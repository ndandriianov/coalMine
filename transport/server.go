package transport

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	handlers *Handlers
}

func NewServer(handlers *Handlers) *Server {
	return &Server{handlers: handlers}
}

func (s *Server) Serve() error {
	router := mux.NewRouter()

	router.Path("/mine/restart").Methods("POST").HandlerFunc(s.handlers.HandleRestartMine)
	router.Path("/mine/stop").Methods("POST").HandlerFunc(s.handlers.HandleStop)
	router.Path("/mine/pause").Methods("POST").HandlerFunc(s.handlers.HandlePause)
	router.Path("/mine/resume").Methods("POST").HandlerFunc(s.handlers.HandleResume)

	router.Path("/mine/balance").Methods("GET").HandlerFunc(s.handlers.HandleGetBalance)

	router.Path("/mine/smallMiner/hire").Methods("POST").HandlerFunc(s.handlers.HandleHireSmallMiner)
	router.Path("/mine/miner/start").Methods("POST").HandlerFunc(s.handlers.HandleRunMiner)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}

	return nil
}
