package Week_01

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	}
	if nil == l2 {
		return l1
	}

	head := &ListNode{}
	tmp := head
	for nil != l1 && nil != l2 {
		if l1.Val > l2.Val {
			tmp.Next = l2
			l2 = l2.Next
		} else {
			tmp.Next = l1
			l1 = l1.Next
		}
		tmp = tmp.Next
	}
	if nil != l1 {
		tmp.Next = l1
	} else if nil != l2 {
		tmp.Next = l2
	}
	return head.Next
}
