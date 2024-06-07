package main

import (
	"bas_api_gateway/handler"
	"bas_api_gateway/proto"
	"context"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/client/grpc"
	micro "go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		ExposeHeaders: []string{"*"},
	}))

	addressServiceTransactionOpt := client.WithAddress(":8084")
	clientServiceTransaction := grpc.NewClient()

	serviceTransaction := micro.NewService(
		micro.Client(clientServiceTransaction),
	)

	serviceTransaction.Init(
		micro.Name("service-transaction"),
		micro.Version("latest"),
	)

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/remove/:id", handler.NewAccount().RemoveAccount)
	accountRoute.POST("/getbalance", handler.NewAccount().GetBalance)

	transferRoute := r.Group("/transfer")
	transferRoute.POST("/transferbank", handler.NewTransaction().TransferBank)

	authRoute := r.Group("/auth")
	authRoute.POST("/login", handler.NewAuth().Login)

	transactionRoute := r.Group("/transaction")
	transactionRoute.GET("/get", func(g *gin.Context) {
		clientResponse, err := proto.NewServiceTransactionService("service-transaction", serviceTransaction.Client()).Login(context.Background(), &proto.LoginRequest{
			Username: "aldo",
			Password: "1234567",
		}, addressServiceTransactionOpt)

		if err != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		g.JSON(http.StatusOK, gin.H{
			"data": clientResponse,
		})
	})

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
