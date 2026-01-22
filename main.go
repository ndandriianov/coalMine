package main

import (
	m "coalMine/mine"
	"fmt"
	"time"
)

func main() {
	theMine := m.NewMine()
	theMine.Run()

	time.Sleep(2 * time.Second)
	theMine.Stop()

	fmt.Println("coal extracted:", theMine.Balance)
}
