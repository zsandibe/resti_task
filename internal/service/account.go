package service

import (
	"context"

	"github.com/zsandibe/resti_task/internal/domain"
	repo "github.com/zsandibe/resti_task/internal/repository"
)

type AccountService struct {
	accountRepo repo.Account
}

func NewAccountService(accountRepo repo.Account) *AccountService {
	return &AccountService{accountRepo: accountRepo}
}

func (s *AccountService) CreateAccount(ctx context.Context, account *domain.CreateAccountInput) (int, error) {
	return s.accountRepo.CreateAccount(ctx, account)
}

func (s *AccountService) GetAllAccountsWithTransactions(ctx context.Context) ([]domain.GetAccountOutput, error) {
	return s.accountRepo.GetAllAccountsWithTransactions(ctx)
}
