func minEatingSpeed(piles []int, h int) int {

	lo, hi := 1,0

	for _, p := range piles {
       if p > hi {
           hi = p
       }
   	}

	var minSpeed int

	for lo <= hi {
		localSpeed := lo + (hi - lo) / 2
		
	
		// find total hours needed using this speed
		localCount := 0
		for _, v := range(piles) {
			localCount += v / localSpeed
			if v % localSpeed != 0 {
				localCount += 1
			}
		}

		if localCount <= h {
			// localSpeed works
			minSpeed = localSpeed
			hi = localSpeed - 1
		} else {
			lo = localSpeed + 1
		}

	}

	return minSpeed

}
