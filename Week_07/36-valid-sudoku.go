package Week_07

func isValidSudoku(board [][]byte) bool {
	rows := make([][]int, 9)
	colums := make([][]int, 9)
	boxes := make([][]int, 9)
	for i := 0; i < 9; i++ {
		rows[i] = make([]int, 9)
		colums[i] = make([]int, 9)
		boxes[i] = make([]int, 9)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c != '.' {
				n := c - '0' - 1 //数字1-9，转换成下标0-8
				boxIndex := (i/3)*3 + j/3

				rows[i][n]++
				colums[j][n]++
				boxes[boxIndex][n]++

				if rows[i][n] > 1 || colums[j][n] > 1 || boxes[boxIndex][n] > 1 {
					return false
				}
			}
		}
	}

	return true
}
