package main

import (
	"fmt"
	"log"
	"os"

	"rayquen/vyra/vyramigrationdatabase/almercadito/clients/services"
	"rayquen/vyra/vyramigrationdatabase/google/api"
	"rayquen/vyra/vyramigrationdatabase/google/auth"
)

const ENV_MODE_GET_TOKEN_FROM_WEB = "VMDB_GET_TOKEN_FROM_WEB"
const ENV_FORCE_OVERWRITE_TOKEN = "VMDB_FORCE_OVERWRITE_TOKEN"
const GET_TOKEN_FROM_WEB_YES = "true"
const FORCE_OVERWRITE_TOKEN_YES = "true"

var API *api.GoogleAPI

func main() {
	fmt.Println("Vyra Migration Database")

	//
	os.Setenv(ENV_MODE_GET_TOKEN_FROM_WEB, "false")
	os.Setenv(ENV_FORCE_OVERWRITE_TOKEN, "false")
	//

	_, errCredential := auth.GetCredential()
	if errCredential != nil {
		log.Fatalf("[ALERT!!] Not found Credentials file: %v", errCredential)
	}

	if os.Getenv(ENV_MODE_GET_TOKEN_FROM_WEB) == GET_TOKEN_FROM_WEB_YES {
		modeGetWebToken()
	} else {
		modeRunService()
	}

	fmt.Println("VMDB Service is Stopped")
}

func modeGetWebToken() {

	fmt.Println("Running get token from web mode...")

	_, err := auth.GetToken()

	if err == nil {
		if os.Getenv(ENV_FORCE_OVERWRITE_TOKEN) != FORCE_OVERWRITE_TOKEN_YES {
			log.Fatalf("Token file exist. Not be must overwrite")
		}
	}

	fmt.Print("Solicitando autenticacion...")

	err = auth.RequestToken()

	if err != nil {
		log.Fatalf("Token inaccesible: %v", err)
	}
}

func modeRunService() {
	fmt.Println("Running service mode...")

	apiConfig, err := auth.Authenticate()

	if err != nil {
		log.Fatalf("Authentication failed: %v", err)
	}

	API, err := api.GetService(apiConfig)

	fmt.Println("Wake up service...")
	err = services.GetAll(API.Service)
	if err != nil {
		log.Fatalf("Algo salio mal: %v", err)
	}

}
