package _2_AddTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{Val: 123}
	cur := l3
	increment := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + increment
		increment = sum / 10
		sum = sum % 10

		l1, l2 = l1.Next, l2.Next

		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}

	// 两个链表节点相同，但是有进位
	// if l1 == nil && l2 == nil && increment == 1 {
	// 	cur.Next = &ListNode{Val: 1}
	// }

	for l1 != nil {
		sum := l1.Val + increment
		increment = sum / 10
		sum = sum % 10

		l1 = l1.Next

		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}

	for l2 != nil {
		sum := l2.Val + increment
		increment = sum / 10
		sum = sum % 10

		l2 = l2.Next

		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
	}

	// 两个链表节点相同，但是有进位
	if l1 == nil && l2 == nil && increment == 1 {
		cur.Next = &ListNode{Val: 1}
	}

	return l3.Next
}
