package manager

import (
	"net/http"

	"github.com/harshvsinghme/uniblox-assmt-backend/dao"
	"github.com/harshvsinghme/uniblox-assmt-backend/global"
	"github.com/harshvsinghme/uniblox-assmt-backend/models"
	"github.com/harshvsinghme/uniblox-assmt-backend/utils"
	"github.com/twinj/uuid"
)

func SetCouponCode(code string) (Err models.Error) {

	err := dao.SetCouponCode(code)
	if err != nil {
		utils.GetError(&Err, err.Error(), http.StatusInternalServerError)
	}
	return
}

func Checkout(userId, CouponCode string) {

	//Overall orders of all users in database
	orders := dao.GetAllOrders()
	ordersCount := len(orders)

	activeCouponCode := dao.GetCouponCode()

	//Logged in user's cart
	cart := dao.GetMyCart(userId)
	//Logged in user's cart items, which will be checked out and an order will be placed for
	var orderedItems []models.OrderItem
	prepareOrderedItems(&cart, &orderedItems)

	newOrder := models.Order{
		OrdId:        uuid.NewV4().String(),
		UserId:       userId,
		OrderedItems: orderedItems,
	}

	newOrder.TotalAmount = calculateTotalAmount(&orderedItems)
	newOrder.TotalPayable = newOrder.TotalAmount

	var atleastOneOrder = ordersCount >= 1
	var isDiscApplicable = (ordersCount % global.CONFIG.NthOrderForDiscount) == 0
	var validCouponCode = activeCouponCode == CouponCode

	if atleastOneOrder && isDiscApplicable && validCouponCode {

		newOrder.CouponCode = CouponCode

		newOrder.TotalPayable = newOrder.TotalAmount - (.1 * newOrder.TotalAmount)
	}

	dao.PlaceOrder(newOrder)

}

func prepareOrderedItems(cart *[]models.CartItem, orderedItems *[]models.OrderItem) {
	var item models.OrderItem
	for _, val := range *cart {
		item.Product = val.Product
		item.Quantity = val.Quantity

		*orderedItems = append(*orderedItems, item)
	}
}

func calculateTotalAmount(orderedItems *[]models.OrderItem) (total float64) {

	for _, val := range *orderedItems {
		total += val.ProdPrice * float64(val.Quantity)
	}

	return total
}
