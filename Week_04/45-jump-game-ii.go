package Week_04

//贪心算法
func jump(nums []int) int {
	maxPosition := 0
	end := 0
	steps := 0
	for i := 0; i < len(nums)-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
