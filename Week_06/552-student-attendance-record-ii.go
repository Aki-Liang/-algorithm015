package Week_06

//动态规划

//每一层在上一层的末尾+A +P +L
// 有题目可知，符合条件的出勤序列有以下六种情况
// A0L0 没有A，末尾没有L
// A0L1 没有A，末尾1个连续L
// A0L2 没有A，末尾2个连续L
// A1L0 1个A，末尾没有L
// A1L1 1个A，末尾1个连续L
// A1L2 1个A，末尾2个连续L

// 每层在以上六种情况的序列末尾+A +P +L

//状态转移方程
// +P，末尾添加P之后状态全部变成末尾没有连续L
// dp[i][A0L0] = dp[i-1][A0L0] +dp[i-1][A0L1] +dp[i-1][A0L2]
// dp[i][A1L0] = dp[i-1][A1L0] +dp[i-1][A1L1] +dp[i-1][A1L2]

// +L，只有末尾没有两个连续L的情况才能追加L
// dp[i][A0L1] = dp[i-1][A0L0]
// dp[i][A0L2] = dp[i-1][A0L1]
// dp[i][A1L1] = dp[i-1][A1L0]
// dp[i][A1L2] = dp[i-1][A1L1]

// +A, 只有没有A的情况才能追加A,末尾追加A之后全部变成尾部没有L的状态
// 因为前面已经计算过+p的情况，所以本层A1L0的情况= 上层A1L*末尾加P + 上层A0L*末尾加A
// dp[i][A1L0] = dp[i-1][A0L0] +dp[i-1][A0L1] +dp[i-1][A0L2] +dp[i][A1L0]
//             = dp[i][A0L0] + dp[i][A1L0]

func checkRecord(n int) int {
	M := 1000000007
	a0l0, a0l1, a0l2, a1l0, a1l1, a1l2 := 1, 0, 0, 0, 0, 0
	for i := 0; i <= n; i++ {
		//+P
		a0l0Tmp := (a0l0 + a0l1 + a0l2) % M
		//+L
		a0l2 = a0l1
		a0l1 = a0l0
		a0l0 = a0l0Tmp
		//+P +A
		a1l0Tmp := (a0l0 + a1l0 + a1l1 + a1l2) % M
		//+L
		a1l2 = a1l1
		a1l1 = a1l0
		a1l0 = a1l0Tmp
	}

	return a1l0
}
