package http

import (
	"context"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Handler interface {
	InitRoutes(ctx context.Context, r chi.Router)
}

func NewRouter(ctx context.Context, handlers ...Handler) chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	for _, h := range handlers {
		h.InitRoutes(ctx, r)
	}

	return r
}
