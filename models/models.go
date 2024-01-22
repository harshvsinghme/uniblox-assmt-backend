package models

type User struct {
	Id, Email string
}

type Product struct {
	Id, Name string
	Price    float64
}

type Order struct {
	Id, UserId string
	Items      []Product
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	TraceId string `json:"traceId,omitempty"`
}
