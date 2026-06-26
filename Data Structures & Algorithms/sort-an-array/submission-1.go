func heapify(nums []int, heapSize int) []int{
	// binary trees
		// go down left - 2i+1
		// go up (i - 1)/2
		// go down right - 2i+2
	if nums == nil {
		return nil
	}
	for i := heapSize - 1; i > 0; i-- {
		pIdx := (i - 1) / 2
		if nums[pIdx] < nums[i] {
			temp := nums[pIdx]
			nums[pIdx] = nums[i]
			nums[i] = temp
		}
	}

	root := nums[0]
	last := nums[heapSize - 1]
	nums[0] = last
	nums[heapSize-1] = root

	return nums
}

func sortArray(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums = heapify(nums, len(nums) - i)
		
	}
	return nums
}
