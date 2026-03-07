package transport

import (
	"context"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewRouter(ctx context.Context, handler Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Route("/entity", func(r chi.Router) {
		r.Post("/", handler.CreateEntity(ctx))
		r.Get("/{id}", handler.GetEntityByID(ctx))
	})

	return r
}
