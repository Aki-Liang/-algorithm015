package Week_06

import "math"

// 动态规划

// 创建二维数组dp
// dp中存储最大连续子数组和的最小值
// dp[i][j]中保存长度为i的数组分为j段之后能得到的最大连续子数组和的最小值
// 设i的前x个数组成了j-1个连续数组，则max(dp[x][j-1], x到i的子数组和）为i长度数组分为j段的最大连续子数组和
// 通过枚举x，取结果的最小值，即可得出i长度数组分为j段后能得到的最大连续子数组和的最小值

func splitArray(nums []int, m int) int {
	n := len(nums)
	dp := make([][]int, n+1)
	sub := make([]int, n+1) //用来快速计算子数组和
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[0][0] = 0
	for i := 0; i < n; i++ {
		sub[i+1] = sub[i] + nums[i]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < min(i, m); j++ {
			for x := 0; x < i; x++ {
				dp[i][j] = min(dp[i][j], max(dp[x][j-1], sub[i]-sub[x]))
			}
		}
	}

	return dp[n][m]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
