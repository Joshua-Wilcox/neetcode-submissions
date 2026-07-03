func maxArea(heights []int) int {
    l, r := 0, len(heights) - 1
    
    max := 0
    for l < r {
        lowerHeight := 0
        if heights[l] < heights[r] {
            lowerHeight = heights[l]
        } else {
            lowerHeight = heights[r]
        }

        thisVol := (r - l) * lowerHeight

        if thisVol > max {
            max = thisVol
        }

        if heights[r] < heights[l] {
            r--
        } else {
            l++
        }

    }

    return max



}
