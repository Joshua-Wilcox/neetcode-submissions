func search(nums []int, target int) int {
	l, r := 0, len(nums) - 1

	for l < r {
		m := l + (r - l) / 2

		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}

	}

	smallestIndex := l
	
	var binarySearch func(l, r int) int
    binarySearch = func(l, r int) int {
        for l <= r {
            m := (l + r) / 2
            if nums[m] == target {
                return m
            } else if nums[m] < target {
                l = m + 1
            } else {
                r = m - 1
            }
        }
        return -1
    }

    result := binarySearch(0, smallestIndex-1)
    if result != -1 {
        return result
    }

    return binarySearch(smallestIndex, len(nums)-1)
}
