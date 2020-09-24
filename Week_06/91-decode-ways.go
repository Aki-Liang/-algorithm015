package Week_06

// 动态规划法
// dp中存储译码方法数量，dp[i]存储str[0...i]的方法总数
//s[i]==0, if s[i-1]==1 || s[i-1]==2, 则dp[i]=dp[i-2],否则return 0
// s[i-1]==1, dp[i] = dp[i-1]+dp[i-2]
// s[i-1]==2 && 1<=s[i]<=6, dp[i] = dp[i-1]+dp[i-2]

// index  -1  0 1 2 3 4 5 6
// str        1 2 1 0 1 2 1
// dp      1  1 2 3 2 2 4 6

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	pre, curr := 1, 1 //dp[-1] = dp[0] = 1
	for i := 1; i < len(s); i++ {
		tmp := curr
		if s[i] == '0' {
			if s[i-1] == '1' || s[i-1] == '2' {
				curr = pre
			} else {
				return 0
			}
		} else if s[i-1] == '1' || (s[i-1] == '2' && s[i] >= '1' && s[i] <= '6') {
			curr = curr + pre
		}
		pre = tmp
	}

	return curr
}
