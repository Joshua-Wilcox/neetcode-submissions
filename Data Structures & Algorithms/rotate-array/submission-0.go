func rotateOnce(nums []int) {

	// [1*,2,3,4,5,6]
	// [6,2*,3,4,5,1]
	// [6,1,3*,4,5,2]
	// [6,1,2,4*,5,3]
	// [6,1,2,3,5*,4]
	// [6,1,2,3,4,5]
	n := len(nums)
	for i := 0; i < n; i++ {
		temp := nums[i]
		nums[i] = nums[n-1]
		nums[n-1] = temp
	}
}

func rotate(nums []int, k int) {
	for i:=0; i<k; i++ {
		rotateOnce(nums)
	}
}
