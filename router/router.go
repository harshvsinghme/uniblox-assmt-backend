package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshvsinghme/uniblox-assmt-backend/global"
	"github.com/harshvsinghme/uniblox-assmt-backend/middleware"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
	"github.com/harshvsinghme/uniblox-assmt-backend/services"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	AuthService := services.AuthService{}
	UserService := services.UserService{}
	ProductService := services.ProductService{}
	OrderService := services.OrderService{}

	router.GET("/status", func(ctx *gin.Context) {
		errorOut := models.Error{}
		ctx.JSON(http.StatusOK, gin.H{"message": "server is running", "errorOut": errorOut})
	})

	//Auth-Service
	router.POST("/auth-service/users/authenticate", AuthService.AuthenticateUser)
	router.PUT("/auth-service/users/logout", middleware.SessionMiddleware(global.ENUM.AnyUser), AuthService.Logout)

	//User-Service
	router.POST("/user-service/cart/add", middleware.SessionMiddleware(global.ENUM.GeneralUser), UserService.AddItemToMyCart)
	router.GET("/user-service/cart/get", middleware.SessionMiddleware(global.ENUM.GeneralUser), UserService.GetMyCart)

	//Product-Service
	router.GET("/product-service/products/get", middleware.SessionMiddleware(global.ENUM.GeneralUser), ProductService.GetAllProducts)

	//Order-Service
	router.POST("/order-service/coupon/set", middleware.SessionMiddleware(global.ENUM.AdminUser), OrderService.SetCouponCode)
	router.POST("/order-service/cart/checkout", middleware.SessionMiddleware(global.ENUM.GeneralUser), OrderService.Checkout)

	return router
}
