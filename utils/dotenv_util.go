package utils

import "github.com/joho/godotenv"

func LoadDotenv() {
	err := godotenv.Load()

	if err != nil {
		panic(nil)
	}
}
