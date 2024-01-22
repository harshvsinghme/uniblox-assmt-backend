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
