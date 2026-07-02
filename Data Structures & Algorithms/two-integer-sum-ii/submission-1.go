func twoSum(numbers []int, target int) []int {
	l := 0
	r := len(numbers) - 1

	for l < r {

		// [1,5,9,13,15,20,29], target = 33
		if numbers[l] + numbers[r] < target {
			l++
		} else if numbers[l] + numbers[r] > target {
			r--
		} else {
				res := []int{l+1, r+1}
		return res
		}

	}

	return []int{}



}
