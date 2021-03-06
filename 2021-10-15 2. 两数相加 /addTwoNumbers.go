package main

import "fmt"

func main() {
	l1 := &ListNode{2,&ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5,&ListNode{6, &ListNode{4, nil}}}

	fmt.Print(addTwoNumbers(l1, l2))
}

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0

	for l1 != nil || l2 != nil {

		n1, n2 := 0, 0

		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		carry, sum = int(sum / 10), sum % 10

		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{carry, nil}
	}

	return
}
