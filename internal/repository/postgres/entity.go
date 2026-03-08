package postgres

import (
	"context"
	"log/slog"

	"github.com/solumD/go-service-template/internal/model"
	"github.com/solumD/go-service-template/pkg/helper"
	"github.com/solumD/go-service-template/pkg/logger"
	"github.com/solumD/go-service-template/pkg/postgres"
)

type entityRepository struct {
	db  *postgres.Postgres
	log *slog.Logger
}

func NewEntityRepository(pg *postgres.Postgres, l *slog.Logger) *entityRepository {
	return &entityRepository{
		db:  pg,
		log: l,
	}
}

func (r *entityRepository) CreateEntity(ctx context.Context, entity *model.Entity) (int, error) {
	fn := helper.GetCurrentFunctionName()
	log := r.log.With(logger.String("fn", fn))

	query := `INSERT INTO entity (name) VALUES ($1) RETURNING id`

	log.Debug("executing query", logger.String("query", query), logger.String("name", entity.Name))

	var entityID int
	err := r.db.Pool().QueryRow(ctx, query, entity.Name).Scan(&entityID)
	if err != nil {
		log.Error("error while executing query", logger.Error(err))

		return 0, err
	}

	return entityID, nil
}

func (r *entityRepository) GetEntityByID(ctx context.Context, id int) (*model.Entity, error) {
	fn := helper.GetCurrentFunctionName()
	log := r.log.With(logger.String("fn", fn))

	query := `SELECT name FROM entity WHERE id = $1`

	log.Debug("executing query", logger.String("query", query), logger.Int("id", id))

	var entityName string
	err := r.db.Pool().QueryRow(ctx, query, id).Scan(&entityName)
	if err != nil {
		log.Error("error while executing query", logger.Error(err))

		return nil, err
	}

	return &model.Entity{ID: id, Name: entityName}, nil
}
