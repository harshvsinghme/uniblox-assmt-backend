package models

type User struct {
	Id, Email string
}

type Product struct {
	ProdId, ProdName string
	ProdPrice        float64
}
type OrderItem struct {
	Product
	Quantity int
}
type Order struct {
	OrdId, UserId, CouponCode string
	TotalAmount, TotalPayable float64
	OrderedItems              []OrderItem
}

type CartItem struct {
	Product
	UserId   string
	Quantity int
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	TraceId string `json:"traceId,omitempty"`
}

type OrderSummary struct {
	ItemsPurchased                           int
	TotalPurchaseAmount, TotalDiscountAmount float64
	DiscountCodes                            []string
}
