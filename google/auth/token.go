package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"rayquen/vyra/vyramigrationdatabase/google/auth/spreadsheets"

	"golang.org/x/oauth2"
)

const TOKEN_FILENAME = "token.json"

func GetToken() (*oauth2.Token, error) {

	tokenPath := os.Getenv(ENV_AUTH_SESSION_PATH)
	tokenPath = path.Clean(tokenPath)

	f, err := os.Open(path.Join(tokenPath, TOKEN_FILENAME))
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

func RequestToken() error {

	data, err := GetCredential()

	if err != nil {
		return err
	}

	config, err := spreadsheets.GetConfigReadonly(data)

	if err != nil {
		return err
	}

	token := getTokenFromWeb(config)

	tokenPath := os.Getenv(ENV_AUTH_SESSION_PATH)
	tokenPath = path.Clean(tokenPath)

	saveToken(path.Join(tokenPath, TOKEN_FILENAME), token)

	return nil
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
