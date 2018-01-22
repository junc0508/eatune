package compute

import (
	"github.com/junc0508/eatune/src/price"
)

func Calculate_history_data(price.Show_candles) {
	var judg bool
	var judg_act bool
	if history.Candles[1].Volume > 5000 {
		if history.Candles[1].Closebid-history.Candles[1].Openbid > 0 {
			jugd_act = true
		} else {
			jugd_act = false
		}
		judg = true
	}
	return jugd, jugd_act

}
