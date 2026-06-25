func majorityElement(nums []int) int {
    myMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		myMap[nums[i]] += 1
		if myMap[nums[i]] > (len(nums) / 2) {
			return nums[i]
		}
	}

	return -1
}
