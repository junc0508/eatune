package compute

import (
	"github.com/junc0508/eatune/src/price"
)

func Calculate_history_data(history price.Show_candles) (bool, bool) {
	var judg bool
	var judg_act bool
	if history.Candles[1].Volume > 5000 {
		if history.Candles[1].Closebid-history.Candles[1].Openbid > 0 {
			judg_act = true
		} else {
			judg_act = false
		}
		judg = true
	}
	return judg, judg_act

}
