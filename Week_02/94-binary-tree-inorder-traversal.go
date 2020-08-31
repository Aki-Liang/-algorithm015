package Week_02

func inorderTraversal(root *TreeNode) []int {
	if nil == root {
		return nil
	}

	res := []int{}
	res = append(res, inorderTraversal(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)
	return res
}
