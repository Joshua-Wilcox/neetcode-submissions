func isAnagram(s string, t string) bool {

	if len(s) != len(t){
		return false
	}

	mapS, mapT := make(map[rune]int), make(map[rune]int)

	for i := 0; i < len(s); i++ {
		mapS[rune(s[i])] += 1
		mapT[rune(t[i])] += 1 
	}

	for i := 0; i < len(s); i++ {
		if mapS[rune(s[i])] != mapT[rune(s[i])] {
			return false
		}
	}
	return true

}
