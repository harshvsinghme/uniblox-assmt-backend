package manager

import "github.com/harshvsinghme/uniblox-assmt-backend/dao"

func AuthenticateUser(email string) string {
	return dao.Authenticate(email)
}

func Logout(userId string) {
	dao.Logout(userId)
}
