package Week_06

// 原数组                  dp
// 1,3,4,8      1,     3+1=4,  4+4=8, 8+8=16
// 3,2,2,4      3+1=4, 4+2=6,  6+2=8, 8+4=12
// 5,7,1,9      4+5=9, 6+7=13, 8+1=9, 9+9=18
// 2,3,2,3      9+2=11,11+3=14,9+2=11,11+3=14

// 状态转移方程
// m>0 && n==0  dp[m][0] = dp[m-1][0]+grid[m][0]
// m==0 && n>0  dp[0][n] = dp[0][n-1]+grid[0][n]
// m>0 && n>0   dp[m][n] = min(dp[m-1][n], dp[m][n-1]) + grid[m][n]

//动态规划解法
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, columns := len(grid), len(grid[0])
	dp := make([]int, columns)
	for m := 0; m < rows; m++ {
		for n := 0; n < columns; n++ {
			if n-1 >= 0 && m-1 >= 0 {
				dp[n] = min(dp[n-1], dp[n]) + grid[m][n]
			} else if n-1 >= 0 && m == 0 {
				dp[n] = dp[n-1] + grid[m][n]
			} else if n == 0 && m-1 >= 0 {
				dp[n] = dp[n] + grid[m][n]
			} else {
				dp[n] = grid[m][n]
			}
		}
	}

	return dp[columns-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
