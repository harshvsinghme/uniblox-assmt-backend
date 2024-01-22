package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshvsinghme/uniblox-assmt-backend/dao"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		errorOut := models.Error{}
		userId := ctx.GetHeader("userId")
		if userId == "" {

			ctx.JSON(http.StatusUnauthorized, gin.H{"errorOut": errorOut})
			ctx.Abort()
			return
		}

		IsAuthenticated := dao.IsAuthenticated(userId)
		if !IsAuthenticated {
			ctx.JSON(http.StatusUnauthorized, gin.H{"errorOut": errorOut})
			ctx.Abort()
			return
		}
		ctx.Set("userId", userId)
		ctx.Next()
	}
}
