package model

type Account struct {
	AccountID string `gorm:"primarykey"`
	Username  string `gorm: "column:username"`
	Password  string `gorm: "column:password"`
	Name      string `gorm: "column:name"`
}

func (a *Account) TableName() string {
	return "account"
}
