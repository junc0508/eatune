package compute

import (
	"github.com/junc0508/eatune/src/price"
	"sort"
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

func Calculate_RCI(history price.Show_candles) []int {
	//tmp := 5 * 4
	resAr := []int{}
	for i := 0; i < len(history.Candles); i++ {
		sort.Sort(history.Candles[i].Closeask))
		resAr = append(resAr, history.Candles[i].Closeask[0])
	}
	return resAr

}
