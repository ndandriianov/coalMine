package mine

import (
	"coalMine/mine/resources"
	"context"
	"fmt"
	"time"
)

func RunPassiveIncome(ctx context.Context) <-chan resources.Coal {
	ch := make(chan resources.Coal)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("PassiveIncome finished")
				return
			default:
				time.Sleep(1 * time.Second)
				ch <- 1
			}
		}
	}()

	return ch
}
