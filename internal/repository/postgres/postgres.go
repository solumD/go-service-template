package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/solumD/go-service-template/internal/model"
	"github.com/solumD/go-service-template/internal/repository"
)

type entityRepository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) repository.Repository {
	return &entityRepository{
		pool: pool,
	}
}

func (r *entityRepository) CreateEntity(ctx context.Context, entity *model.Entity) (int64, error) {
	query := `INSERT INTO entity (name) VALUES ($1) RETURNING id`

	var entityID int64
	err := r.pool.QueryRow(ctx, query, entity.Name).Scan(&entityID)
	if err != nil {
		return 0, fmt.Errorf("[postgres] failed to create entity: %v (name: %s)", err, entity.Name)
	}

	return entityID, nil
}

func (r *entityRepository) GetEntity(ctx context.Context, id int64) (*model.Entity, error) {
	query := `SELECT name FROM entity WHERE id = $1`

	var entityName string
	err := r.pool.QueryRow(ctx, query, id).Scan(&entityName)
	if err != nil {
		return nil, fmt.Errorf("[postgres] failed to get entity: %v (id: %d)", err, id)
	}

	return &model.Entity{ID: id, Name: entityName}, nil
}
