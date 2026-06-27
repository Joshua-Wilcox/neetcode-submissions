type Solution struct{}

func (s *Solution) Encode(strs []string) string {
    outputStr := ""

    for i := 0; i < len(strs); i++ {
        outputStr += strconv.Itoa(len(strs[i])) + "#" + strs[i]
    }

    return outputStr

}

func (s *Solution) Decode(encoded string) []string {
    var output []string
    for encoded != "" {
        lenStr := ""
        for encoded[0] != '#' {
            lenStr += string(encoded[0])
            encoded = encoded[1:]
        }
        encoded = encoded[1:]
        
        countOfNextStr, _ := strconv.Atoi(lenStr)
        
        thisStr := encoded[:countOfNextStr]
        output = append(output, thisStr)
        
        encoded = encoded[countOfNextStr:]
    }
    return output
}