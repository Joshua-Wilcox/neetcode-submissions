func containsNearbyDuplicate(nums []int, k int) bool {

	n := len(nums)
	for i := 0; i < n; i++ {
		for j := 1; j <= k && i + j < n; j++ {
			if nums[i] == nums[i + j] {
				return true
			}
		}
	}
	return false
}
