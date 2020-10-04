package Week_06

import "sort"

func leastInterval(tasks []byte, n int) int {
	m := make([]int, 26)
	for _, c := range tasks {
		m[c-'A']++
	}
	sort.Ints(m)
	maxVal := m[25] - 1
	idleSlots := maxVal * n
	for i := 24; i >= 0 && m[i] > 0; i-- {
		idleSlots -= min(m[i], maxVal)
	}
	if idleSlots > 0 {
		return idleSlots + len(tasks)
	}
	return len(tasks)
}
