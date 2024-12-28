package repository

import "context"

// SomeRepository interface of something repository
type SomeRepository interface {
	// some repository methods
	SomeMethod(ctx context.Context, someArgs ...interface{}) (interface{}, error)
}
