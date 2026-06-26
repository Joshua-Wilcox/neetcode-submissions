func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int) // counting map

	for i := 0; i < len(nums); i++ {
		m[nums[i]] += 1
	}
	// [1,2,2,2,3,3]
	// 1 => 1
	// 2 => 3
	// 3 => 2

	var output []int
	for i := 0; i < k; i++ {
		maxVal := 0
		maxKey := 0
		for key, value := range m {
			if value > maxVal {
				maxVal = value
				maxKey = key
			}
		}
		output = append(output, maxKey)
		delete(m, maxKey)
	}
	return output
}
