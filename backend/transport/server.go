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
	// methods that work anytime
	router := mux.NewRouter()

	router.Path("/mine/start").Methods("POST").HandlerFunc(s.handlers.HandleStartMine)
	router.Path("/mine/start").Methods("GET").HandlerFunc(s.handlers.HandleMineStarted)
	router.Path("/mine/restart").Methods("POST").HandlerFunc(s.handlers.HandleRestartMine)
	router.Path("/mine/stop").Methods("POST").HandlerFunc(s.handlers.HandleStop)

	router.Path("/mine/pause").Methods("GET").HandlerFunc(s.handlers.IsOnPause)

	router.Path("/mine/balance").Methods("GET").HandlerFunc(s.handlers.HandleGetBalance)
	router.Path("/mine/miner").Methods("GET").HandlerFunc(s.handlers.HandleGetMiners)
	router.Path("/mine/equipment").Methods("GET").HandlerFunc(s.handlers.HandleGetEquipment)

	// methods that don't work if not started
	protected := router.NewRoute().Subrouter()
	protected.Use(HasStartedMiddleware(s.handlers))

	protected.Path("/mine/pause").Methods("POST").HandlerFunc(s.handlers.HandlePause)
	protected.Path("/mine/resume").Methods("POST").HandlerFunc(s.handlers.HandleResume)

	// methods that don't work on pause
	pauseProtected := protected.NewRoute().Subrouter()
	pauseProtected.Use(PauseMiddleware(s.handlers))

	pauseProtected.Path("/mine/miner/hire").Methods("POST").HandlerFunc(s.handlers.HandleHireMiner)
	pauseProtected.Path("/mine/miner/start").Methods("POST").HandlerFunc(s.handlers.HandleRunMiner)
	pauseProtected.Path("/mine/equipment").Methods("POST").HandlerFunc(s.handlers.HandleBuyEquipment)

	handler := CorsMiddleware(router)

	if err := http.ListenAndServe(":9091", handler); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}

	return nil
}
