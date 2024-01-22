package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	manager "github.com/harshvsinghme/uniblox-assmt-backend/business"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
)

type ProductService struct {
}

func (service ProductService) GetAllProducts(ctx *gin.Context) {
	errorOut := models.Error{}

	products := manager.GetAllProducts()

	ctx.JSON(http.StatusOK, gin.H{
		"errorOut": errorOut,
		"products": products,
	})
}
