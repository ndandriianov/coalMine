package mine

import (
	"coalMine/helpers"
	"coalMine/mine/miners"
	"coalMine/mine/pauseController"
	"coalMine/mine/resources"
	"context"
	"fmt"
	"sync"
)

type Service struct {
	miners     []miners.Miner
	minersMtx  sync.Mutex
	Balance    resources.Coal
	BalanceMtx sync.Mutex
	coalChan   chan resources.Coal

	producers *sync.WaitGroup
	consumers *sync.WaitGroup

	ctx       context.Context
	cancel    context.CancelFunc
	startOnce sync.Once

	pc *pauseController.PauseController
}

func NewService() *Service {
	mineContext, mineCancel := context.WithCancel(context.Background())

	return &Service{
		miners:   make([]miners.Miner, 0, 10),
		Balance:  0,
		coalChan: make(chan resources.Coal),

		producers: &sync.WaitGroup{},
		consumers: &sync.WaitGroup{},

		ctx:    mineContext,
		cancel: mineCancel,

		pc: pauseController.NewPauseController(),
	}
}

func (s *Service) Start() {
	s.startOnce.Do(func() {
		s.ctx, s.cancel = context.WithCancel(context.Background())
		s.producers = &sync.WaitGroup{}
		s.consumers = &sync.WaitGroup{}

		s.producers.Add(1)
		go PassiveIncome(s.ctx, s.producers, s.pc, s.coalChan)

		s.consumers.Add(1)
		go s.collectCoal()

		go func() {
			s.producers.Wait()
			close(s.coalChan)
		}()
	})
}

func (s *Service) Stop() {
	s.cancel()
	fmt.Println("Stopping the mine")

	s.producers.Wait()
	s.consumers.Wait()

	fmt.Println("Service stopped")
}

func (s *Service) Pause() {
	s.pc.Pause()
}

func (s *Service) Resume() {
	s.pc.Resume()
}

// HireMiner adds a miner to the mine.
//
// Returns the miners id.
func (s *Service) HireMiner() int {
	miner := miners.NewSmallMiner(s.pc, s.coalChan)

	s.minersMtx.Lock()
	defer s.minersMtx.Unlock()
	s.miners = append(s.miners, miner)

	return len(s.miners)
}

func (s *Service) RunMiner(id int) bool {
	s.minersMtx.Lock()
	currentMiner, ok := helpers.SafeGet(id, s.miners)
	s.minersMtx.Unlock()

	if !ok {
		return false
	}

	s.producers.Add(1)
	go currentMiner.Run(s.ctx, s.producers)

	return true
}

func (s *Service) RunAllNotStartedMiners() {
	s.minersMtx.Lock()
	for _, miner := range s.miners {
		if !miner.HasStarted() {
			s.producers.Add(1)
			go miner.Run(s.ctx, s.producers)
		}
	}
	s.minersMtx.Unlock()
}

func (s *Service) collectCoal() {
	defer s.consumers.Done()

	for coal := range s.coalChan {
		s.BalanceMtx.Lock()
		s.Balance += coal
		s.BalanceMtx.Unlock()
	}

	fmt.Println("finished collecting coal")
}
