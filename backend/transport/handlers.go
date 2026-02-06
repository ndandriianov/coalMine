package transport

import (
	"coalMine/mine"
	"coalMine/mine/equipment"
	mineErrors "coalMine/mine/errors"
	"coalMine/transport/dto"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

func (h *Handlers) IsOnPause(w http.ResponseWriter, r *http.Request) {
	pause := dto.Pause{IsOnPause: h.mine.IsOnPause()}
	err := json.NewEncoder(w).Encode(pause)
	if err != nil {
		logHttpWriteFailure()
	}
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
		} else if errors.Is(err, mineErrors.ErrNotEnoughCoal) {
			http.Error(w, err.Error(), http.StatusForbidden)
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

// Equipment

func (h *Handlers) HandleBuyEquipment(w http.ResponseWriter, r *http.Request) {
	var purchase dto.EquipmentPurchase
	if err := json.NewDecoder(r.Body).Decode(&purchase); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var eqType equipment.Type
	switch strings.ToLower(purchase.Equipment) {
	case "pickaxe":
		eqType = equipment.Pickaxe
	case "ventilation":
		eqType = equipment.Ventilation
	case "minecarts":
		eqType = equipment.Minecarts
	default:
		http.Error(w, "equipment type is not specified", http.StatusBadRequest)
		return
	}

	err := h.mine.BuyEquipment(eqType)
	if err != nil {
		if errors.Is(err, mineErrors.ErrEquipmentIsAlreadyBought) {
			http.Error(w, err.Error(), http.StatusConflict)
		} else if errors.Is(err, mineErrors.ErrNotEnoughCoal) {
			http.Error(w, err.Error(), http.StatusForbidden)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *Handlers) HandleGetEquipment(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(h.mine.GetEquipment())
	if err != nil {
		logHttpWriteFailure()
	}
}
