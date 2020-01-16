package main

import (
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"net/http"
	"strconv"
)

func fetch(url string) ([]Offer, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	xlFile, err := xlsx.OpenBinary(bytes)
	if err != nil {
		return nil, err
	}
	var offers []Offer
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows[1:] {
			cellLen := len(row.Cells)
			if cellLen > 0 {

				id, err := strconv.Atoi(row.Cells[0].String())
				if err != nil {
					return nil, err
				}
				url := row.Cells[cellLen-1].String()
				offers = append(offers, Offer{id, url})
			}
		}
	}
	return offers, nil
}

//Offer represents Cian offer
type Offer struct {
	ID  int
	URL string
}
