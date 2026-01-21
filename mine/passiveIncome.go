package mine

import (
	"coalMine/mine/resources"
	"context"
	"fmt"
	"sync"
	"time"
)

func PassiveIncome(ctx context.Context, wg *sync.WaitGroup, ch chan resources.Coal) {
	defer wg.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("PassiveIncome finished")
			return
		case <-ticker.C:
			select {
			case <-ctx.Done():
				fmt.Println("PassiveIncome finished")
				return
			case ch <- 1:
			}
		}
	}
}
