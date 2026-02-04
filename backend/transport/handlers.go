package transport

import (
	"coalMine/mine"
	mineErrors "coalMine/mine/errors"
	"coalMine/transport/dto"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type Handlers struct {
	mine *mine.Service
}

func NewHandlers() *Handlers {
	m := mine.NewService()
	return &Handlers{
		mine: m,
	}
}

func (h *Handlers) HandleStartMine(w http.ResponseWriter, r *http.Request) {
	h.mine.Start()
}

func (h *Handlers) HandleMineStarted(w http.ResponseWriter, r *http.Request) {
	started := dto.Started{Started: h.mine.Started()}
	err := json.NewEncoder(w).Encode(started)
	if err != nil {
		logHttpWriteFailure()
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

func logHttpWriteFailure() {
	fmt.Println("failed to write http response")
}

// MINER HANDLERS

func (h *Handlers) HandleHireMiner(w http.ResponseWriter, r *http.Request) {
	minerType := r.URL.Query().Get("type")
	if minerType == "" {
		http.Error(w, "type is required", http.StatusBadRequest)
		return
	}

	id, err := h.mine.HireMiner(minerType)
	if err != nil {
		if errors.Is(err, mineErrors.ErrInvalidMinerType) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(map[string]int{
		"id": id,
	})
	if err != nil {
		fmt.Println("serializing error")
	}
}

func (h *Handlers) HandleRunMiner(w http.ResponseWriter, r *http.Request) {
	rawId := r.URL.Query().Get("id")
	if id, err := strconv.Atoi(rawId); err == nil {
		if !h.mine.RunMiner(id) {
			w.WriteHeader(http.StatusNotFound)
		}
		return
	}
	h.mine.RunAllNotStartedMiners()
}

func (h *Handlers) HandleGetMiners(w http.ResponseWriter, r *http.Request) {
	minerType := r.URL.Query().Get("type")
	miners := h.mine.GetMiners(minerType)

	err := json.NewEncoder(w).Encode(miners)
	if err != nil {
		logHttpWriteFailure()
	}
}
