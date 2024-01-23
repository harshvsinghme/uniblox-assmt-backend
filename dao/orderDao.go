package dao

import (
	"github.com/harshvsinghme/uniblox-assmt-backend/dbUtils"
)

func SetCouponCode(code string) error {
	return dbUtils.InMemoryDBClient.SetCouponCode(code)
}

func GetCouponCode() string {
	return dbUtils.InMemoryDBClient.GetCouponCode()
}
