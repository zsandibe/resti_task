package pkg

import (
	"errors"
	"fmt"
	"unicode"
)

var ErrIncorrectGroup = errors.New("incorrect group")

func ValidateAccount(name string, balance float64) error {
	if err := validateOnlyLetters(name); err != nil {
		return err
	}
	if err := validateBalance(balance); err != nil {
		return err
	}
	return nil
}

func ValidateTransaction(balance float64, group string, acc_id int, acc2_id int) error {
	if err := validateBalance(balance); err != nil {
		return err
	}

	if acc_id == 0 {
		return errors.New("acc_id can`t be 0")
	}

	if group == "income" || group == "outcome" {
		return nil
	} else if group == "transfer" {
		fmt.Println(acc2_id)
		if acc2_id == 0 {
			return errors.New("acc2_id can`t be 0")
		}
		return nil
	}
	return ErrIncorrectGroup
}

func validateOnlyLetters(s string) error {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return errors.New("name must contain only letters")
		}
	}
	if s == "" {
		return errors.New("name is required and cannot be empty")
	}
	return nil
}

func validateBalance(b float64) error {
	if b <= 0 {
		return errors.New("balance must be greater than 0")
	}
	return nil
}
