package price

import (
	//"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// instruments USD_JPY
type Get_rate struct {
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
type Get_price struct {
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
type Get_candles struct {
	Instument   string         `json:"instrument"`
	Granularity string         `json:"granularity"`
	Candles     []Candles_rate `json:"candles"`
}
type Candles_rate struct {
	Time     string `json:"time"`
	Openbid  int    `json:"openBid"`
	Openask  int    `json:"openAsk"`
	Highbid  int    `json:"highBid"`
	Highask  int    `json:"highAsk"`
	LowBid   int    `json:"lowbid"`
	Closebid int    `json:"closeBid"`
	Closeask int    `json:"closeAsk"`
	Volume   int    `json:"volume"`
	Complete bool   `json:"complete"`
}

func Get_current_rate(End_point string, Token string) []byte {
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

func Get_history_rate(End_point string, Token string) []byte {
	//var
	client := &http.Client{}
	req, err := http.NewRequest("GET", End_point+"/v1/candles?accountId=8307075&instrument=USD_JPY&granuarity=H4&count=2", nil)

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
