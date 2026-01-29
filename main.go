package main

import (
	"coalMine/transport"
	"fmt"
)

func main() {
	handlers := transport.NewHandlers()
	server := transport.NewServer(handlers)

	if err := server.Serve(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
