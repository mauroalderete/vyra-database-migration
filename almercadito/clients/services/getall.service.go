package services

import (
	"fmt"

	"google.golang.org/api/sheets/v4"
)

func GetAll(srv *sheets.Service) error {
	readRange := spreadsheet_page + "!A1:I24"

	resp, err := srv.Spreadsheets.Values.Get(spreadsheet_id, readRange).Do()

	if err != nil {
		return err
	}

	if len(resp.Values) == 0 {
		fmt.Println("Ups! sin datos")
	} else {
		for _, row := range resp.Values {
			fmt.Printf("%s: %s\n", row[0], row[2])
		}
	}

	return nil
}
