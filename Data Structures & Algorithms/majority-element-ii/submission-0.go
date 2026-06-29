func majorityElement(nums []int) []int {
	m := make(map[int]int)

	for _, val := range(nums) {
		m[val]++
	}
	var res []int
	for key, val := range(m) {
		if val > len(nums) / 3 {
			res = append(res, key)
		}
	}
	return res
}
