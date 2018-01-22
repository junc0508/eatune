package compute

import (
	"github.com/junc0508/eatune/src/price"
)

const (
	th_vol  = 50
	ind     = 1
	th_diff = 0
)

func Calculate_history_data(history price.Show_candles) (bool, bool) {
	var judg bool
	var judg_act bool
	if history.Candles[ind].Volume > th_vol {
		if history.Candles[ind].Closebid-history.Candles[ind].Openbid > th_diff {
			judg_act = true
		} else {
			judg_act = false
		}
		judg = true
	}
	return judg, judg_act

}
