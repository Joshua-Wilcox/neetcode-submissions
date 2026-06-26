func sortColors(nums []int) {
    m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}
	mainIter := 0
	for i := 0; i <= 2; i++{
		for j :=0; j < m[i]; j++ {
			nums[mainIter] = i
			mainIter++
		}
	}
	return
}
