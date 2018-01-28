package compute

import (
	"fmt"
	"github.com/junc0508/eatune/src/price"
	//"sort"
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

func Calculate_RCI(history price.Show_candles) []float64 {
	//tmp := 5 * 4
	//resAr := []int{}
	fset := []float64{}
	for i := 0; i < len(history.Candles); i++ {
		//fset := []float64{0.055, 0.815, 1.0, 0.107}
		fset = append(fset, history.Candles[i].Closeask)
		//fmt.Print(sort.Float64s(history.Candles.Closeask))
		//resAr = append(resAr, history.Candles[i].Closeask[0])
	}
	fmt.Print(fset)
	return fset

}
