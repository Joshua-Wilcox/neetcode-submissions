func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	setMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		setMap[nums[i]] = 1
	}

	var starts []int
	for i := 0; i < len(nums); i++ {
		_, exists := setMap[nums[i] - 1]
		if !exists {
			starts = append(starts, nums[i])
		}
	}
	fmt.Println(starts)
	res := 1
	for i := 0; i < len(starts); i++ {
		thisRes := 1
		thisIterator := starts[i] + 1
		for setMap[thisIterator] == 1 {

			thisRes++
			thisIterator++

		}
		if thisRes > res {
			res = thisRes
		}
	}

	return res
	}
