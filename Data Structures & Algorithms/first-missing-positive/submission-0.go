func firstMissingPositive(nums []int) int {

	m := make(map[int]int)

	for _, val := range(nums) {
		m[val]++
	}

	i := 1
	for {
		_, exists := m[i]
		if !exists{
			return i
		}
		i++
	}
}
