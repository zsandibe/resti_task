package repo_errs

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrNotEnoughBalance = errors.New("not enough balance")
)
