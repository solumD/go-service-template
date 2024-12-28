package service

import "context"

// SomeService interface of something service
type SomeService interface {
	// some service methods
	SomeMethod(ctx context.Context, someArgs ...interface{}) (interface{}, error)
}
