func siftDown(nums []int, heapSize int, i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2
		largest := i

		if left < heapSize && nums[left] > nums[largest] {
			largest = left
		}
		if right < heapSize && nums[right] > nums[largest] {
			largest = right
		}
		if largest == i {
			return
		}
		nums[i], nums[largest] = nums[largest], nums[i]
		i = largest
	}
}

func buildHeap(nums []int, heapSize int) {
	for i := heapSize/2 - 1; i >= 0; i-- {
		siftDown(nums, heapSize, i)
	}
}

func sortArray(nums []int) []int {
	n := len(nums)
	buildHeap(nums, n)
	for i := n - 1; i > 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		siftDown(nums, i, 0)
	}
	return nums
}
