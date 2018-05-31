package routing

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// DiagnosticsRouter to handle health and ready checks
func DiagnosticsRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/healthz", homeHandler()).Methods(http.MethodGet)
	r.HandleFunc("/readyz", homeHandler()).Methods(http.MethodGet)
	return r
}

func readyzHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request is processing: %s", r.URL.Path)
		fmt.Fprint(w, "READY YO!")
	}
}

func healthzHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request is processing: %s", r.URL.Path)
		fmt.Fprint(w, "HEALTH YO!")
	}
}
