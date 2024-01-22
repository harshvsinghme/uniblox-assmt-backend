package dao

import "github.com/harshvsinghme/uniblox-assmt-backend/dbUtils"

func Authenticate(email string) string {
	return dbUtils.InMemoryDBClient.Authenticate(email)
}
