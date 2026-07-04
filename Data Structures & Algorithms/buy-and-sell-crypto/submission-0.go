func maxProfit(prices []int) int {
	l,r := 0,0

	res := 0

	for r < len(prices){

		if prices[r] > prices[l] {
			profit := prices[r] - prices[l]
			if profit > res {
				res = profit
			}
		} else {
			l = r
		}
		r++
	}

	return res



}
