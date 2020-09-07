package Week_03

var res [][]int

func permute(nums []int) [][]int {
	res = [][]int{}
	pathNums := []int{}
	used := make([]bool, len(nums))
	backTrack(nums, pathNums, used)
	return res
}

func backTrack(nums, pathNums []int, used []bool) {
	if len(nums) == len(pathNums) {
		tmp := make([]int, len(pathNums))
		copy(tmp, pathNums)
		res = append(res, tmp)
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			pathNums = append(pathNums, nums[i])
			backTrack(nums, pathNums, used)
			pathNums = pathNums[:len(pathNums)-1]
			used[i] = false
		}
	}
}
