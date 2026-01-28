package mine

import (
	"coalMine/mine/miners"
	"coalMine/mine/pauseController"
	"coalMine/mine/resources"
	"context"
	"fmt"
	"sync"
)

type Service struct {
	miners    []miners.Miner
	Balance   resources.Coal
	producers *sync.WaitGroup
	consumers *sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
	isRunning bool
	runMtx    sync.Mutex

	pc *pauseController.PauseController
}

func NewMine() *Service {
	mineContext, mineCancel := context.WithCancel(context.Background())

	return &Service{
		miners:  make([]miners.Miner, 0, 10),
		Balance: 0,

		producers: &sync.WaitGroup{},
		consumers: &sync.WaitGroup{},

		ctx:       mineContext,
		cancel:    mineCancel,
		isRunning: false,

		pc: pauseController.NewPauseController(),
	}
}

func (s *Service) Start() {
	s.runMtx.Lock()
	if s.isRunning {
		fmt.Println("Mine is already running")
		s.runMtx.Unlock()
		return
	}

	s.isRunning = true
	s.runMtx.Unlock()

	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.producers = &sync.WaitGroup{}
	s.consumers = &sync.WaitGroup{}

	coalChan := make(chan resources.Coal)

	s.producers.Add(1)
	go PassiveIncome(s.ctx, s.producers, coalChan)

	s.consumers.Add(1)
	go func() {
		defer s.consumers.Done()

		for coal := range coalChan {
			s.Balance += coal
		}
	}()

	go func() {
		s.producers.Wait()
		close(coalChan)
	}()
}

func (s *Service) Stop() {
	s.runMtx.Lock()
	defer s.runMtx.Unlock()

	if !s.isRunning {
		fmt.Println("Mine is already not running")
		return
	}

	s.cancel()
	fmt.Println("Stopping the mine")

	s.producers.Wait()
	s.consumers.Wait()

	fmt.Println("Service stopped")
	s.isRunning = false
}

func (s *Service) Pause() {
	s.pc.Pause()
}

func (s *Service) Resume() {
	s.pc.Resume()
}

func (s *Service) HireMiner() {
	miner := miners.NewSmallMiner(s.ctx, s.pc)
	s.miners = append(s.miners, miner)

	s.producers.Add(1)
	go s.miners[0].Run(s.producers)
}
