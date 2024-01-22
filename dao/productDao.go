package dao

import (
	"github.com/harshvsinghme/uniblox-assmt-backend/dbUtils"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
)

func GetAllProducts() []models.Product {
	products := dbUtils.InMemoryDBClient.GetAllProducts()
	return products
}
