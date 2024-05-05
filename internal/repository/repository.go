package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/zsandibe/resti_task/internal/domain"
	"github.com/zsandibe/resti_task/internal/entity"
	"github.com/zsandibe/resti_task/internal/repository/postgres"
)

type Account interface {
	CreateAccount(ctx context.Context, account *domain.CreateAccountInput) (int, error)
	GetAllAccountsWithTransactions(ctx context.Context) ([]domain.GetAccountOutput, error)
}

type Transaction interface {
	GetAllTransactionsByAccountId(ctx context.Context, id int) ([]entity.Transaction, error)
	CreateTransaction(ctx context.Context, input domain.CreateTransactionInput) error
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
