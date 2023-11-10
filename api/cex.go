package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"mikasanita.com/go/cryptomasters/datatypes"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characters required, got %d", len(currency))
	}
	upCurrency := strings.ToUpper(currency) // returning a new string
	res, err := http.Get(fmt.Sprintf(apiUrl, upCurrency))
	if err != nil {
		return nil, err
	}
	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body) // the thread will wait there until the body is ready cause it comes in chunks
		if err != nil {
			return nil, err
		} // in case the connection was closed before the body was ready, example our internet connection was lost
		// json := string(bodyBytes)
		err = json.Unmarshal(bodyBytes, &response)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, fmt.Errorf("received status code: %v", res.StatusCode)
	}
	rate := &datatypes.Rate{Currency: currency, Price: response.Bid}
	return rate, nil
}
