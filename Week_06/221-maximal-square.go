package Week_06

// 动态规划解法
//     原始矩阵              dp                 dp[j][i]中存储以该坐标为右下角的正方形的大小
//    i 0 1 2 3 4        i 0,1,2,3,4           如 dp[1][2]中存储的是 dp[0][1]，dp[0][2]，dp[1][1]，dp[1][2]，四个节点围成的2*2的正方形
//  j                  j
//  0   0,1,1,1,0      0   0,1,1,1,0
//  1   1,1,1,1,0      1   1,1,2,2,0
//  2   0,1,1,1,1      2   0,1,2,3,1
//  3   0,1,1,1,1      3   0,1,2,3,2
//  4   0,0,1,1,1      4   0,0,1,2,3

// 状态转移方程
// dp[m][n] = min(dp[m][n-1], dp[m-1][n-1],dp[m-1][n])+1

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	maxSide := 0
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min3(dp[i][j-1], dp[i-1][j-1], dp[i-1][j]) + 1
				maxSide = max(maxSide, dp[i][j])
			}
		}
	}
	return maxSide * maxSide
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
