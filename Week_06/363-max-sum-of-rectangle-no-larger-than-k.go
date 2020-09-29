package Week_06

import "math"

func maxSumSubmatrix(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	row, col := len(matrix), len(matrix[0])
	for r := 0; r < row; r++ {
		for c := 1; c < col; c++ {
			matrix[r][c] += matrix[r][c-1]
		}
	}
	res := math.MinInt32
	for left := 0; left < col; left++ {
		for right := left; right < col; right++ {
			for top := 0; top < row; top++ {
				sum := 0
				for bottom := top; bottom < row; bottom++ {
					if left == 0 {
						sum += matrix[bottom][right]
					} else {
						sum += matrix[bottom][right] - matrix[bottom][left-1]
					}
					if sum <= k && sum > res {
						res = sum
					}
				}
			}
		}
	}

	return res
}
