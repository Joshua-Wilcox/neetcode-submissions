func hasDuplicate(nums []int) bool {
    var numMap = make(map[int]int)

	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] += 1
		if numMap[nums[i]] > 1 {
			return true
		}
	}
	return false
}
