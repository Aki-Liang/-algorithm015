package Week_04

func searchMatrix(matrix [][]int, target int) bool {
	nums := []int{}
	for _, items := range matrix {
		nums = append(nums, items...)
	}
	return searchList(nums, target)
}

func searchList(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	index := len(nums) / 2
	if target == nums[index] {
		return true
	} else if target < nums[index] {
		return searchList(nums[:index], target)
	} else {
		return searchList(nums[index+1:], target)
	}
}
