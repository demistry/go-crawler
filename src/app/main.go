package main

import (
	"fmt"
	"github.com/Demistry/MovieCrawler/src/utilities"
	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load(utilities.ENV_FILE)
	if err != nil{
		fmt.Println("Error loading environment file ", err.Error())
	}

	//other initialization...
}
