package miners

import (
	"coalMine/mine/resources"
	"context"
	"fmt"
	"sync"
	"time"
)

type SmallMiner struct {
	ctx context.Context

	cond     *sync.Cond
	state    MinerState
	stateMtx sync.Mutex

	energy        int
	coalExtracted resources.Coal
	infoMtx       sync.Mutex
}

func NewSmallMiner(ctx context.Context) *SmallMiner {
	m := &SmallMiner{
		ctx:    ctx,
		state:  NotStarted,
		energy: 30,
	}

	m.cond = sync.NewCond(&m.stateMtx)

	return m
}

func (m *SmallMiner) Run(group *sync.WaitGroup) {
	defer group.Done()

	m.stateMtx.Lock()
	if m.state != NotStarted {
		fmt.Println("Cannot run started miner")
		m.stateMtx.Unlock()
		return
	}

	m.state = Working
	m.stateMtx.Unlock()

	//ticker := time.NewTicker(3 * time.Second)

	for {
		select {
		case <-m.ctx.Done():
			fmt.Println("small miner was forced to finish his work")
			return
		//case <-ticker.C:
		//	if m.energy == 0 {
		//		fmt.Println("small miner finished his work")
		//		return
		//	}
		//	m.infoMtx.Lock()
		//
		//	coalExtracted := resources.Coal(1)
		//	ch <- coalExtracted
		//	m.coalExtracted += coalExtracted
		//	m.energy--
		//
		//	m.infoMtx.Unlock()
		//
		//	fmt.Println("small miner extracted coal")
		default:
			m.stateMtx.Lock()
			for m.state == Pause {
				m.cond.Wait()
			}
			m.stateMtx.Unlock()

			fmt.Println("small miner working")
			time.Sleep(3 * time.Second)
		}
	}
}

func (m *SmallMiner) Pause() {
	m.stateMtx.Lock()
	m.state = Pause
	m.stateMtx.Unlock()
}

func (m *SmallMiner) Resume() {
	m.stateMtx.Lock()
	m.state = Working
	m.stateMtx.Unlock()
	m.cond.Signal()
}

func (m *SmallMiner) Info() MinerInfo {
	m.infoMtx.Lock()
	defer m.infoMtx.Unlock()

	return MinerInfo{
		EnergyLeft:    m.energy,
		CoalExtracted: int(m.coalExtracted),
	}
}
