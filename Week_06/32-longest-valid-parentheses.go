package Week_06

// 动态规划法
// index:0   1   2   3   4   5   6   7
//       (   )   )   (   (   (   )   )
// dp:   0   2   0   0   0   0   2   4
// 			         \           \2+/2=4
// 			          \__________0+/
func longestValidParentheses(s string) int {
	ans := 0
	dp := make([]int, len(s)) //用来保存每个节点对应的有效长度
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				//与前一个正好组成一对
				if i >= 2 {
					//前面连着的一组有效长度+2
					dp[i] = dp[i-2] + 2
				} else {
					//前面没括号了，那当前节点和前面的括号组成一个单独的有效子串，长度为2
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				// 检查当前节点的）能否和之前的子串组成一个新的大子串
				// i-dp[i-1] > 0检查略过前一子串之后是否存在一个节点可以和当前节点组成有效子串
				// 如果节点存在，则通过s[i-dp[i-1]-1] == '('来判断该节点是否和当节点匹配
				if i-dp[i-1] >= 2 {
					//如果前面还有更多的节点
					//当前节点的有效长度=前一个有效子串的长度+再往前一个相邻节点的有效长度+两个括号
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					//前面没有更多节点了
					//当前节点的有效长度=前一个有效子串的长度+两个括号
					dp[i] = dp[i-1] + 2
				}

			}
		}
		//保存最大的子串长度
		ans = max(ans, dp[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
