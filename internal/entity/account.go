package entity

import "time"

type Account struct {
	Id           int       `json:"id"`
	Name         string    `json:"username"`
	Balance      float64   `json:"balance"`
	Created_date time.Time `json:"created_date"`
}
