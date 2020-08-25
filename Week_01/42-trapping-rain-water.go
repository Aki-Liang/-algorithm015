package Week_01

func trap(height []int) int {
	length := len(height)
	left := 0
	leftMax := 0
	right := length - 1
	rightMax := 0
	var res int
	for left <= right {
		if leftMax < rightMax {
			res += getMax(0, leftMax-height[left])
			leftMax = getMax(leftMax, height[left])
			left++
		} else {
			res += getMax(0, rightMax-height[right])
			rightMax = getMax(rightMax, height[right])
			right--
		}
	}
	return res
}

func getMax(a, b int) int {
	if a < b {
		return b
	}

	return a
}
