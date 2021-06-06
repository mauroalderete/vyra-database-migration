package api

import (
	"context"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const MASTER_CLIENTS_API = "1BPGEDtDsiHKNfJylUFfEy9esnYY1If6SAKHW82psthA"

type GoogleAPI struct {
	Client  *http.Client
	Service *sheets.Service
}

func GetService(api *GoogleAPI) (*GoogleAPI, error) {
	srv, err := sheets.NewService(context.Background(), option.WithHTTPClient(api.Client))

	if err != nil {
		return api, err
	}

	api.Service = srv

	return api, nil
}
