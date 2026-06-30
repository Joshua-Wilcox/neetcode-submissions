func reverseString(s []byte) {
	for i := 0; i < (len(s) / 2); i++ {
		j := len(s) - i - 1
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
	}
}
