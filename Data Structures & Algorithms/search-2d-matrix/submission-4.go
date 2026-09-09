func searchMatrix(matrix [][]int, target int) bool {
	var flatmtx []int

	for _, v := range(matrix) {
		flatmtx = append(flatmtx, v...)
	}

	res := false

	lo, hi := 0, len(flatmtx) -1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if flatmtx[mid] == target {
			res = true
			break
		} else if flatmtx[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return res

}
