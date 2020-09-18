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

func searchMatrixV2(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	if target > matrix[m-1][n-1] || target < matrix[0][0] {
		return false
	}
	left, right := 0, m*n-1
	for left <= right {
		mid := left + (right-left)>>1
		r := mid / n
		c := mid % n
		if matrix[r][c] == target {
			return true
		} else if matrix[r][c] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return false
}
