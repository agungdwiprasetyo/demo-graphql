package main

import (
	"log"

	"github.com/agungdwiprasetyo/demo-graphql/config"
	env "github.com/joho/godotenv"
)

func main() {
	if err := env.Load(".env"); err != nil {
		log.Fatal(err)
	}

	conf := config.New()
	service := NewService(conf)
	service.ServeHTTP(8080)
}
