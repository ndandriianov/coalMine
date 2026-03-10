package miners

import (
	"coalMine/mine/pauseController"
	"coalMine/mine/resources"
	"context"
	"fmt"
	"sync"
	"time"
)

type StrongMiner struct {
	pc        *pauseController.PauseController
	startOnce sync.Once
	started   bool // for making list of unstarted miner

	energy            int
	coalPerExtraction resources.Coal
	coalExtracted     resources.Coal
	sleepTimeSeconds  int
	infoMtx           sync.Mutex

	coalChan chan resources.Coal
}

func NewStrongMiner(pc *pauseController.PauseController, coalChan chan resources.Coal) *StrongMiner {
	m := &StrongMiner{
		pc:      pc,
		started: false,

		energy:            60,
		coalPerExtraction: 10,
		coalExtracted:     0,
		sleepTimeSeconds:  1,

		coalChan: coalChan,
	}

	return m
}

func (m *StrongMiner) Run(ctx context.Context, group *sync.WaitGroup) {
	defer group.Done()

	m.startOnce.Do(func() {
		m.started = true
		ticker := time.NewTicker(time.Duration(m.sleepTimeSeconds) * time.Second)
		defer ticker.Stop()

		for {
			if m.energy <= 0 {
				fmt.Println("strong miner ran out of energy!")
				return
			}
			select {
			case <-ctx.Done():
				fmt.Println("strong miner was forced to finish his work")
				return
			case <-ticker.C:
				if err := m.pc.WaitIfPaused(ctx); err != nil {
					fmt.Println("strong miner was forced to finish his work")
					return
				}

				select {
				case <-m.pc.PauseChan():
					fmt.Println("strong miner on pause")
					if err := m.pc.WaitIfPaused(ctx); err != nil {
						fmt.Println("strong miner was forced to finish his work")
						return
					}
				case m.coalChan <- m.coalPerExtraction:
					fmt.Printf("strong miner put %v coal to coal chan\n", m.coalPerExtraction)
					m.infoMtx.Lock()
					m.coalExtracted += m.coalPerExtraction
					m.coalPerExtraction += 3
					m.energy--
					m.infoMtx.Unlock()
				}
			}
		}
	})
}

func (m *StrongMiner) Info() MinerInfo {
	m.infoMtx.Lock()
	defer m.infoMtx.Unlock()

	return MinerInfo{
		Type:              "strong",
		EnergyLeft:        m.energy,
		CoalPerExtraction: m.coalPerExtraction,
		CoalExtracted:     int(m.coalExtracted),
		Started:           m.started,
		SleepTimeSeconds:  m.sleepTimeSeconds,
	}
}

func (m *StrongMiner) HasStarted() bool {
	return m.started
}
