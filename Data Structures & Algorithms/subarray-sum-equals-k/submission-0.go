func subarraySum(nums []int, k int) int {
    res, curSum := 0, 0
    m := map[int]int{0: 1}

    for _, num := range nums {
        curSum += num
        diff := curSum - k
        res += m[diff]
        m[curSum]++
    }

    return res
}