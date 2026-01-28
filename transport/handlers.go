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

func NewHandlers(mine *mine.Service) *Handlers {
	return &Handlers{
		mine: mine,
	}
}

func (h *Handlers) HandleStart(w http.ResponseWriter, r *http.Request) {
	h.mine.Start()
}

func (h *Handlers) HandleStop(w http.ResponseWriter, r *http.Request) {
	h.mine.Stop()
}

func (h *Handlers) HandleGetBalance(w http.ResponseWriter, r *http.Request) {
	balance := dto.Balance{Balance: int(h.mine.Balance)}
	err := json.NewEncoder(w).Encode(balance)
	if err != nil {
		logHttpWriteFailure()
	}
}

func logHttpWriteFailure() {
	fmt.Println("failed to write http response")
}

// MINER HANDLERS

func (h *Handlers) HandleRunSmallMiner(w http.ResponseWriter, r *http.Request) {
	h.mine.HireMiner()
}

func (h *Handlers) HandlePauseSmallMiner(w http.ResponseWriter, r *http.Request) {
	h.mine.PauseMiner()
}

func (h *Handlers) HandleResumeSmallMiner(w http.ResponseWriter, r *http.Request) {
	h.mine.ResumeMiner()
}
