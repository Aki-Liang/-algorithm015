package Week_01

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	k := m + n - 1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			k -= 1
			i -= 1
		} else {
			nums1[k] = nums2[j]
			k -= 1
			j -= 1
		}
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		k -= 1
		j -= 1
	}

	fmt.Println(nums1)
}
