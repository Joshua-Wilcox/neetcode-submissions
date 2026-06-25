func longestCommonPrefix(strs []string) string {
    var output string
	letterMap := make(map[int]byte)

	for i:=0; i < len(strs[0]); i++{
		letterMap[i] = strs[0][i]
	}

	maxIndex := 201
	for i:=0; i < len(strs); i++ {
		if strs[i] == "" {
			return ""
		}
		for j:=0; j < len(strs[i]); j++{
			pfxChar, exists := letterMap[j]
			if pfxChar != strs[i][j] || !(exists){
				if j < maxIndex {
					maxIndex = j
				}
			}
		}

		if len(strs[i]) < maxIndex {
			maxIndex = len(strs[i])
		}
	}

	fmt.Println(maxIndex)


	for i := 0; i < maxIndex; i++ {
		output += string(strs[0][i])
	}
	return output
}
