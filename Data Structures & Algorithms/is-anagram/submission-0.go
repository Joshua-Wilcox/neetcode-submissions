func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	var countMap = make(map[byte][]int)

	for i := 0; i < len(s); i++ {
		countMap[s[i]] = make([]int, 2)
		countMap[t[i]] = make([]int, 2)
	}

	for i := 0; i < len(s); i++ {
		countMap[s[i]][0] += 1
		countMap[t[i]][1] += 1
	}

	for i := 0; i < len(s); i++ {
		if countMap[s[i]] == nil || countMap[t[i]] == nil{
			return false
		}
		if countMap[s[i]][0] != countMap[s[i]][1] {
			return false
		}
	}
	return true

}
