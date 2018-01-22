package price

import (
	//"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Get_rate struct {
	Instruments []Instruments_rate `json:"instruments"`
}
type Instruments_rate struct {
	Instrument    string `json:"instrument"`
	Displayname   string `json:"displayName"`
	Pip           string `json:"pip"`
	Maxtradeunits int    `json:"maxTradeUnits"`
}

func Get_current_rate(End_point string, Token string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest("GET", End_point+"/v1/prices?instruments=USD_JPY", nil)
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
