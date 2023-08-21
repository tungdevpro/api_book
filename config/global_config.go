package config

import "github.com/joho/godotenv"

func GlobalConfig() {
	if err := godotenv.Load(); err != nil {

	}
}
