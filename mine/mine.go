package mine

import (
	"coalMine/mine/miners"
	"coalMine/mine/resources"
	"context"
	"fmt"
	"sync"
)

type Mine struct {
	miners    []miners.Miner
	Balance   resources.Coal
	Producers *sync.WaitGroup
	Consumers *sync.WaitGroup
	ctx       context.Context
	cancel    context.CancelFunc
}

func NewMine() Mine {
	producers := sync.WaitGroup{}
	consumers := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	return Mine{
		miners:    make([]miners.Miner, 10),
		Balance:   0,
		Producers: &producers,
		Consumers: &consumers,
		ctx:       ctx,
		cancel:    cancel,
	}
}

func (m *Mine) Run() {
	coalChan := make(chan resources.Coal)

	m.Producers.Add(1)
	go PassiveIncome(m.ctx, m.Producers, coalChan)

	m.Consumers.Add(1)
	go func() {
		defer m.Consumers.Done()

		for coal := range coalChan {
			m.Balance += coal
		}
	}()

	go func() {
		m.Producers.Wait()
		close(coalChan)
	}()
}

func (m *Mine) Stop() {
	m.cancel()
	fmt.Println("Stopping the mine")

	m.Producers.Wait()
	m.Consumers.Wait()

	fmt.Println("Mine stopped")
}
