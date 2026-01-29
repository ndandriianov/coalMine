package transport

import (
	"coalMine/mine"
	"coalMine/transport/dto"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handlers struct {
	mine *mine.Service
}

func NewHandlers() *Handlers {
	m := mine.NewService()
	m.Start()
	return &Handlers{
		mine: m,
	}
}

func (h *Handlers) HandleRestartMine(w http.ResponseWriter, r *http.Request) {
	h.mine = mine.NewService()
	h.mine.Start()
}

func (h *Handlers) HandleStop(w http.ResponseWriter, r *http.Request) {
	h.mine.Stop()
}

func (h *Handlers) HandlePause(w http.ResponseWriter, r *http.Request) {
	h.mine.Pause()
}

func (h *Handlers) HandleResume(w http.ResponseWriter, r *http.Request) {
	h.mine.Resume()
}

func (h *Handlers) HandleGetBalance(w http.ResponseWriter, r *http.Request) {
	balance := dto.Balance{Balance: int(h.mine.Balance)}
	err := json.NewEncoder(w).Encode(balance)
	if err != nil {
		logHttpWriteFailure()
	}
}

// MINER HANDLERS

func logHttpWriteFailure() {
	fmt.Println("failed to write http response")
}

func (h *Handlers) HandleRunSmallMiner(w http.ResponseWriter, r *http.Request) {
	h.mine.HireMiner()
}
