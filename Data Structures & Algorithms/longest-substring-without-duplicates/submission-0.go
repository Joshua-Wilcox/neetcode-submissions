func lengthOfLongestSubstring(s string) int {
    res := 0
    l := 0
    lastSeen := make(map[byte]int)

    for r := 0; r < len(s); r++ {
        if idx, ok := lastSeen[s[r]]; ok && idx >= l {
            l = idx + 1
        }
        lastSeen[s[r]] = r
        if r-l+1 > res {
            res = r - l + 1
        }
    }
    return res
}