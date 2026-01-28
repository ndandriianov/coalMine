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
	ctx       context.Context
	pc        *pauseController.PauseController
	startOnce sync.Once

	energy        int
	coalExtracted resources.Coal
	infoMtx       sync.Mutex
}

func NewSmallMiner(ctx context.Context, pc *pauseController.PauseController) *SmallMiner {
	m := &SmallMiner{
		ctx: ctx,
		pc:  pc,

		energy:        30,
		coalExtracted: 0,
	}

	return m
}

func (m *SmallMiner) Run(group *sync.WaitGroup) {
	defer group.Done()

	m.startOnce.Do(func() {
		ticker := time.NewTicker(3 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-m.ctx.Done():
				fmt.Println("small miner was forced to finish his work")
				return
			case <-ticker.C:
				if err := m.pc.WaitIfPaused(m.ctx); err != nil {
					return
				}

				fmt.Println("small miner working")
			}
		}
	})
}

func (m *SmallMiner) Info() MinerInfo {
	m.infoMtx.Lock()
	defer m.infoMtx.Unlock()

	return MinerInfo{
		EnergyLeft:    m.energy,
		CoalExtracted: int(m.coalExtracted),
	}
}
