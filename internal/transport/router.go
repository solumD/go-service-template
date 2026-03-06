package transport

import (
	"context"
	"log/slog"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(ctx context.Context, log *slog.Logger, handler Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Route("/entity", func(r chi.Router) {
		r.Post("/", handler.CreateEntity(ctx))
		r.Get("/{id}", handler.GetEntityByID(ctx))
	})

	return r
}
