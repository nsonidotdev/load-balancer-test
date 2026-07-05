package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/nsonidotdev/load-balancer-test/apps/api/internal/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file", err.Error())
	}

	dbConfig := database.GetEnvConfig()
	conn, err := database.Connect(dbConfig)
	if err != nil {
		log.Fatal("Could not connect to the database", err.Error())
	}
	defer conn.Close(context.Background())

	fmt.Println("Connected to the database succesfully")
}
