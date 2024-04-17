package main

import (
	"fmt"

	"github.com/AllenOps08/TranslatorBot/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("There is something wrong while reading the file")
		return
	}
	router.Routing()
}