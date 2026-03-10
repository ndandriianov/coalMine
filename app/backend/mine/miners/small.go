package miners

import (
	"coalMine/mine/pauseController"
	"coalMine/mine/resources"
	"context"
	"fmt"
	"sync"
	"time"
)

type SmallMiner struct {
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

func NewSmallMiner(pc *pauseController.PauseController, coalChan chan resources.Coal) *SmallMiner {
	m := &SmallMiner{
		pc:      pc,
		started: false,

		energy:            30,
		coalPerExtraction: 1,
		coalExtracted:     0,
		sleepTimeSeconds:  3,

		coalChan: coalChan,
	}

	return m
}

func (m *SmallMiner) Run(ctx context.Context, group *sync.WaitGroup) {
	defer group.Done()

	m.startOnce.Do(func() {
		m.started = true
		ticker := time.NewTicker(time.Duration(m.sleepTimeSeconds) * time.Second)
		defer ticker.Stop()

		for {
			if m.energy <= 0 {
				fmt.Println("small miner ran out of energy!")
				return
			}
			select {
			case <-ctx.Done():
				fmt.Println("small miner was forced to finish his work")
				return
			case <-ticker.C:
				if err := m.pc.WaitIfPaused(ctx); err != nil {
					fmt.Println("small miner was forced to finish his work")
					return
				}

				select {
				case <-m.pc.PauseChan():
					fmt.Println("small miner on pause")
					if err := m.pc.WaitIfPaused(ctx); err != nil {
						fmt.Println("small miner was forced to finish his work")
						return
					}
				case m.coalChan <- m.coalPerExtraction:
					fmt.Printf("small miner put %v coal to coal chan\n", m.coalPerExtraction)
					m.infoMtx.Lock()
					m.coalExtracted += m.coalPerExtraction
					m.energy--
					m.infoMtx.Unlock()
				}
			}
		}
	})
}

func (m *SmallMiner) Info() MinerInfo {
	m.infoMtx.Lock()
	defer m.infoMtx.Unlock()

	return MinerInfo{
		Type:              "small",
		EnergyLeft:        m.energy,
		CoalPerExtraction: m.coalPerExtraction,
		CoalExtracted:     int(m.coalExtracted),
		Started:           m.started,
		SleepTimeSeconds:  m.sleepTimeSeconds,
	}
}

func (m *SmallMiner) HasStarted() bool {
	return m.started
}
