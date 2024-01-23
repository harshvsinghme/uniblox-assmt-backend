package dao

import (
	"github.com/harshvsinghme/uniblox-assmt-backend/dbUtils"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
)

func SetCouponCode(code string) error {
	return dbUtils.InMemoryDBClient.SetCouponCode(code)
}

func GetCouponCode() string {
	return dbUtils.InMemoryDBClient.GetCouponCode()
}

func GetAllOrders() []models.Order {
	return dbUtils.InMemoryDBClient.GetAllOrders()
}

func PlaceOrder(newOrder models.Order) {
	dbUtils.InMemoryDBClient.PlaceOrder(newOrder)
}
