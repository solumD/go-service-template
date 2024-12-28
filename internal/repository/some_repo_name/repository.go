package somereponame

import (
	"context"

	"github.com/solumD/go-service-template/internal/client/db"
	"github.com/solumD/go-service-template/internal/repository"
)

const (
// columns' names
)

type repo struct {
	db db.Client
}

// New returns new repository object
func New(db db.Client) repository.SomeRepository {
	return &repo{
		db: db,
	}
}

// SomeMethod ...
func (r *repo) SomeMethod(_ context.Context, _ ...interface{}) (interface{}, error) {
	// something queries
	return struct{}{}, nil
}
