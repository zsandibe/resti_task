package postgres

import "github.com/jmoiron/sqlx"

type TransactionRepo struct {
	db *sqlx.DB
}

func NewTransactionRepo(db *sqlx.DB) *TransactionRepo {
	return &TransactionRepo{
		db: db,
	}
}
