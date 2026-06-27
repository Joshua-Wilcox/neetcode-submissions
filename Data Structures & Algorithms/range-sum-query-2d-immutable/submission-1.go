type NumMatrix struct {
    sums [][]int // sums[r][c] = total of everything from (0,0) up to row r-1, col c-1
}

func Constructor(matrix [][]int) NumMatrix {
    rows := len(matrix)
    cols := len(matrix[0])

    // One extra row and column of zeros around the top/left edge.
    // This means we never have to check "am I at the border?" later.
    sums := make([][]int, rows+1)
    for r := range sums {
        sums[r] = make([]int, cols+1)
    }

    for r := 0; r < rows; r++ {
        for c := 0; c < cols; c++ {
            // Running total: this cell, plus the block above it,
            // plus the block to its left, minus the corner we
            // just counted twice.
            sums[r+1][c+1] = matrix[r][c] +
                sums[r][c+1] + // block above
                sums[r+1][c] - // block to the left
                sums[r][c]     // overlap counted twice, remove it
        }
    }

    return NumMatrix{sums: sums}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    return this.sums[row2+1][col2+1] - // whole block to bottom-right corner
        this.sums[row1][col2+1] - // remove the strip above
        this.sums[row2+1][col1] + // remove the strip to the left
        this.sums[row1][col1] // add back the corner removed twice
}