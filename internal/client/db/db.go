package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Handler function which runs inside transaction
type Handler func(ctx context.Context) error

// Client interface of database client
type Client interface {
	DB() DB
	Close() error
}

// TxManager interface of a transaction manager that executes the specified handler in a transaction
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

// Query wrapper above the query, storing the query name and the query itself
// Query name is used for logging and potentially for tracing
type Query struct {
	Name     string
	QueryRaw string
}

// Transactor interface of transaction beginner
type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

// SQLExecer interface which combines NamedExecer and QueryExecer
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// NamedExecer interface for working with named queries using tags in structures
type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// QueryExecer interface for working with common requests
type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// Pinger interface for checking the connection to the database
type Pinger interface {
	Ping(ctx context.Context) error
}

// DB database interface
type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}
