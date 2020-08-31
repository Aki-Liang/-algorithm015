package Week_02

import "container/list"

func preorder(root *Node) []int {
	if nil == root {
		return nil
	}
	res := []int{}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		elem := stack.Back()
		node := elem.Value.(*Node)
		res = append(res, node.Val)
		stack.Remove(elem)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.PushBack(node.Children[i])
		}
	}
	return res
}
