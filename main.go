package main

import (
	"fmt"
	"log"

	"rayquen/vyra/vyramigrationdatabase/almercadito/clients/services"
	"rayquen/vyra/vyramigrationdatabase/google/api"
	"rayquen/vyra/vyramigrationdatabase/google/auth"
)

var API *api.GoogleAPI

func main() {
	fmt.Println("Vyra Migration Database")

	_, err := auth.GetCredential()
	if err != nil {
		log.Fatalf("[ALERT!!] Not found Credentials file: %v", err)
	}

	fmt.Println("Authenticating...")

	apiConnection, err := auth.Authenticate()

	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	fmt.Println("Get services...")
	API, err := api.GetService(apiConnection)

	fmt.Println("Query clients...")
	fmt.Println()
	err = services.GetAll(API.Service)
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}

	fmt.Println()
	fmt.Println("Stopping service...")
	fmt.Println()
	fmt.Println("Bye!")
}
