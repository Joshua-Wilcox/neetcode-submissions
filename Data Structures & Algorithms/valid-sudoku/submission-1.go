func isValidSudoku(board [][]byte) bool {


    // rows
    for i := 0; i < 9; i++ {
        m := make(map[byte]int)

        for j := 0; j < 9; j++ {
            thisCell := board[i][j]
            if thisCell == '.' {
                continue
            }
            m[thisCell] += 1

            if m[thisCell] > 1 {
                fmt.Println("duplicate row ", thisCell, "found at ", i,j)
                return false
            }
        }

    }

    // columns
    for i := 0; i < 9; i++ {
        m := make(map[byte]int)
        for j := 0; j < 9; j++ {
            thisCell := board[j][i]
            if thisCell == '.' {
                continue
            }
            m[thisCell] += 1
            if m[thisCell] > 1 {
                fmt.Println("duplicate col ", thisCell, "found at ", i,j)
                return false
            }
        }


    }

    // squares
    var x,y int

    for k := 0; k < 9; k++{
        m := make(map[byte]int)
        for i := 0; i < 3; i++ {
            for j := 0; j <3; j++ {
                
                x = ((k/3 * 3)+ i)
                y = ((k%3 * 3) + j)
                thisCell := board[x][y]
                fmt.Println(x,y,m, thisCell)
                if thisCell == '.' {
                    continue
                }
                
                m[thisCell] += 1

                if m[thisCell] > 1 {
                    fmt.Println("duplicate box ", thisCell, "found at ", i,j)
                    return false
                }
            }
        }
    }

    return true

}

// [
// [".",".",".",".","5",".",".","1","."],
// [".","4",".","3",".",".",".",".","."],
// [".",".",".",".",".","3",".",".","1"],
// ["8",".",".",".",".",".",".","2","."],
// [".",".","2",".","7",".",".",".","."],
// [".","1","5",".",".",".",".",".","."],
// [".",".",".",".",".","2",".",".","."],
// [".","2",".","9",".",".",".",".","."],
// [".",".","4",".",".",".",".",".","."]]
