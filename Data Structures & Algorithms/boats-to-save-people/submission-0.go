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
