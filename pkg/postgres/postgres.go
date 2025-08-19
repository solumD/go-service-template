package postgres

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	maxPoolSize  = 1
	connAttempts = 10
	connTimeout  = 2 * time.Second
)

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	Pool *pgxpool.Pool
}

func New(dsn string) *Postgres {
	pg := &Postgres{
		maxPoolSize:  maxPoolSize,
		connAttempts: connAttempts,
		connTimeout:  connTimeout,
	}

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize)

	for ; pg.connAttempts > 0; pg.connAttempts-- {
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}

		log.Printf("failed to connect to database, attempts left: %d", pg.connAttempts)

		time.Sleep(pg.connTimeout)

		pg.connAttempts--
	}

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	return pg
}

func (pg *Postgres) Close() {
	if pg.Pool != nil {
		pg.Pool.Close()
	}
}
