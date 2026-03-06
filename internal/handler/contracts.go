package handler

import (
	"context"
	"net/http"
)

type Handler interface {
	CreateEntity(ctx context.Context) http.HandlerFunc
	GetEntity(ctx context.Context) http.HandlerFunc
}
