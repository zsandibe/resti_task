package entity

import "time"

type Transaction struct {
	Id           int       `json:"id"`
	Value        float64   `json:"value"`
	Account_id   int       `json:"account_id"`
	Group        string    `json:"group"`
	Account2_id  int       `json:"account2_id"`
	Created_date time.Time `json:"created_date"`
}
