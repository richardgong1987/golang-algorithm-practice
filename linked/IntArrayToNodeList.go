package linked

func IntArrayToNodeList(arr []int) *ListNode {
	root := &ListNode{
		Val: arr[0],
	}
	ret := root
	for index, ele := range arr {
		if index != 0 {
			node := &ListNode{
				Val: ele,
			}
			root.Next = node
			root = root.Next
		}
	}

	return ret
}
