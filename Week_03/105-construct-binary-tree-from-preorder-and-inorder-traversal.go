package Week_03

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) || len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}
	//找到中序遍历的跟节点
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])

	return root
}
