package auth

import (
	"io/ioutil"
	"os"
	"path"
)

const CREDENTIAL_FILENAME = "credentials.json"

func GetCredential() ([]byte, error) {

	credentialPath := os.Getenv(ENV_AUTH_SESSION_PATH)
	credentialPath = path.Clean(credentialPath)

	data, err := ioutil.ReadFile(path.Join(credentialPath, CREDENTIAL_FILENAME))

	return data, err
}
