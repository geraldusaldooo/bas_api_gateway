package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountInterface interface {
	GetAccount(*gin.Context)
	CreateAccount(*gin.Context)
	UpdateAccount(*gin.Context)
	RemoveAccount(*gin.Context)
	GetBalance(*gin.Context)
}

type accountImplement struct{}

func NewAccount() AccountInterface {
	return &accountImplement{}
}

func (a *accountImplement) GetAccount(g *gin.Context) {

	QueryParam := g.Request.URL.Query()

	name := QueryParam.Get("name")

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    name,
	})
}

type BodyPayloadAccount struct {
	AccountID string
	Name      string
	Address   string
}

func (a *accountImplement) CreateAccount(g *gin.Context) {

	BodyPayload := BodyPayloadAccount{}
	err := g.BindJSON(&BodyPayload)

	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    BodyPayload,
	})
}

func (a *accountImplement) UpdateAccount(g *gin.Context) {

	QueryParam := g.Request.URL.Query()

	name := QueryParam.Get("name")

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    name,
	})
}

func (a *accountImplement) RemoveAccount(g *gin.Context) {

	id := g.Param("id")

	g.JSON(http.StatusOK, gin.H{
		"message": "Account removed successfully",
		"data":    id,
	})
}

type BodyPayloadBalance struct{}

func (a *accountImplement) GetBalance(g *gin.Context) {

	bodyPayloadBal := BodyPayloadBalance{}
	err := g.BindJSON(&bodyPayloadBal)

	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Hello guys this API rest for later",
	})
}
