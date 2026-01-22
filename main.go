package main

import (
	"coalMine/mine"
	"coalMine/transport"
	"fmt"
)

func main() {
	m := mine.NewMine()
	handlers := transport.NewHandlers(m)
	server := transport.NewServer(handlers)

	if err := server.Serve(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
