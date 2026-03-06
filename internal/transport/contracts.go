package transport

import (
	"context"
	"net/http"
)

type Handler interface {
	CreateEntity(ctx context.Context) http.HandlerFunc
	GetEntityByID(ctx context.Context) http.HandlerFunc
}
