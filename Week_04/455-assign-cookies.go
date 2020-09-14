package Week_04

import "sort"

//贪心算法
func findContentChildren(g []int, s []int) int {
	//尽可能多的满足孩子，所以从胃口最小的孩子开始满足
	sort.Ints(g)
	sort.Ints(s)
	i := 0
	for j := 0; i < len(g) && j < len(s); {
		if g[i] <= s[j] {
			i++
		}
		j++
	}

	return i
}
