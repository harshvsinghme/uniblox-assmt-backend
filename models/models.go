package models

type User struct {
	Id, Email string
}

type Product struct {
	ProdId, ProdName string
	ProdPrice        float64
}

type Order struct {
	OrdId, UserId string
	OrderedItems  []Product
}

type CartItem struct {
	Product
	UserId string
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	TraceId string `json:"traceId,omitempty"`
}
