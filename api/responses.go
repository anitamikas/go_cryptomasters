package api

type CEXResponse struct {
	Timestamp             string  `json:"timestamp"`
	Low                   string  `json:"low"` // mapping definition between a json and go structure
	High                  string  `json:"high"`
	Last                  string  `json:"last"`
	Volume                string  `json:"volume"`
	Volume30D             string  `json:"volume30d"`
	Bid                   float64 `json:"bid"`
	Ask                   float64 `json:"ask"`
	PriceChange           string  `json:"priceChange"`
	PriceChangePercentage string  `json:"priceChangePercentage"`
	Pair                  string  `json:"pair"`
}

// https://mholt.github.io/json-to-go/
