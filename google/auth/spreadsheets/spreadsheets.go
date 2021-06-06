package spreadsheets

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const AUTH_SPREADSHEETS_READONLY = "https://www.googleapis.com/auth/spreadsheets.readonly"

func GetConfigReadonly(b []byte) (*oauth2.Config, error) {
	return getConfig(b, AUTH_SPREADSHEETS_READONLY)
}

func getConfig(b []byte, url string) (*oauth2.Config, error) {

	config, err := google.ConfigFromJSON(b, AUTH_SPREADSHEETS_READONLY)
	if err != nil {
		return nil, err
	}

	return config, nil
}
