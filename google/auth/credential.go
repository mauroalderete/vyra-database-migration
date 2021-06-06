package auth

import (
	"io/ioutil"
)

const CREDENTIALS_FILENAME = "credentials.json"

func GetCredential() ([]byte, error) {
	data, err := ioutil.ReadFile(CREDENTIALS_FILENAME)

	return data, err
}
