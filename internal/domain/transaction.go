package domain

type CreateTransactionInput struct {
	Value       float64 `json:"value"`
	Account_id  int     `json:"account_id"`
	Group       string  `json:"group"`
	Account2_id int     `json:"account2_id,omitempty"`
}

type GetTransactionsInput struct {
	Id int `json:"id"`
}
