func mergeAlternately(word1 string, word2 string) string {
	one, two := 0,0

	res := ""
	for one < len(word1) || two < len(word2) {

		if one < len(word1){
			res += string(word1[one])
			one ++
		}

		if two < len(word2){
			res += string(word2[two])
			two++
		}


	}

	return res
}
