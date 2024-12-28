package transaction

import (
	"context"

	"github.com/solumD/go-service-template/internal/client/db"
	"github.com/solumD/go-service-template/internal/client/db/pg"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

type manager struct {
	db db.Transactor
}

// NewTransactionManager returns new transaction manager
func NewTransactionManager(db db.Transactor) db.TxManager {
	return &manager{
		db: db,
	}
}

// transaction executes a transation
func (m *manager) transaction(ctx context.Context, opts pgx.TxOptions, fn db.Handler) (err error) {
	// if it is a nested transaction, we skip the
	// initiation of the new transaction and execute the handler.
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	// starting transaction
	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}

	// adding transaction in context
	ctx = pg.MakeContextTx(ctx, tx)

	// defer to commit transaction or rollback
	// through recover() if there was an error
	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("recovered after panic: %v", err)
		}

		if err != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				err = errors.Wrapf(err, "errRollback: %v", errRollback)
			}

			return
		}

		// committing transaction
		if err == nil {
			err = tx.Commit(ctx)
			if err != nil {
				err = errors.Wrap(err, "failed to commit transaction")
			}
		}
	}()

	// executing query inside transaction
	if err = fn(ctx); err != nil {
		err = errors.Wrap(err, "failed to execute code inside transaction")
	}

	return err
}

// ReadCommited executes transaction with Read-Commited isolation level
func (m *manager) ReadCommitted(ctx context.Context, f db.Handler) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}

	return m.transaction(ctx, txOpts, f)
}
