package compute

import (
	"fmt"
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

func Calculate_RCI(history price.Show_candles) float64 {
	//tmp := 5 * 4
	//resAr := []int{}
	var tmp_p, tmp_c, tmp float64
	fset := []float64{}
	fset_c := []float64{}
	for i := 0; i < len(history.Candles)-1; i++ {
		fset = append(fset, history.Candles[i].Closeask)
	}
	//fmt.Print(fset)
	sort.Float64s(fset)
	for i := 0; i < len(history.Candles)-1; i++ {
		for j := 0; j < len(history.Candles)-1; j++ {
			fset = append(fset, history.Candles[i].Closeask)
			if history.Candles[i].Closeask == fset[j] {
				tmp_p = tmp_p + float64((i-j)*(i-j))
			}
		}
	}
	tmp_p = (1 - (6*tmp_p)/(5*24)) * 100

	for i := 1; i < len(history.Candles); i++ {
		fset_c = append(fset, history.Candles[i].Closeask)
	}
	//fmt.Print(fset)
	sort.Float64s(fset_c)
	for i := 1; i < len(history.Candles); i++ {
		for j := 1; j < len(history.Candles); j++ {
			fset = append(fset_c, history.Candles[i].Closeask)
			if history.Candles[i].Closeask == fset_c[j] {
				tmp_c = tmp_c + float64((i-j)*(i-j))
			}
		}
	}
	//fmt.Print("\n", tmp)
	tmp_c = (1 - (6*tmp_c)/(5*24)) * 100

	tmp = tmp_c - tmp_p
	fmt.Println(tmp_p, tmp_c, tmp)
	return tmp

}
