package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to setup")
	fmt.Println("Please write your SportsApiKey:")
	err := os.Setenv("SPORTS_API_KEY", GetUserInput())
	if err != nil {
		log.Fatal(err)
	}
}

func GetUserInput() string {
	var key string
	_, err := fmt.Scanln(&key)
	if err != nil {
		log.Fatal(err)
	}

	return key
}
