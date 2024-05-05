package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/zsandibe/resti_task/internal/domain"
	"github.com/zsandibe/resti_task/internal/entity"
)

type AccountRepo struct {
	db *sqlx.DB
}

func NewAccountRepo(db *sqlx.DB) *AccountRepo {
	return &AccountRepo{
		db: db,
	}
}

func (r *AccountRepo) CreateAccount(ctx context.Context, inp *domain.CreateAccountInput) (int, error) {
	var id int
	query := `
		INSERT INTO accounts (name,balance,created_at)
		VALUES ($1,$2,NOW())
		RETURNING ID
	`

	err := r.db.QueryRowContext(ctx, query, inp.Name, inp.Balance).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AccountRepo) GetAllAccountsWithTransactions(ctx context.Context) ([]domain.GetAccountOutput, error) {
	var accounts []domain.GetAccountOutput

	query := `
        SELECT a.id, a.name, a.balance, a.created_at, 
               t.id AS transaction_id, t.value, t.account_id, t.group, t.account2_id, t.created_at AS transaction_created_at
        FROM accounts a
        LEFT JOIN transactions t ON a.id = t.account_id
        ORDER BY a.id, t.created_at DESC
    `

	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	currentAccountId := 0
	var currentAccount *domain.GetAccountOutput

	for rows.Next() {
		var id int
		var name string
		var balance float64
		var createdDate time.Time
		var transactionId, account2Id, accountId sql.NullInt64
		var value sql.NullFloat64
		var group sql.NullString
		var transactionCreatedAt sql.NullTime

		err = rows.Scan(&id, &name, &balance, &createdDate,
			&transactionId, &value, &accountId, &group, &account2Id, &transactionCreatedAt)
		if err != nil {
			return nil, err
		}

		if currentAccountId != id {
			if currentAccount != nil {
				accounts = append(accounts, *currentAccount)
			}
			currentAccount = &domain.GetAccountOutput{
				Id:           id,
				Name:         name,
				Balance:      balance,
				Created_date: createdDate,
				Transactions: []entity.Transaction{},
			}
			currentAccountId = id
		}

		if transactionId.Valid {
			trans := entity.Transaction{
				Id: int(transactionId.Int64),
			}
			if value.Valid {
				trans.Value = value.Float64
			}
			if accountId.Valid {
				trans.Account_id = int(accountId.Int64)
			}
			if group.Valid {
				trans.Group = group.String
			}
			if account2Id.Valid {
				trans.Account2_id = int(account2Id.Int64)
			}
			if transactionCreatedAt.Valid {
				trans.Created_date = transactionCreatedAt.Time
			}

			currentAccount.Transactions = append(currentAccount.Transactions, trans)
		}
	}
	if currentAccount != nil {
		accounts = append(accounts, *currentAccount)
	}

	return accounts, nil

}
