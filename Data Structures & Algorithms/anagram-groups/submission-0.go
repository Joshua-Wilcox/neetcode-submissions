func isAnagram(str1 string, str2 string) bool {
	map1 := make(map[byte]int) // char counter
	map2 := make(map[byte]int) // char counter
	for i := 0; i < len(str1); i++ {
		map1[str1[i]] += 1
		map2[str2[i]] += 1
	}

	for key, _ := range map1 {
		if map1[key] != map2[key] {
			return false
		}
	}

	return true

}

func splitAnagrams(strs []string) [][]string {
	// requires all input strings to be of the same length
	// ["act","hat","cat"] => [["hat"],["act", "cat"]]

	m := make(map[int][]string)
	used := make(map[int]bool)
	for i := 0; i < len(strs); i++ {
		if used[i] {
			continue
		}
		m[i] = append(m[i], strs[i])
		used[i] = true
		for j := 0; j < len(strs); j++ {
			if i != j && !used[j] && isAnagram(strs[i], strs[j]) {
				m[i] = append(m[i], strs[j])
				used[j] = true
			}
		}
	}
	output := make([][]string, 0)
	for _, val := range m {
		output = append(output, val)
	}
	return output
}

func groupAnagrams(strs []string) [][]string {
	// given array, group all strings that have the same num of chars

	m := make(map[int][]string,0)
	for i := 0; i < len(strs); i++ {
		thisStr := strs[i]
		m[len(thisStr)] = append(m[len(thisStr)], thisStr)
	}

	output := make([][]string, 0)
	for _, val := range m {
		splits := splitAnagrams(val)
		output = append(output, splits...)
	}
	return output
	
}
