package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to setup")
	fmt.Println("Please write your SportsApiKey:")
	err := os.Setenv("SportsApiKey", GetUserInput())
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserInput() string {
	var key string
	fmt.Scanln(&key)
	return key
}
