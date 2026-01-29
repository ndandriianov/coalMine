package mine

import (
	"coalMine/mine/pauseController"
	"coalMine/mine/resources"
	"context"
	"fmt"
	"sync"
	"time"
)

func PassiveIncome(
	ctx context.Context,
	wg *sync.WaitGroup,
	pc *pauseController.PauseController,
	coalChan chan resources.Coal,
) {
	defer wg.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("PassiveIncome finished")
			return
		case <-ticker.C:
			if err := pc.WaitIfPaused(ctx); err != nil {
				fmt.Println("PassiveIncome finished")
				return
			}

			select {
			case <-pc.PauseChan():
				fmt.Println("passive income pause")
				if err := pc.WaitIfPaused(ctx); err != nil {
					fmt.Println("PassiveIncome finished")
					return
				}
			case coalChan <- 1:
				fmt.Println("passive income +1")
			}
		}
	}
}
