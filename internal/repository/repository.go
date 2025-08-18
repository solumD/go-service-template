package repository

import (
	"context"

	"github.com/solumD/go-service-template/internal/model"
)

type Repository interface {
	CreateEntity(ctx context.Context, entity *model.Entity) (int64, error)
	GetEntity(ctx context.Context, id int64) (*model.Entity, error)
}
