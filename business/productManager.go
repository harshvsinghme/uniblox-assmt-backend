package manager

import (
	"github.com/harshvsinghme/uniblox-assmt-backend/dao"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
)

func GetAllProducts() []models.Product {
	return dao.GetAllProducts()
}
