package handler

import (
	"context"
	"log/slog"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	mw "github.com/solumD/go-service-template/internal/handler/middleware"
)

func NewRouter(ctx context.Context, log *slog.Logger, handler Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(mw.NewLoggerMW(log))
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Route("/entity", func(r chi.Router) {
		r.Post("/", handler.CreateEntity(ctx))
		r.Get("/{id}", handler.GetEntity(ctx))
	})

	return r
}
