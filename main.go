package main

import (
	"bas_api_gateway/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/remove/:id", handler.NewAccount().RemoveAccount)
	accountRoute.POST("/getbalance", handler.NewAccount().GetBalance)

	transactionRoute := r.Group("/transaction")
	transactionRoute.POST("/transferbank", handler.NewTransaction().TransferBank)

	authRoute := r.Group("/auth")
	authRoute.POST("/login", handler.NewAuth().Login)
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
