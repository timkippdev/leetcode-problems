package main

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	result := addTwoNumbers(l1, l2)
	if result.Val != 7 {
		t.FailNow()
	}
	if result.Next == nil || result.Next.Val != 0 {
		t.FailNow()
	}
	if result.Next.Next == nil || result.Next.Next.Val != 8 {
		t.FailNow()
	}
}
