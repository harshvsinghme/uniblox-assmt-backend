package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	manager "github.com/harshvsinghme/uniblox-assmt-backend/business"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
	"github.com/harshvsinghme/uniblox-assmt-backend/utils"
)

type UserService struct {
}

func (service UserService) AddItemToMyCart(ctx *gin.Context) {
	errorOut := models.Error{}

	userId := utils.GetValueFromContext(ctx, "userId")

	type ReqBody struct {
		ProdId string
	}

	var reqBody ReqBody

	if err := ctx.BindJSON(&reqBody); err != nil {
		log.Println(err)
		utils.GetError(&errorOut, "invalid request body", 400)
		ctx.JSON(http.StatusBadRequest, gin.H{"errorOut": errorOut})
		return
	}

	manager.AddItemToMyCart(userId, reqBody.ProdId)

	ctx.JSON(http.StatusOK, gin.H{
		"errorOut": errorOut,
	})
}

func (service UserService) GetMyCart(ctx *gin.Context) {
	errorOut := models.Error{}
	userId := utils.GetValueFromContext(ctx, "userId")

	cart := manager.GetMyCart(userId)

	ctx.JSON(http.StatusOK, gin.H{
		"errorOut": errorOut,
		"cart":     cart,
	})
}
