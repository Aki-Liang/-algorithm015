package Week_04

import "math"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	idxMap := map[string]int{}
	for i, word := range wordList {
		//建立单词和index的映射关系
		idxMap[word] = i
	}

	if _, ok := idxMap[beginWord]; !ok {
		//将beginword添加到列表尾部，方便寻找
		wordList = append(wordList, beginWord)
		idxMap[beginWord] = len(wordList) - 1
	}
	if _, ok := idxMap[endWord]; !ok {
		return [][]string{}
	}
	n := len(wordList)
	edges := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if transformCheck(wordList[i], wordList[j]) {
				edges[i] = append(edges[i], j)
				edges[j] = append(edges[j], i)
			}
		}
	}

	res := [][]string{}
	queue := [][]int{[]int{idxMap[beginWord]}}
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[idxMap[beginWord]] = 0

	for i := 0; i < len(queue); i++ {
		now := queue[i]              //当前层？
		last := now[len(now)-1]      //获取最后一个单词
		if last == idxMap[endWord] { //如果最后一个但是是结束词就填充返回结果集
			tmp := []string{}
			for _, index := range now {
				tmp = append(tmp, wordList[index])
			}
			res = append(res, tmp)
		} else { //不然就还得加
			for _, to := range edges[last] { //遍历最后一个单词能转换的所有单词
				if cost[last]+1 <= cost[to] { //因为初始值是MaxInt32，如果一个单词没有被用过，这里肯定是to的值小于last的值
					cost[to] = cost[last] + 1 //标记为已使用
					tmp := make([]int, len(now))
					copy(tmp, now)
					tmp = append(tmp, to)      //新词追加到末尾
					queue = append(queue, tmp) //放到下一轮迭代
				}
			}
		}
	}
	return res
}

func transformCheck(from, to string) bool {
	for i := 0; i < len(from); i++ {
		if from[i] != to[i] {
			return from[i+1:] == to[i+1:]
		}
	}
	return false
}
