package Week_04

//BFS法
func ladderLength(beginWord string, endWord string, wordList []string) int {
	//建立备选词的映射
	wdict := map[string][]int{}
	for index, w := range wordList {
		for i := range w {
			k := w[0:i] + "*" + w[i+1:]
			if _, ok := wdict[k]; !ok {
				wdict[k] = []int{}
			}
			wdict[k] = append(wdict[k], index)
		}
	}
	// 建立一个数组，标识一下用过的word
	used := make([]bool, len(wordList))
	layer := 1 //初始化BFS的当前层数为1

	wordList = append(wordList, beginWord)
	q := []int{len(wordList) - 1} //这里面存需要搜索的词的index
	for len(q) > 0 {
		nextQ := []int{} //下一层迭代需要查询的词的index集合
		layer++
		for _, index := range q {
			w := wordList[index]
			for i := range w {
				key := w[0:i] + "*" + w[i+1:]
				for _, wIndex := range wdict[key] {
					if wordList[wIndex] == endWord { //接龙完成
						return layer
					}
					if !used[wIndex] { //检查这个词是否被用过
						used[wIndex] = true
						nextQ = append(nextQ, wIndex)
					}
				}
			}
		}
		q = nextQ
	}
	return 0
}
