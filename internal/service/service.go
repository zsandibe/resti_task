package service

import (
	"context"

	"github.com/zsandibe/resti_task/internal/domain"
	"github.com/zsandibe/resti_task/internal/entity"
	repo "github.com/zsandibe/resti_task/internal/repository"
)

type Account interface {
	CreateAccount(ctx context.Context, inp *domain.CreateAccountInput) (int, error)
	GetAllAccountsWithTransactions(ctx context.Context) ([]domain.GetAccountOutput, error)
}

type Transaction interface {
	CreateTransaction(ctx context.Context, inp domain.CreateTransactionInput) error
	GetAllTransactionsByAccountId(ctx context.Context, id int) ([]entity.Transaction, error)
}

type Service struct {
	Account
	Transaction
}

func NewService(repo *repo.Repository) *Service {
	return &Service{
		Account:     NewAccountService(repo.Account),
		Transaction: NewTransactionService(repo.Transaction),
	}
}
