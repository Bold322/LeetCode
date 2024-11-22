package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pizda := &ListNode{0, nil}
	bodolt := pizda
	sanah := 0

	for l1 != nil || l2 != nil {
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		} else {
			x = 0
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		} else {
			y = 0
		}

		sum := x + y + sanah
		sanah = sum / 10
		bodolt.Next = &ListNode{sum % 10, nil}
		bodolt = bodolt.Next

		printList(pizda)
	}

	if sanah > 0 {
		bodolt.Next = &ListNode{sanah, nil}
	}
	printList(pizda)

	return pizda.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{9, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{9, nil}}}}

	result := addTwoNumbers(l1, l2)

	fmt.Println("Result:")

	printList(result)
}
