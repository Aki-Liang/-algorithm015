package Week_08

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int][]int)
	for _, value := range arr2 {
		m[value] = []int{}
	}
	other := []int{}
	for _, value := range arr1 {
		if countArr, ok := m[value]; ok {
			countArr = append(countArr, value)
			m[value] = countArr
		} else {
			other = append(other, value)
		}
	}
	sort.Ints(other)
	res := []int{}
	for _, v := range arr2 {
		res = append(res, m[v]...)
	}
	res = append(res, other...)
	return res
}
