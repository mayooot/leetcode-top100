package _2_AddTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	increment := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + increment
		increment = sum / 10
		sum = sum % 10
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		sum := l1.Val + increment
		increment = sum / 10
		sum = sum % 10
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		sum := l2.Val + increment
		increment = sum / 10
		sum = sum % 10
		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
		l2 = l2.Next
	}

	if l1 == nil && l2 == nil && increment == 1 {
		cur.Next = &ListNode{Val: 1}
	}
	return pre.Next
}
