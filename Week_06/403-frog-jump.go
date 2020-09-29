package Week_06

// 动态规划
// 使用map来进行存储，key为石头位置，value保存可能挑到该石头的步数

func canCross(stones []int) bool {
	stepMap := make(map[int]map[int]int)
	for _, stone := range stones {
		stepMap[stone] = make(map[int]int)
	}
	stepMap[0] = map[int]int{0: 0}
	for _, stone := range stones {
		for step := range stepMap[stone] {
			for i := step - 1; i <= step+1; i++ {
				steps, ok := stepMap[stone+i]
				if ok {
					steps[i] = i
				}
			}
		}
	}
	steps := stepMap[stones[len(stones)-1]]
	return len(steps) > 0
}
