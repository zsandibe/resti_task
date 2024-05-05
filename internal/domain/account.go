package domain

import (
	"time"

	"github.com/zsandibe/resti_task/internal/entity"
)

type CreateAccountInput struct {
	Name    string
	Balance float64
}

type GetAccountOutput struct {
	Id           int                  `json:"id"`
	Name         string               `json:"username"`
	Balance      float64              `json:"balance"`
	Created_date time.Time            `json:"created_date"`
	Transactions []entity.Transaction `json:"transactions"`
}
