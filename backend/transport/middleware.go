package transport

import "net/http"

var allowedOrigins = map[string]bool{
	"http://localhost:5173": true,
}

var allowedMethods = "GET, POST, OPTIONS"
var allowedHeaders = "Content-Type, Authorization, Accept"

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")
		if isOriginAllowed(origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
		}

		w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
		w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func isOriginAllowed(origin string) bool {
	return allowedOrigins[origin]
}

func HasStartedMiddleware(h *Handlers) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if !h.mine.Started() {
				http.Error(w, "game has not started", http.StatusLocked)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func PauseMiddleware(h *Handlers) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if h.mine.IsOnPause() {
				http.Error(w, "game is on pause", http.StatusLocked)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
