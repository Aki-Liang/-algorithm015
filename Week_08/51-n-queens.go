package Week_08

import "math/bits"

func solveNQueens(n int) [][]string {
	solutions := [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	solve(queens, n, 0, 0, 0, 0, &solutions)
	return solutions
}

func solve(queens []int, n, row, columns, diagonals1, diagonals2 int, solutions *[][]string) {
	if row == n {
		board := generateBoard(queens, n)
		*solutions = append(*solutions, board)
		return
	}
	availablePositions := ((1 << n) - 1) & (^(columns | diagonals1 | diagonals2))
	for availablePositions != 0 {
		position := availablePositions & (-availablePositions)
		availablePositions = availablePositions & (availablePositions - 1)
		column := bits.OnesCount(uint(position - 1))
		queens[row] = column
		solve(queens, n, row+1, columns|position, (diagonals1|position)>>1, (diagonals2|position)<<1, solutions)
		queens[row] = -1
	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
