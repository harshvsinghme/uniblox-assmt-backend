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
	Authenticate(email string) string
	IsAuthenticated(userId string) bool
}

type InMemoryDB struct {
	IDB
}

var InMemoryDBClient InMemoryDB

var users []models.User
var products []models.Product

var orders []models.Order

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

	orders = []models.Order{}
}

func (InMemoryDBClient *InMemoryDB) GetAllProducts() []models.Product {

	return products
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
