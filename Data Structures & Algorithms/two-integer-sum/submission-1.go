func twoSum(nums []int, target int) []int {
	
	// nums[j] = target - nums[i]

	myMap := make(map[int]int)
	numsj := 0
	for i := 0; i < len(nums); i++ {
		numsj = target - nums[i]

		j, exists := myMap[numsj]

		if exists && i != j {
			return []int{j,i}
		}

		myMap[nums[i]] = i

	}
	return []int{-1,-1}
}
