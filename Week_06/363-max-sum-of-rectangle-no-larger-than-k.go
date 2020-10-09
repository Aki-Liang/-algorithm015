package Week_06

import "math"

//更新解法
func maxSumSubmatrixV2(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows := len(matrix)
	cols := len(matrix[0])
	max := math.MinInt32
	for l := 0; l < cols; l++ { //枚举左边界
		rowSum := make([]int, rows)
		for r := l; r < cols; r++ { // 枚举右边界
			for i := 0; i < rows; i++ { // 按每一行累计到 dp
				rowSum[i] += matrix[i][r]
			}
			max = MaxInt(max, dpmax(rowSum, k))
			if max == k {
				return k
			}
		}
	}
	return max
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dpmax(arr []int, k int) int {
	rollMax, rollSum := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		if rollSum > 0 {
			rollSum += arr[i]
		} else {
			rollSum = arr[i]
		}
		if rollSum > rollMax {
			rollMax = rollSum
		}
	}
	if rollMax <= k {
		return rollMax
	}

	max := math.MinInt32
	for i := 0; i < len(arr); i++ {
		sum := 0
		for r := i; r < len(arr); r++ {
			sum += arr[r]
			if sum > max && sum <= k {
				max = sum
			}
			if max == k {
				return k
			}
		}
	}

	return max
}

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
