package Week_01

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l < 2 {
		return l
	}
	i := 0
	for j := 1; j < l; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}
