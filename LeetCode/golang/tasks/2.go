package tasks

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	val := 0

	if l1 == nil && l2 == nil {
		return head
	}

	for l1 != nil || l2 != nil || val > 0 {
		sum := val

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		val = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return head.Next
}
