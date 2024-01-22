package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
	"github.com/harshvsinghme/uniblox-assmt-backend/services"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	AuthService := services.AuthService{}

	router.GET("/status", func(ctx *gin.Context) {
		errorOut := models.Error{}
		ctx.JSON(http.StatusOK, gin.H{"message": "server is running", "errorOut": errorOut})
	})

	router.POST("/auth-service/authenticate", AuthService.AuthenticateUser)

	return router
}
