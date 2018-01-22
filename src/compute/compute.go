package compute

func Calculate_history_data(End_point string, Token string) {
	var jugd bool
	var jugd_act bool
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
