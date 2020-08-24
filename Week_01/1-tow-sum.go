package Week_01

func twoSum(nums []int, target int) []int {

	resMap := make(map[int]int)
	for index, num := range nums {
		idx, ok := resMap[target-num]
		if ok {
			return []int{idx, index}
		}
		resMap[num] = index
	}

	return nil
}
