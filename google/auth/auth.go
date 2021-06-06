package auth

import (
	"context"
	"net/http"
	"rayquen/vyra/vyramigrationdatabase/google/api"
	"rayquen/vyra/vyramigrationdatabase/google/auth/spreadsheets"

	"golang.org/x/oauth2"
)

func Authenticate() (*api.GoogleAPI, error) {
	cred, err := GetCredential()
	if err != nil {
		return nil, err
	}

	config, err := spreadsheets.GetConfigReadonly(cred)
	if err != nil {
		return nil, err
	}

	client, err := getClient(config)

	if err != nil {
		return nil, err
	}

	api := api.GoogleAPI{Client: client, Service: nil}

	return &api, nil
}

func getClient(config *oauth2.Config) (*http.Client, error) {
	tok, err := GetToken()
	if err != nil {
		return nil, err
	}

	return config.Client(context.Background(), tok), nil
}
