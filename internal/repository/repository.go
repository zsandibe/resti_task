package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zsandibe/resti_task/internal/repository/postgres"
)

type Account interface {
}

type Transaction interface {
}

type Repository struct {
	Account
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Account:     postgres.NewAccountRepo(db),
		Transaction: postgres.NewTransactionRepo(db),
	}
}
