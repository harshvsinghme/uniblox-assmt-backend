package global

import (
	"log"

	"github.com/joho/godotenv"
)

type globalconfig struct {
	NthOrderForDiscount int
}

type listofsecrets struct {
}

type enum struct {
	GeneralUser string
	AdminUser   string
	AnyUser     string
}

var SECRETS = listofsecrets{}
var CONFIG = globalconfig{}
var ENUM = enum{}

func LoadGlobals(envPath string) {

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	CONFIG.NthOrderForDiscount = 3

	ENUM.GeneralUser = "GENERAL"
	ENUM.AdminUser = "ADMIN"
	ENUM.AnyUser = "ANY"

}
