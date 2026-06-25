func removeElement(nums []int, val int) int {
    n := len(nums)
    currentSwapIdx := n-1
    returnVal := n

    for i := 0; i < n; i++ {
        fmt.Println(nums)

        if nums[i] == val && i < currentSwapIdx{ 
            currentTailVal := nums[currentSwapIdx]
            for currentTailVal == val && currentSwapIdx > i {
                currentSwapIdx -= 1
                currentTailVal = nums[currentSwapIdx]
            }

            if currentSwapIdx <= i {
                 break
            }

            nums[currentSwapIdx] = val
            nums[i] = currentTailVal
            fmt.Println("swapped these two indexes :", i, currentSwapIdx)
            currentSwapIdx -= 1
        }
        
    }

    for i := 0; i < n; i++ {
        if nums[i] == val {
            returnVal -= 1
        }
    }

    return returnVal
}
