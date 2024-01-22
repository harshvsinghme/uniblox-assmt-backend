package dao

import (
	"github.com/harshvsinghme/uniblox-assmt-backend/dbUtils"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
)

func GetAllProducts() []models.Product {
	products := dbUtils.InMemoryDBClient.GetAllProducts()
	return products
}

func GetProductById(ProdId string) (models.Product, bool) {
	return dbUtils.InMemoryDBClient.GetProductById(ProdId)

}
