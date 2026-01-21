package mine

import (
	"coalMine/mine/miners"
	"coalMine/mine/resources"
	"context"
	"sync"
)

type Mine struct {
	miners    []miners.Miner
	Balance   resources.Coal
	Producers *sync.WaitGroup
	Consumers *sync.WaitGroup
}

func NewMine() Mine {
	producers := sync.WaitGroup{}
	consumers := sync.WaitGroup{}
	return Mine{
		miners:    make([]miners.Miner, 10),
		Balance:   0,
		Producers: &producers,
		Consumers: &consumers,
	}
}

func (m *Mine) Run(ctx context.Context) {
	coalChan := make(chan resources.Coal)

	m.Producers.Add(1)
	go PassiveIncome(ctx, m.Producers, coalChan)

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
