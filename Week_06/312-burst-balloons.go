package Week_06

//动态规划法

// 在数组两端添加两个值为1的边界元素方便计算

// 创建二维数组作为dp，存储数组上区间的最大金币数量

// 如dp[i][j]中保存数组上下标i到下标j的开区间内戳气球能获得的最大金币数量

// 假设每个区间最后一个被戳爆的球下标为k，可得出状态转移方程
// dp[i][j] =      dp[i][k]       +      dp[k][j]       + nums[i]*nums[k]*nums[j]
// ------------[i,k]区间最大金币数    [k,j]区间最大金币数   因为前面的已经被戳爆了，所以这里nums[i]和nums[j]就成了k的相邻气球

// nums    	1,  3,  1,  5,  8,  1

// dp  		0,  1,  2,  3,  4,  5
// 		 0  0,  0,  3, 30,159,167
// 		 1  0,  0,  0, 15,135,159
// 		 2  0,  0,  0,  0, 40, 40
// 		 3  0,  0,  0,  0,  0, 40
// 		 4  0,  0,  0,  0,  0,  0
// 		 5  0,  0,  0,  0,  0,  0

func maxCoins(nums []int) int {
	n := len(nums)
	rec := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		rec[i] = make([]int, n+2)
	}
	val := make([]int, n+2)
	val[0], val[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		val[i] = nums[i-1]
	}

	//最小的区间至少包含三个元素，从尾部往前扩散，k从小往大递增
	//确保需要使用的子区间已经被计算过
	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := val[i] * val[k] * val[j]
				sum += rec[i][k] + rec[k][j]
				rec[i][j] = max(rec[i][j], sum)
			}
		}
	}
	return rec[0][n+1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
