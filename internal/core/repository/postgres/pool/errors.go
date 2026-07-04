package core_postgres_pool

import "errors"

var (
	ErrNoRows             = errors.New("no rows")
	ErrUnknown            = errors.New("unknown")
	ErrViolatesForeignKey = errors.New("violates foreign key")
)
