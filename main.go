package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-crawler/src/utilities"
)

func main() {
	err := godotenv.Load(utilities.ENV_FILE)
	if err != nil {
		fmt.Println("Error loading environment file ", err.Error())
	}

	//other initialization...
}
