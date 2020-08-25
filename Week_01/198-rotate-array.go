package Week_01

import "fmt"

func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length
	count := 0
	for i := 0; count < length; i++ {
		prev := nums[i]
		current := i
		for {
			next := (current + k) % length
			prev, nums[next] = nums[next], prev
			count++
			current = next
			if current == i {
				break
			}
		}
	}

	fmt.Println(nums)
}
