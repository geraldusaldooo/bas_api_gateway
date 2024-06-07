package model

import "time"

type Transaction struct {
	ID              int        `gorm:"primarykey"`
	AccountID       string     `gorm: "column:account_id"`
	BankID          string     `gorm: "column:bank_id"`
	Amount          int        `gorm: "column:amount"`
	TransactionDate *time.Time `gorm: "column:transaction_date"`
}

func (a *Transaction) TableName() string {
	return "transaction"
}
