package models

type TransferBalanceModel struct {
	ToAccountNumber int `json:"to_account_number"`
	Amount          int `json:"amount"`
}
