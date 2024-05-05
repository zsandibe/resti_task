package service

import (
	repo "github.com/zsandibe/resti_task/internal/repository"
)

type TransactionService struct {
	transactionRepo repo.Transaction
}

func NewTransactionService(transactionRepo repo.Transaction) *TransactionService {
	return &TransactionService{transactionRepo: transactionRepo}
}
