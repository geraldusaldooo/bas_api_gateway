package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionInterface interface {
	TransferBank(*gin.Context)
}

type transactionImplement struct{}

func NewTransaction() TransactionInterface {
	return &transactionImplement{}
}

type BodyPayloadTransaction struct{}

func (b *transactionImplement) TransferBank(g *gin.Context) {

	bodyPayloadTxn := BodyPayloadTransaction{}
	err := g.BindJSON(&bodyPayloadTxn)

	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Hello guys this API rest for later",
	})
}
