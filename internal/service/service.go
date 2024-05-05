package service

import (
	repo "github.com/zsandibe/resti_task/internal/repository"
)

type Account interface {
}

type Transaction interface {
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
