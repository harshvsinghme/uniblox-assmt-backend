package dao

import (
	"github.com/harshvsinghme/uniblox-assmt-backend/dbUtils"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
)

func AddItemToMyCart(cartItem models.CartItem) {
	dbUtils.InMemoryDBClient.AddtoUserCart(cartItem)
}

func GetMyCart(userId string) []models.CartItem {
	return dbUtils.InMemoryDBClient.GetMyCart(userId)
}
