// much smarter method after looking at solution
// i really like how it uses such a big key value - tells me how powerful keys can be beyond just as an identifier

func groupAnagrams(strs []string) [][]string {
    m := make(map[[26]int][]string)

    for _, str := range strs {
        var count [26]int
        for _, c := range str {
            count[c-'a']++
        }
        m[count] = append(m[count], str)
    }

    var result [][]string
    for _, group := range m {
        result = append(result, group)
    }
    return result
}