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
	producers := sync.WaitGroup{}
	consumers := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	return &Service{
		miners:    make([]miners.Miner, 10),
		Balance:   0,
		producers: &producers,
		consumers: &consumers,
		ctx:       ctx,
		cancel:    cancel,
	}
}

func (m *Service) Start() {
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
