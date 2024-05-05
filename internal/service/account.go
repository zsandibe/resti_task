package service

import (
	repo "github.com/zsandibe/resti_task/internal/repository"
)

type AccountService struct {
	accountRepo repo.Account
}

func NewAccountService(accountRepo repo.Account) *AccountService {
	return &AccountService{accountRepo: accountRepo}
}
