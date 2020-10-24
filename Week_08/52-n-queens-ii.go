package Week_08

func totalNQueens(n int) int {
	count := 0
	solveII(n, 0, 0, 0, 0, &count)
	return count
}

func solveII(n, row, columns, diagonals1, diagonals2 int, count *int) {
	if row == n {
		*count++
		return
	}

	bits := (^(columns | diagonals1 | diagonals2)) & ((1 << n) - 1)
	for bits != 0 {
		p := bits & (-bits)      //取到最低位的1
		bits = bits & (bits - 1) //在p的位置上放上皇后
		solveII(n, row+1, columns|p, (diagonals1|p)<<1, (diagonals2|p)>>1, count)
	}

}
