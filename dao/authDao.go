package dao

import "github.com/harshvsinghme/uniblox-assmt-backend/dbUtils"

func Authenticate(email string) string {
	return dbUtils.InMemoryDBClient.Authenticate(email)
}

func IsAuthenticated(userId string, userType string) int {
	return dbUtils.InMemoryDBClient.IsAuthenticated(userId, userType)
}

func Logout(userId string) {
	dbUtils.InMemoryDBClient.Logout(userId)
}
