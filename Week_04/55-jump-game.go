package Week_04

func canJump(nums []int) bool {
	maxPosition := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i >= maxPosition && maxPosition < length-1 {
			return false
		}
	}

	return true
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
