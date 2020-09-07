package Week_03

func combine(n int, k int) [][]int {
	res := [][]int{}
	if n <= 0 || k <= 0 || n < k {
		return res
	}
	nums := []int{}
	trackBackCombine(n, 1, k, nums, &res)
	return res
}

func trackBackCombine(n, m, k int, nums []int, res *[][]int) {
	if len(nums) == k {
		tmp := make([]int, k)
		copy(tmp, nums)
		*res = append(*res, tmp)
		return
	}

	for ; m <= n; m++ {
		nums = append(nums, m)
		trackBackCombine(n, m+1, k, nums, res)
		nums = nums[:len(nums)-1]
	}
}
