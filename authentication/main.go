package main

import (
	"fmt"
	"log"
	"strings"

	"rayquen/vyra/vyramigrationdatabase/google/auth"
)

func main() {
	fmt.Println("Vyra Migration Database - Authentication Service")
	fmt.Println()

	data, err := auth.GetToken()

	if err == nil {

		if data != nil {
			fmt.Println("Token file exist. You want overwrite? y/n")

			var response string
			_, err := fmt.Scan(&response)
			if err != nil {
				log.Fatalf("Unable to read response code: %v", err)
			}

			response = strings.ToLower(response)

			if !(response == "y" || response == "yes") {
				fmt.Println()
				fmt.Println("Bye!")
				return
			}
		}
	}

	fmt.Println("Request authentication...")

	err = auth.RequestToken()

	if err != nil {
		log.Fatalf("Token inaccesible: %v", err)
	}

	fmt.Println("Token file saved succefull")
	fmt.Println()
	fmt.Println("Bye!")
}
