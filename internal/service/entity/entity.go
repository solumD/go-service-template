package entity

import (
	"context"
	"fmt"

	"github.com/solumD/go-service-template/internal/model"
	"github.com/solumD/go-service-template/internal/repository"
	"github.com/solumD/go-service-template/internal/service"
	"github.com/solumD/go-service-template/pkg/logger"
	"go.uber.org/zap"
)

type entityService struct {
	repository repository.Repository
}

func New(r repository.Repository) service.EntityService {
	return &entityService{
		repository: r,
	}
}

func (s *entityService) SaveEntity(ctx context.Context, entity *model.Entity) (int64, error) {
	// - some validations
	// - some checks
	// - etc.

	entityID, err := s.repository.SaveEntity(ctx, entity)
	if err != nil {
		logger.Error("[service] failed to save entity in repository",
			zap.Error(err),
			zap.Any("entity", entity),
		)

		return 0, fmt.Errorf("[service] failed to save entity in repository: %v", err)
	}

	return entityID, nil
}

func (s *entityService) GetEntity(ctx context.Context, id int64) (*model.Entity, error) {
	// - some validations
	// - some checks
	// - etc.

	entity, err := s.repository.GetEntity(ctx, id)
	if err != nil {
		logger.Error("[service] failed to get entity from repository",
			zap.Error(err),
			zap.Int64("id", id),
		)

		return nil, fmt.Errorf("[service] failed to get entity from repository: %v", err)
	}

	logger.Info("[service] got entity from repository",
		zap.Any("entity", entity),
	)

	return entity, nil
}
