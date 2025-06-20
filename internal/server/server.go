package server

import (
	"fmt"
	"net/http"

	"internal/config"
	"internal/handler"
)

type Server struct {
	cfg config.Config
}

func New(cfg config.Config) *Server {
	return &Server{cfg: cfg}
}

func (s *Server) Start() error {
	http.HandleFunc("/", handler.Home)

	addr := fmt.Sprintf(":%s", s.cfg.Port)
	fmt.Printf("Server is running at http://localhost%s\n", addr)
	return http.ListenAndServe(addr, nil)
}

