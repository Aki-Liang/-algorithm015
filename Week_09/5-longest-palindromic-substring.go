package Week_09

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return s
	}

	dp := make([]bool, n)
	maxLen := 0
	left := 0
	right := 0
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			if i == j {
				dp[j] = true
			} else if s[i] == s[j] && (j+1 == i || (j+1 < i && dp[j+1])) {
				dp[j] = true
				l := i - j + 1
				if l > maxLen {
					maxLen, left, right = l, j, i
				}
			} else {
				dp[j] = false
			}
		}
	}
	return s[left : right+1]
}
