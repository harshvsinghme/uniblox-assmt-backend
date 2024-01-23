package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	manager "github.com/harshvsinghme/uniblox-assmt-backend/business"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
	"github.com/harshvsinghme/uniblox-assmt-backend/utils"
)

type OrderService struct {
}

func (service OrderService) SetCouponCode(ctx *gin.Context) {
	errorOut := models.Error{}
	zeroError := models.Error{}

	type ReqBody struct {
		CouponCode string
	}

	var reqBody ReqBody

	if err := ctx.BindJSON(&reqBody); err != nil {
		log.Println(err)
		utils.GetError(&errorOut, "invalid request body", 400)
		ctx.JSON(http.StatusBadRequest, gin.H{"errorOut": errorOut})
		return
	}

	errorOut = manager.SetCouponCode(reqBody.CouponCode)
	if errorOut != zeroError {
		log.Println(errorOut.Message)
		ctx.JSON(http.StatusInternalServerError, gin.H{"errorOut": errorOut})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"errorOut": errorOut,
	})
}

func (service OrderService) Checkout(ctx *gin.Context) {
	errorOut := models.Error{}

	userId := utils.GetValueFromContext(ctx, "userId")

	type ReqBody struct {
		CouponCode string
	}

	var reqBody ReqBody

	if err := ctx.BindJSON(&reqBody); err != nil {
		log.Println(err)
		utils.GetError(&errorOut, "invalid request body", 400)
		ctx.JSON(http.StatusBadRequest, gin.H{"errorOut": errorOut})
		return
	}

	manager.Checkout(userId, reqBody.CouponCode)

	ctx.JSON(http.StatusOK, gin.H{
		"errorOut": errorOut,
	})
}

func (service OrderService) OrdersSummary(ctx *gin.Context) {
	errorOut := models.Error{}

	var orderSummary models.OrderSummary

	orderSummary = manager.OrdersSummary()

	ctx.JSON(http.StatusOK, gin.H{
		"errorOut":     errorOut,
		"orderSummary": orderSummary,
	})
}
