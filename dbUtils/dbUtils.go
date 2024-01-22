package dbUtils

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/harshvsinghme/uniblox-assmt-backend/models"
	"github.com/twinj/uuid"
)

type IDB interface {
	GetAllProducts() []models.Product
	GetProductById(ProdId string) (models.Product, bool)

	Authenticate(email string) string
	IsAuthenticated(userId string) bool
	Logout()

	AddtoUserCart(item models.CartItem)
	GetMyCart(userId string) []models.CartItem
}

type InMemoryDB struct {
	IDB
}

var InMemoryDBClient InMemoryDB

var users []models.User
var products []models.Product
var orders []models.Order
var cart []models.CartItem

func InitDB() {
	users = []models.User{}

	products = []models.Product{}
	fileContent, err := ioutil.ReadFile("data/products.json")
	if err == nil {
		err = json.Unmarshal(fileContent, &products)
		if err != nil {
			log.Print("Failed parsing products data", err)
			products = []models.Product{}
		}
	} else {
		log.Println("Error reading products data", err)
	}

	cart = []models.CartItem{}

	orders = []models.Order{}
}

func (InMemoryDBClient *InMemoryDB) GetAllProducts() []models.Product {

	return products
}

func (InMemoryDBClient *InMemoryDB) GetProductById(ProdId string) (currProd models.Product, found bool) {

	for i := range products {
		currProd = products[i]

		if currProd.ProdId == ProdId {
			found = true
			break
		}
	}
	return
}

func (InMemoryDBClient *InMemoryDB) Authenticate(email string) string {

	userRecord := models.User{}
	alreadyExists := false

	for i := range users {
		userRecord = users[i]
		if userRecord.Email == email {
			alreadyExists = true
			break
		}
	}

	if alreadyExists {
		return userRecord.Id
	}

	newUser := models.User{
		Id:    uuid.NewV4().String(),
		Email: email,
	}
	users = append(users, newUser)

	return newUser.Id

}

func (InMemoryDBClient *InMemoryDB) IsAuthenticated(userId string) bool {

	userRecord := models.User{}
	found := false

	for i := range users {
		userRecord = users[i]
		if userRecord.Id == userId {
			found = true
			break
		}
	}

	return found

}

func (InMemoryDBClient *InMemoryDB) Logout(userId string) {

	for i, v := range users {
		if v.Id == userId {
			users = append(users[:i], users[i+1:]...)
			break
		}
	}
	// fmt.Println(users)
}

func (InMemoryDBClient *InMemoryDB) AddtoUserCart(item models.CartItem) {

	var alreadyExists bool

	for i := range cart {

		currItem := cart[i]

		if currItem.ProdId == item.ProdId && currItem.UserId == item.UserId {
			alreadyExists = true
			break
		}
	}

	if !alreadyExists {
		cart = append(cart, item)
	}

}

func (InMemoryDBClient *InMemoryDB) GetMyCart(userId string) (myCart []models.CartItem) {

	for i := range cart {
		currItem := cart[i]

		if currItem.UserId == userId {
			myCart = append(myCart, currItem)
		}
	}
	return
}
