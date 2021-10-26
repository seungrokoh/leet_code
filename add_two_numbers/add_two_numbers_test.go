package add_two_numbers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	expected := &ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}}
	result := addTwoNumbers(l1, l2)
	assert.Equal(t, expected, result)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwo(l1, l2, 0)
}

func addTwo(l1 *ListNode, l2 *ListNode, alpha int) *ListNode {
	if l1 == nil && l2 == nil {
		if alpha > 0 {
			return &ListNode{Val: alpha}
		}
		return nil
	}
	if l1 != nil && l2 != nil {
		var listNode ListNode
		listNode.Val = (alpha + l1.Val + l2.Val) % 10
		alpha = (alpha + l1.Val + l2.Val) / 10
		listNode.Next = addTwo(l1.Next, l2.Next, alpha)
		return &listNode
	}
	if l1 != nil {
		var listNode ListNode
		listNode.Val = (alpha + l1.Val) % 10
		listNode.Next = addTwo(l1.Next, nil, (alpha+l1.Val)/10)
		return &listNode
	}
	var listNode ListNode
	listNode.Val = (alpha + l2.Val) % 10
	listNode.Next = addTwo(nil, l2.Next, (alpha+l2.Val)/10)
	return &listNode
}
