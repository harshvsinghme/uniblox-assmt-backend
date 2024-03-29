package services

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	manager "github.com/harshvsinghme/uniblox-assmt-backend/business"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
	"github.com/harshvsinghme/uniblox-assmt-backend/utils"
)

type AuthService struct {
}

func (service AuthService) AuthenticateUser(ctx *gin.Context) {
	errorOut := models.Error{}

	type ReqBody struct {
		Email string `json:"email"`
	}

	var reqBody ReqBody

	if err := ctx.BindJSON(&reqBody); err != nil {
		log.Println(err)
		utils.GetError(&errorOut, "invalid request body", 400)
		ctx.JSON(http.StatusBadRequest, gin.H{"errorOut": errorOut})
		return
	}

	userId := manager.AuthenticateUser(reqBody.Email)

	ctx.SetCookie("userId", userId, int(2*time.Hour), "/", "localhost", false, true) // "userId" is the cookie name
	ctx.JSON(http.StatusOK, gin.H{
		"errorOut": errorOut,
		"userId":   userId,
	})
}

func (service AuthService) Logout(ctx *gin.Context) {
	errorOut := models.Error{}

	userId := utils.GetValueFromContext(ctx, "userId")

	manager.Logout(userId)

	ctx.SetCookie("userId", "", -1, "/", "localhost", false, true) // clear cookie

	ctx.JSON(http.StatusOK, gin.H{
		"errorOut": errorOut,
	})
}
