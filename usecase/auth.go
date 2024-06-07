package usecase

import (
	"bas_api_gateway/model"
	"bas_api_gateway/utils"
)

type Login struct{}

type LoginInterface interface {
	Autentikasi(string, string) bool
}

func NewLogin() LoginInterface {
	return &Login{}
}

func (masuk *Login) Autentikasi(username, password string) bool {

	accounts := model.Account{}
	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()
	defer db.Close()

	orm.Find(&accounts, "username = ? AND password = ? ", username, password)
	if accounts.AccountID == "" {

		return false
	}

	return true

	// if Username == "admin" && Password == "admin123" {
	// 	return true
	// }
	// return false
}
