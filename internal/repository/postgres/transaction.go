package postgres

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/zsandibe/resti_task/internal/domain"
	"github.com/zsandibe/resti_task/internal/entity"
	"github.com/zsandibe/resti_task/internal/repository/repo_errs"
	"golang.org/x/net/context"
)

type TransactionRepo struct {
	db *sqlx.DB
}

func NewTransactionRepo(db *sqlx.DB) *TransactionRepo {
	return &TransactionRepo{
		db: db,
	}
}

func (r *TransactionRepo) CreateTransaction(ctx context.Context, inp domain.CreateTransactionInput) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `INSERT INTO transactions (value, account_id, "group", account2_id, created_at)
              VALUES ($1, $2, $3, $4, NOW())`
	_, err = tx.ExecContext(ctx, query, inp.Value, inp.Account_id, inp.Group, inp.Account2_id)
	if err != nil {
		return err
	}

	var currentBalance float64
	err = tx.QueryRowContext(ctx, "SELECT balance FROM accounts WHERE id = $1", inp.Account_id).Scan(&currentBalance)
	if err != nil {
		return err
	}

	switch inp.Group {
	case "transfer":
		if inp.Account2_id == 0 {
			return errors.New("account2_id is required for transfer transactions")
		}

		if currentBalance < inp.Value {
			return repo_errs.ErrNotEnoughBalance
		}

		_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - $1 WHERE id = $2 AND balance >= $1", inp.Value, inp.Account_id)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + $1 WHERE id = $2", inp.Value, inp.Account2_id)
		if err != nil {
			return err
		}

	case "income":

		_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance + $1 WHERE id = $2", inp.Value, inp.Account_id)
		if err != nil {
			return err
		}

	case "outcome":
		if currentBalance < inp.Value {
			return repo_errs.ErrNotEnoughBalance
		}

		_, err = tx.ExecContext(ctx, "UPDATE accounts SET balance = balance - $1 WHERE id = $2 AND balance >= $1", inp.Value, inp.Account_id)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *TransactionRepo) GetAllTransactionsByAccountId(ctx context.Context, id int) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	query := `
		SELECT transactions.id,transactions.value,transactions.account_id,
		transactions.group,transactions.account2_id,transactions.created_at
		FROM transactions WHERE account_id = $1
	`

	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction entity.Transaction
		if err := rows.Scan(&transaction.Id, &transaction.Value, &transaction.Account_id, &transaction.Group, &transaction.Account2_id, &transaction.Created_date); err != nil {
			return transactions, err
		}

		transactions = append(transactions, transaction)
	}
	return transactions, nil
}
