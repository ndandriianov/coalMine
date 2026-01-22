package transport

import (
	"coalMine/mine"
	"net/http"
)

type Handlers struct {
	m mine.Mine
}

func (h *Handlers) HandleStop(w http.ResponseWriter, r *http.Request) {

}
