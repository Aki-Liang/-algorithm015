package Week_04

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			//左侧递增
			if target <= nums[mid] && target >= nums[left] {
				right = mid - 1 //在左侧
			} else {
				left = mid + 1 // 在右侧
			}
		} else {
			//右侧递增
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1 // 在右侧
			} else {
				right = mid - 1 //在左侧
			}
		}
	}
	return -1
}
