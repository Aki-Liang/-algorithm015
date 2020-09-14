package Week_04

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if nums[left] <= nums[right] {
			//证明这个区间上是递增的
			return nums[left]
		}
		//这个区间被旋转过，这里获取中点
		mid := left + (right-left)>>1
		if nums[left] <= nums[mid] {
			//左边是连续的，最小数字在右边
			left = mid + 1
		} else {
			//左边不连续，最小数字在左边
			right = mid
		}
	}
	return -1
}
