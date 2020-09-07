package Week_03

import "sort"

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	res := [][]int{}
	backTrackUnique(nums, 0, &res)
	return res
}

// 回溯函数实现
// i表示本次函数需要放置的元素位置
func backTrackUnique(nums []int, i int, res *[][]int) {
	l := len(nums)
	if l-1 == i {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	// nums[0:i]是已经决定的部分，nums[i:]是待决定部分，同时待选元素也都在nums[i:]
	for j := i; j < l; j++ {
		//跳过相同的值
		if i != j && nums[i] == nums[j] {
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
		backTrackUnique(nums, i+1, res)
	}

	//还原调整过的数组
	for j := l - 1; j > i; j-- {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
