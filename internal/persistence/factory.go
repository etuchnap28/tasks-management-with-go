package persistence

import (
	"context"
	"fmt"
	"github.com/gobuffalo/pop/v6"
)

type RepositoryFactory[R Repository] struct {
	db      *pop.Connection
	setupFn func(*R)
}

func NewRepositoryFactory[R Repository](db *pop.Connection) *RepositoryFactory[R] {
	var factory RepositoryFactory[R]
	factory = RepositoryFactory[R]{db: db}
	return &factory
}

func (factory *RepositoryFactory[R]) WithTransaction(ctx context.Context, fn atomicOp[R]) error {
	tx, err := factory.db.NewTransactionContext(ctx)
	if err != nil {
		return err
	}
	repo := R{db: tx}
	if factory.setupFn != nil {
		factory.setupFn(&repo)
	}
	if err := fn(&repo); err != nil {
		if rbErr := tx.TX.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction error: %v, rollback error: %v", err, rbErr)
		}
		return err
	}
	return tx.TX.Commit()
}

func (factory *RepositoryFactory[R]) WithSetupFunction(setupFn func(*R)) *RepositoryFactory[R] {
	factory.setupFn = setupFn
	return factory
}
