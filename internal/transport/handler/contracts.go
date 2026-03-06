package handler

import (
	"context"

	"github.com/solumD/go-service-template/internal/model"
)

type EntityUsecase interface {
	CreateEntity(ctx context.Context, entity *model.Entity) (int, error)
	GetEntityByID(ctx context.Context, id int) (*model.Entity, error)
}
