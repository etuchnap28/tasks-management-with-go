package persistence

import "github.com/gobuffalo/pop/v6"

type Repository interface {
	~struct{ db *pop.Connection }
}

type atomicOp[R Repository] func(repo *R) error
