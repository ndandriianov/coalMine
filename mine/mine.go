package mine

import (
	"coalMine/mine/miners"
	"coalMine/mine/resources"
	"context"
	"fmt"
	"sync"
)

type Mine struct {
	miners  []miners.Miner
	balance resources.Coal
	Group   sync.WaitGroup
}

func NewMine() Mine {
	return Mine{
		miners:  make([]miners.Miner, 10),
		balance: 0,
	}
}

func (m *Mine) Run(ctx context.Context) {
	coalChan := RunPassiveIncome(ctx)

	m.Group.Add(1)
	go func() {
		defer m.Group.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("mine ended its work! coal extracted:", m.balance)
				return
			case coal := <-coalChan:
				m.balance += coal
			}
		}
	}()
}
