package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kristux/gophercon/pkg/routing"
	"github.com/kristux/gophercon/pkg/webserver"
)

// go run .cmd/gophercon/gophercon
func main() {
	log.Printf("Service is starting...")
	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		log.Fatal("Service port was not set")
	}
	r := routing.BaseRouter()
	ws := webserver.New("", port, r)
	log.Fatal(ws.Start())
}

func homeHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request is processing: %s", r.URL.Path)
		fmt.Fprint(w, "Hello! Your request was processed.")
	}
}
