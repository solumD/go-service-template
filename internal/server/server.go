package server

import (
	"context"
	"net/http"
	"time"
)

const (
	ReadTimeout    = 10 * time.Second
	WriteTimeout   = 10 * time.Second
	MaxHeaderBytes = 1 << 20 // 1 MB

)

type Server struct {
	httpServer *http.Server
}

func New(addr string, router http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           addr,
			Handler:        router,
			ReadTimeout:    ReadTimeout,
			WriteTimeout:   WriteTimeout,
			MaxHeaderBytes: MaxHeaderBytes,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
