package dbUtils

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/harshvsinghme/uniblox-assmt-backend/models"
)

type IDB interface {
	GetAllProducts() ([]models.Product, error)
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

func (InMemoryDBClient *InMemoryDB) GetAllProducts() ([]models.Product, error) {
	var err error

	return products, err
}
