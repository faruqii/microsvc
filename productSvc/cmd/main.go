package main

import (
	"fmt"
	"product-service/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	app.Run()
}
