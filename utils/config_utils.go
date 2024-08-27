package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path"
)

var devEnv = "environments/.env"

func LoadEnv() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Unable to get working directory")
	}
	err = godotenv.Load(path.Join(wd, devEnv))
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}
