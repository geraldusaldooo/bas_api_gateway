package usecase

type Login struct{}

type LoginInterface interface {
	Autentikasi(Username, Password string) bool
}

func NewLogin() LoginInterface {
	return &Login{}
}

func (masuk *Login) Autentikasi(Username, Password string) bool {
	if Username == "admin" && Password == "admin123" {
		return true
	}
	return false
}
