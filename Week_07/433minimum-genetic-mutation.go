package Week_07

// 双向BFS解法
var mutationMap = map[uint8][3]string{
	'A': [...]string{"T", "G", "C"},
	'C': [...]string{"T", "G", "A"},
	'T': [...]string{"A", "G", "C"},
	'G': [...]string{"T", "A", "C"},
}

func idxOf(str string, bank []string) int {
	for i, s := range bank {
		if s == str {
			return i
		}
	}
	return -1
}

func minMutation(start string, end string, bank []string) int {
	if idxOf(end, bank) == -1 {
		return -1
	}

	if start == end {
		return 0
	}

	used := make([]bool, len(bank))
	startQueue := []string{start}
	endQueue := []string{end}
	count := 0

	for len(startQueue) > 0 {
		count++
		l := len(startQueue)
		for _, curr := range startQueue {
			for i := 0; i < len(curr); i++ {
				for _, c := range mutationMap[curr[i]] {
					gene := curr[:i] + c + curr[i+1:]
					if idx := idxOf(gene, endQueue); idx != -1 {
						return count
					}

					if idx := idxOf(gene, bank); idx != -1 && !used[idx] {
						used[idx] = true
						startQueue = append(startQueue, gene)
					}
				}
			}
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
		}
	}

	return -1
}
