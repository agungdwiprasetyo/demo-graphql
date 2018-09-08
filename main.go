package main

import (
	"log"

	env "github.com/joho/godotenv"
)

func main() {
	if err := env.Load(".env"); err != nil {
		log.Fatal(err)
	}
	service := NewService()
	service.ServeHTTP(8080)
}
