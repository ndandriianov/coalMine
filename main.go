package main

import (
	m "coalMine/mine"
	"context"
	"fmt"
	"time"
)

func main() {
	myCtx, myCancel := context.WithCancel(context.Background())

	theMine := m.NewMine()
	theMine.Run(myCtx)

	time.Sleep(2 * time.Second)
	myCancel()

	theMine.Producers.Wait()
	theMine.Consumers.Wait()

	fmt.Println("coal extracted:", theMine.Balance)
}
