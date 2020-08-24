package Week_01

import "fmt"

func moveZeroes(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if 0 != nums[j] {
			nums[i] = nums[j]
			if j > i {
				nums[j] = 0
			}
			i++
		}
	}
	fmt.Println(nums)
}
