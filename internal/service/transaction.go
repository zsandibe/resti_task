package service

import (
	"context"

	"github.com/zsandibe/resti_task/internal/domain"
	"github.com/zsandibe/resti_task/internal/entity"
	repo "github.com/zsandibe/resti_task/internal/repository"
)

type TransactionService struct {
	transactionRepo repo.Transaction
}

func NewTransactionService(transactionRepo repo.Transaction) *TransactionService {
	return &TransactionService{transactionRepo: transactionRepo}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, inp domain.CreateTransactionInput) error {
	return s.transactionRepo.CreateTransaction(ctx, inp)
}

func (s *TransactionService) GetAllTransactionsByAccountId(ctx context.Context, id int) ([]entity.Transaction, error) {
	return s.transactionRepo.GetAllTransactionsByAccountId(ctx, id)
}
