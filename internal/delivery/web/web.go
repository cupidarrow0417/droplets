package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

// New initializes a new webapp server.
func New(cfg Config) *Server {
	router := mux.NewRouter()

	return &Server{
		router: router,
	}
}

// Config represents server configuration.
type Config struct {
	TemplateDir string
}

// Server represents an instance of webapp server.
type Server struct {
	cfg    Config
	router *mux.Router
}

func (srv *Server) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	srv.router.ServeHTTP(wr, req)
}