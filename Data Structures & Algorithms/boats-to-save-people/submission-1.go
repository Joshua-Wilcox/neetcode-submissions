func countingSort(inList []int) []int {
    res := []int{}
    m := make(map[int]int)

    for i := 0; i < len(inList); i++ {
        m[i]++
    }

    for key, val := range m {
        for i := 0; i < val; i++ {
            res = append(res, key)
        }
    }

    return res


}
func numRescueBoats(people []int, limit int) int {
    
    sort.Ints(people)
    l, r := 0, len(people) - 1

    count := 0
    // [1l,2,2r,3.,3.]
    for l <= r {
        if people[r] + people[l] > limit {
            r--
        } else {
            l++
            r--
        }
        count++
    }
    return count
}
