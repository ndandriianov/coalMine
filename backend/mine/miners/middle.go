package miners

import (
	"coalMine/mine/pauseController"
	"coalMine/mine/resources"
	"context"
	"fmt"
	"sync"
	"time"
)

type MiddleMiner struct {
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

func NewMiddleMiner(pc *pauseController.PauseController, coalChan chan resources.Coal) *MiddleMiner {
	m := &MiddleMiner{
		pc:      pc,
		started: false,

		energy:            45,
		coalPerExtraction: 3,
		coalExtracted:     0,
		sleepTimeSeconds:  2,

		coalChan: coalChan,
	}

	return m
}

func (m *MiddleMiner) Run(ctx context.Context, group *sync.WaitGroup) {
	defer group.Done()

	m.startOnce.Do(func() {
		m.started = true
		ticker := time.NewTicker(time.Duration(m.sleepTimeSeconds) * time.Second)
		defer ticker.Stop()

		for {
			if m.energy <= 0 {
				fmt.Println("middle miner ran out of energy!")
				return
			}
			select {
			case <-ctx.Done():
				fmt.Println("middle miner was forced to finish his work")
				return
			case <-ticker.C:
				if err := m.pc.WaitIfPaused(ctx); err != nil {
					fmt.Println("middle miner was forced to finish his work")
					return
				}

				select {
				case <-m.pc.PauseChan():
					fmt.Println("middle miner on pause")
					if err := m.pc.WaitIfPaused(ctx); err != nil {
						fmt.Println("middle miner was forced to finish his work")
						return
					}
				case m.coalChan <- m.coalPerExtraction:
					fmt.Printf("middle miner put %v coal to coal chan\n", m.coalPerExtraction)
					m.infoMtx.Lock()
					m.coalExtracted += m.coalPerExtraction
					m.energy--
					m.infoMtx.Unlock()
				}
			}
		}
	})
}

func (m *MiddleMiner) Info() MinerInfo {
	m.infoMtx.Lock()
	defer m.infoMtx.Unlock()

	return MinerInfo{
		Type:              "middle",
		EnergyLeft:        m.energy,
		CoalPerExtraction: m.coalPerExtraction,
		CoalExtracted:     int(m.coalExtracted),
		Started:           m.started,
		SleepTimeSeconds:  m.sleepTimeSeconds,
	}
}

func (m *MiddleMiner) HasStarted() bool {
	return m.started
}
