package Week_02

import "container/list"

func levelOrder(root *Node) [][]int {
	if nil == root {
		return nil
	}

	res := [][]int{}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		items := []int{}
		for i := 0; i < queue.Len(); i++ {
			elem := queue.Front()
			node := elem.Value.(*Node)
			items = append(items, node.Val)
			queue.Remove(elem)
			for _, n := range node.Children {
				queue.PushBack(n)
			}
		}
		res = append(res, items)
	}
	return res
}
