package manager

import (
	"github.com/harshvsinghme/uniblox-assmt-backend/dao"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
)

func AddItemToMyCart(userId, ProdId string) {

	//Validate Product existence in db
	product, found := dao.GetProductById(ProdId)
	if !found {
		return
	}

	cartItem := models.CartItem{
		UserId:   userId,
		Product:  product,
		Quantity: 1,
	}
	dao.AddItemToMyCart(cartItem)
}

func GetMyCart(userId string) []models.CartItem {
	return dao.GetMyCart(userId)
}
