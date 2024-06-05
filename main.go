package main

import (
	"bas_api_gateway/usecase"
	"fmt"
)

func main() {

	LoginAuth := usecase.NewLogin()

	Username := "admin"
	Password := "admin123"

	if LoginAuth.Autentikasi(Username, Password) {
		fmt.Println("Login berhasil")
	} else {
		fmt.Println("Login gagal")
	}
}
