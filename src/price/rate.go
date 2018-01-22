package price

import (
	//"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// instruments USD_JPY
type Show_rate struct {
	Instruments []Instruments_rate `json:"instruments"`
}
type Instruments_rate struct {
	Instrument    string `json:"instrument"`
	Displayname   string `json:"displayName"`
	Pip           string `json:"pip"`
	Maxtradeunits int    `json:"maxTradeUnits"`
}

// Instrument End
// get_price
type Show_price struct {
	Prices []Prices_rate `json:"prices"`
}
type Prices_rate struct {
	Instrument string `json:"instrument"`
	Time       string `json:"time"`
	Bid        int    `json:"bid"`
	Ask        int    `json:"ask"`
}

// get_price end
// get_candles
type Show_candles struct {
	Instument   string         `json:"instrument"`
	Granularity string         `json:"granularity"`
	Candles     []Candles_rate `json:"candles"`
}
type Candles_rate struct {
	Time     string  `json:"time"`
	Openbid  float32 `json:"openBid"`
	Openask  float32 `json:"openAsk"`
	Highbid  float32 `json:"highBid"`
	Highask  float32 `json:"highAsk"`
	LowBid   float32 `json:"lowbid"`
	Closebid float32 `json:"closeBid"`
	Closeask float32 `json:"closeAsk"`
	Volume   int     `json:"volume"`
	Complete bool    `json:"complete"`
}

func Show_current_rate(End_point string, Token string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", End_point+"/v1/prices?accountId=8307075&instruments=USD_JPY", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+Token)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body

}

func Show_history_rate(End_point string, Token string) []byte {
	//var
	client := &http.Client{}
	req, err := http.NewRequest("GET", End_point+"/v1/candles?accountId=8307075&instrument=USD_JPY&granularity=H4&count=2", nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+Token)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body

}
