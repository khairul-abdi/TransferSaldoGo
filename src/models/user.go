package models

type CustomerModel struct {
	CustomerNumber int    `json:"customer_number" bson:"customer_number"`
	CustomerName   string `json:"customer_name" bson:"customer_name"`
}

type AccountModel struct {
	AccountNumber  int `json:"account_number" bson:"account_number"`
	CustomerNumber int `json:"customer_number" bson:"customer_number"`
	Balance        int `json:"balance" bson:"balance"`
}
