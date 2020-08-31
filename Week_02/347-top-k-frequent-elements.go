package Week_02

import "container/heap"

type Node struct {
	Val   int
	Times int
}

type KFreqHeap []*Node

func (a KFreqHeap) Len() int           { return len(a) }
func (a KFreqHeap) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a KFreqHeap) Less(i, j int) bool { return a[i].Times < a[j].Times }
func (a *KFreqHeap) Push(x interface{}) {
	*a = append(*a, x.(*Node))
}
func (a *KFreqHeap) Pop() interface{} {
	old := *a
	n := len(old)
	x := old[n-1]
	*a = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	if k == 0 || len(nums) == 0 {
		return []int{}
	}

	//count tims
	countMap := make(map[int]int)
	for _, v := range nums {
		countMap[v] = countMap[v] + 1
	}

	kHeap := &KFreqHeap{}
	if k > len(countMap) {
		k = len(countMap)
	}
	size := 0

	for val, times := range countMap {
		node := &Node{
			Val:   val,
			Times: times,
		}

		if size < k {
			heap.Push(kHeap, node)
			size++
		} else {
			if times > (*kHeap)[0].Times {
				heap.Pop(kHeap)
				heap.Push(kHeap, node)
			}
		}

	}
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(kHeap).(*Node).Val)
	}

	return res
}
