package mine

import (
	"coalMine/mine/miners"
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
}

func NewMine() *Service {
	return &Service{
		miners:  make([]miners.Miner, 10),
		Balance: 0,
	}
}

func (m *Service) Start() {
	m.ctx, m.cancel = context.WithCancel(context.Background())
	m.producers = &sync.WaitGroup{}
	m.consumers = &sync.WaitGroup{}

	coalChan := make(chan resources.Coal)

	m.producers.Add(1)
	go PassiveIncome(m.ctx, m.producers, coalChan)

	m.consumers.Add(1)
	go func() {
		defer m.consumers.Done()

		for coal := range coalChan {
			m.Balance += coal
		}
	}()

	go func() {
		m.producers.Wait()
		close(coalChan)
	}()
}

func (m *Service) Stop() {
	m.cancel()
	fmt.Println("Stopping the mine")

	m.producers.Wait()
	m.consumers.Wait()

	fmt.Println("Service stopped")
}
