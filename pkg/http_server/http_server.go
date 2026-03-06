package httpserver

import (
	"context"
	"log"
	"net/http"
	"time"
)

const (
	ReadTimeout    = 10 * time.Second
	WriteTimeout   = 10 * time.Second
	MaxHeaderBytes = 1 << 20 // 1 MB
)

type server struct {
	httpServer *http.Server
}

func New(addr string, router http.Handler) *server {
	return &server{
		httpServer: &http.Server{
			Addr:           addr,
			Handler:        router,
			ReadTimeout:    ReadTimeout,
			WriteTimeout:   WriteTimeout,
			MaxHeaderBytes: MaxHeaderBytes,
		},
	}
}

func (s *server) Run() {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("%v\n", err)
		}
	}()

}

func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
