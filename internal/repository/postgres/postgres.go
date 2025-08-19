package postgres

import (
	"context"
	"fmt"

	"github.com/solumD/go-service-template/internal/model"
	"github.com/solumD/go-service-template/internal/repository"
	"github.com/solumD/go-service-template/pkg/postgres"
)

type entityRepository struct {
	pg *postgres.Postgres
}

func New(postgres *postgres.Postgres) repository.Repository {
	return &entityRepository{
		pg: postgres,
	}
}

func (r *entityRepository) CreateEntity(ctx context.Context, entity *model.Entity) (int64, error) {
	query := `INSERT INTO entity (name) VALUES ($1) RETURNING id`

	var entityID int64
	err := r.pg.Pool().QueryRow(ctx, query, entity.Name).Scan(&entityID)
	if err != nil {
		return 0, fmt.Errorf("[postgres] failed to create entity: %v (name: %s)", err, entity.Name)
	}

	return entityID, nil
}

func (r *entityRepository) GetEntity(ctx context.Context, id int64) (*model.Entity, error) {
	query := `SELECT name FROM entity WHERE id = $1`

	var entityName string
	err := r.pg.Pool().QueryRow(ctx, query, id).Scan(&entityName)
	if err != nil {
		return nil, fmt.Errorf("[postgres] failed to get entity: %v (id: %d)", err, id)
	}

	return &model.Entity{ID: id, Name: entityName}, nil
}
