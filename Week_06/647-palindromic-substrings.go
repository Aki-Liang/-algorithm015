package Week_06

//回文串特点，开头和结束是同一个字符，且去掉开头和结尾的字符后的子串也是回文串

// 动态规划解法
//  这里使用dp表
//   dp[m][n] == true, 表示s[m...n]是回文子串
//   n 0,1,2,3,4,5
// m   a,a,b,a,c,a
// 0 a T,T,F,F,F,F
// 1 a   T,F,T,F,F
// 2 b     T,F,F,F
// 3 a       T,F,T
// 4 c         T,F
// 5 a           T

// 状态转移方程
// 1. m==n
//     dp[m][n]=true
// 2. n-m == 1 && s[m]==s[n],只有两个字符，且两个字符相等，
//     dp[m][n]=true
// 3. n-m >1 && s[m]==s[n],大于两个字符，则需要看当前范围的子串是否为回文串
//     dp[m][n] = dp[m+1][n-1]

func countSubstrings(s string) int {
	count := 0
	length := len(s)

	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}

	//根据状态转移方程3，这里需要先遍历列
	for n := 0; n < length; n++ {
		for m := 0; m <= n; m++ {

			if m == n { //状态转移方程1
				dp[m][n] = true
				count++
			} else if n-m == 1 && s[m] == s[n] { //状态转移方程2
				dp[m][n] = true
				count++
			} else if n-m > 1 && s[m] == s[n] && dp[m+1][n-1] { //状态转移方程3
				dp[m][n] = true
				count++
			}
		}
	}

	return count
}
