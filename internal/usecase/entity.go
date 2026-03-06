package usecase

import (
	"context"
	"fmt"

	"github.com/solumD/go-service-template/internal/model"
	"github.com/solumD/go-service-template/pkg/logger"
	"go.uber.org/zap"
)

type entityUsecase struct {
	repository EntityRepository
}

func New(r EntityRepository) *entityUsecase {
	return &entityUsecase{
		repository: r,
	}
}

func (uc *entityUsecase) CreateEntity(ctx context.Context, entity *model.Entity) (int64, error) {
	// - some validations
	// - some checks
	// - some usecases
	// - etc.

	entityID, err := uc.repository.CreateEntity(ctx, entity)
	if err != nil {
		logger.Error("[service] failed to save entity in repository",
			zap.Error(err),
			zap.Any("entity", entity),
		)

		return 0, fmt.Errorf("[service] failed to save entity in repository: %v", err)
	}

	return entityID, nil
}

func (uc *entityUsecase) GetEntity(ctx context.Context, id int64) (*model.Entity, error) {
	// - some validations
	// - some checks
	// - some usecases
	// - etc.

	entity, err := uc.repository.GetEntity(ctx, id)
	if err != nil {
		logger.Error("[service] failed to get entity from repository",
			zap.Error(err),
			zap.Int64("id", id),
		)

		return nil, fmt.Errorf("[service] failed to get entity from repository: %v", err)
	}

	return entity, nil
}
