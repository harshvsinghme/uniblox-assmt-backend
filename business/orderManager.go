package manager

import (
	"net/http"

	"github.com/harshvsinghme/uniblox-assmt-backend/dao"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
	"github.com/harshvsinghme/uniblox-assmt-backend/utils"
)

func SetCouponCode(code string) (Err models.Error) {

	err := dao.SetCouponCode(code)
	if err != nil {
		utils.GetError(&Err, err.Error(), http.StatusInternalServerError)
	}
	return
}
