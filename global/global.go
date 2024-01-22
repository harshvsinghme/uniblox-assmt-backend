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

var SECRETS = listofsecrets{}
var CONFIG = globalconfig{}

func LoadGlobals(envPath string) {

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error Loading .env file")
	}

	CONFIG.NthOrderForDiscount = 3

}
