package model

type Bank struct {
	BankCode string `gorm:"primarykey"`
	Name     string `gorm: "column:name"`
	Address  string `gorm: "column:address"`
}

func (a *Bank) TableName() string {
	return "bank"
}
