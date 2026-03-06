package usecase

import (
	"context"
	"log/slog"

	"github.com/solumD/go-service-template/internal/model"
	"github.com/solumD/go-service-template/pkg/logger"
)

type entityUsecase struct {
	repository EntityRepository
	log        *slog.Logger
}

func NewEntityUsecase(r EntityRepository, l *slog.Logger) *entityUsecase {
	return &entityUsecase{
		repository: r,
		log:        l,
	}
}

func (uc *entityUsecase) CreateEntity(ctx context.Context, entity *model.Entity) (int, error) {
	const fn = "entityUsecase.CreateEntity"
	log := uc.log.With(logger.String("fn", fn))

	entityID, err := uc.repository.CreateEntity(ctx, entity)
	if err != nil {
		log.Error("failed to save entity in repository", logger.Error(err))

		return 0, err
	}

	return entityID, nil
}

func (uc *entityUsecase) GetEntityByID(ctx context.Context, id int) (*model.Entity, error) {
	const fn = "entityUsecase.GetEntityByID"
	log := uc.log.With(logger.String("fn", fn))

	entity, err := uc.repository.GetEntityByID(ctx, id)
	if err != nil {
		log.Error("failed to get entity by id from repository", logger.Error(err))

		return nil, err
	}

	return entity, nil
}
