package main

/*
https://leetcode.com/problems/add-two-numbers/

You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resultNode := &ListNode{}
	head := resultNode

	carry := false
	p1 := *l1
	p2 := *l2

	for {
		sum := p1.Val + p2.Val
		if carry {
			sum++
		}
		carry = sum >= 10

		head.Next = &ListNode{Val: sum % 10}
		head = head.Next

		if p1.Next == nil && p2.Next == nil {
			break
		}

		if p1.Next != nil {
			p1 = *p1.Next
		} else {
			p1.Val = 0
		}

		if p2.Next != nil {
			p2 = *p2.Next
		} else {
			p2.Val = 0
		}
	}

	if carry {
		head.Next = &ListNode{Val: 1}
	}

	return resultNode.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
