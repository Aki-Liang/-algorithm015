package Week_04

// 贪心算法
func maxProfit(prices []int) int {
	res := 0
	for i := 1; i < len(prices); i++ {
		n := prices[i] - prices[i-1]
		if n > 0 {
			res += n
		}
	}
	return res
}
