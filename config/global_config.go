package config

import (
	"api_book/helpers"
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	if err := godotenv.Load(); err != nil {
		helpers.Fatal(err)
	}
}

func Get(k string) string {
	return os.Getenv(k)
}
