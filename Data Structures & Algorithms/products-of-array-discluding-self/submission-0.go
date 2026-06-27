func productExceptSelf(nums []int) []int {
    output := make([]int, len(nums))
    for i := range output {
        output[i] = 1
    }

	// product of everything to the right of index i
    mulRight := 1
    for i := 0; i < len(nums); i++ {
        output[i] = mulRight
        mulRight *= nums[i]
    }

	// product of everything to the left on index i
	// doing it in place on output, making it left prod * right prod, which is prod of all except self
    mulLeft := 1
    for i := len(nums) - 1; i >= 0; i-- {
        output[i] *= mulLeft
        mulLeft *= nums[i]
    }



    return output
}