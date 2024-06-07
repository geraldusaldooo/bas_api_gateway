package handler

import (
	"bas_api_gateway/model"
	"bas_api_gateway/utils"
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

	BodyPayloadBank := model.Transaction{}
	err := g.BindJSON(&BodyPayloadBank)

	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()
	defer db.Close()

	result := orm.Create(&BodyPayloadBank)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Transaction created successfully",
		"data":    BodyPayloadBank,
	})
}
