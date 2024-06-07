package handler

import (
	"bas_api_gateway/model"
	"bas_api_gateway/utils"
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

	accounts := []model.Account{}
	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()
	defer db.Close()

	// q := orm
	// if name !=""{
	// 	q = q.Where("name = ?", name)
	// }
	// result := q.Find(&accounts)

	//result := orm.Where("name=?", name).Find(&accounts)

	result := orm.Find(&accounts, "name = ?", name)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    accounts,
	})
}

// type BodyPayloadAccount struct {
// 	AccountID string
// 	Name      string
// 	Address   string
// }

func (a *accountImplement) CreateAccount(g *gin.Context) {

	BodyPayload := model.Account{}
	err := g.BindJSON(&BodyPayload)

	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()
	defer db.Close()

	result := orm.Create(BodyPayload)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Account created successfully",
		"data":    BodyPayload,
	})
}

func (a *accountImplement) UpdateAccount(g *gin.Context) {

	BodyPayloadUpd := model.Account{}
	err := g.BindJSON(&BodyPayloadUpd)

	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	id := g.Param("id")

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()
	defer db.Close()

	user := model.Account{}
	orm.First(&user, "account_id = ?", id)

	if user.AccountID == "" {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "data not found",
		})
		return
	}

	user.Name = BodyPayloadUpd.Name
	user.Username = BodyPayloadUpd.Username
	user.Password = BodyPayloadUpd.Password
	orm.Save(user)

	g.JSON(http.StatusOK, gin.H{
		"message": "Account updated successfully",
		"data":    BodyPayloadUpd,
	})
}

func (a *accountImplement) RemoveAccount(g *gin.Context) {

	id := g.Param("id")

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()
	defer db.Close()

	result := orm.Where("account_id = ?", id).Delete(&model.Account{})
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Account removed successfully",
		"data":    id,
	})
}

type BodyPayloadBalance struct {
	Account_ID string
	Month      int
}

func (a *accountImplement) GetBalance(g *gin.Context) {

	BodyPayloadBalance := BodyPayloadBalance{}
	err := g.BindJSON(&BodyPayloadBalance)

	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()
	defer db.Close()

	sumResult := struct {
		Total int
	}{}

	result := orm.Model(&model.Transaction{}).
		Select("sum(amount) as total").Where("account_id = ? AND date_part( 'Month' , transaction_date) = ?", BodyPayloadBalance.Account_ID, BodyPayloadBalance.Month).
		Group("account_id").
		Scan(&sumResult)

	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get total successfully",
		"data":    sumResult,
	})

}
