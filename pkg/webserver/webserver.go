package webserver

import (
	"context"
	"net"
	"net/http"
)

// WebServer type
type WebServer struct {
	http.Server
	address string
}

// New creates a new WebServer
func New(host, port string, h http.Handler) *WebServer {
	var ws WebServer

	ws.Addr = net.JoinHostPort(host, port)
	ws.Handler = h

	return &ws
}

// Start starts up the web server
func (s *WebServer) Start() error {
	return s.ListenAndServe()
}

// Stop Stops the WebServer
func (s *WebServer) Stop() error {
	return s.Shutdown(context.Background())
}
