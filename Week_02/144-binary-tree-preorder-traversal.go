package Week_02

func preorderTraversal(root *TreeNode) []int {
	if nil == root {
		return nil
	}
	res := []int{root.Val}
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}
