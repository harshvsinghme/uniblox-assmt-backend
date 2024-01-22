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
		cookie, err := ctx.Request.Cookie("userId")
		// fmt.Println(cookie, err)
		if err != nil {

			ctx.JSON(http.StatusUnauthorized, gin.H{"errorOut": errorOut})
			ctx.Abort()
			return
		}

		userId := cookie.Value
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
