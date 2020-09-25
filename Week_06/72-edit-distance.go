package Week_06

// 动态规划
// 将两个字符串继续分解成子字符串的转化
// 构建二维数组dp, #表示空字符串
//     #  r  o  s
// #   0  1  2  3
// h   1  1  2  3
// o   2  2  1  2
// r   3  2  2  2
// s   4  3  3  2
// e   5  4  4  3

// dp每个元素存储两个子串转换的最少步骤数（各子串转换操作最少步骤+1）
// dp[1][1]=1 表示h转换成r需要一步
// dp[2][3] = 3 表示ho转换成ros需要2步【h to r, add s】
// 当两个字符串结尾字符相同时，转换步数等于去掉结尾的两个子串的转换步数
// 如 ho => ro,

// 状态转移方程
// 结尾不同， dp[i][j] = min(dp[i-1][j-1]+1, dp[i][j-1]+1, dp[i-1][j]+1)
// 结尾相同， dp[i][j] = min(dp[i-1][j-1], dp[i][j-1]+1, dp[i-1][j]+1)

// dp[i-1][j-1]: 两个字符串都去掉最后一个字符后转换最少步骤数
// dp[i][j-1]，dp[i-1][j]: 字符串1或字符串2去掉最后一个字符后的转换最少步骤数，
func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	if m == 0 || n == 0 {
		return m + n
	}

	//添加空字符串转换的一列
	m += 1
	n += 1
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = i
	}
	for i := 0; i < n; i++ {
		dp[0][i] = i
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min3(dp[i-1][j-1], dp[i][j-1]+1, dp[i-1][j]+1)
			} else {
				dp[i][j] = min3(dp[i-1][j-1]+1, dp[i][j-1]+1, dp[i-1][j]+1)
			}
		}
	}

	return dp[m-1][n-1]
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
