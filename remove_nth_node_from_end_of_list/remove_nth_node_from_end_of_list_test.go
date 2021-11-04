package remove_nth_node_from_end_of_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestRemoveNthFromEnd(t *testing.T) {
	ln := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	result := removeNthFromEnd(ln, 2)
	var actual []int
	expected := []int{1, 2, 3, 5}
	for result != nil {
		actual = append(actual, result.Val)
		result = result.Next
	}
	assert.ElementsMatch(t, expected, actual)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var lnArr []*ListNode
	for head != nil {
		lnArr = append(lnArr, &ListNode{
			Val:  head.Val,
			Next: head.Next,
		})
		head = head.Next
	}
	length := len(lnArr)
	var last *ListNode
	for i, ln := range lnArr {
		if length-n != i {
			if head == nil {
				head = &ListNode{
					Val: ln.Val,
				}
				last = head
			} else {
				last.Next = &ListNode{
					Val: ln.Val,
				}
				last = last.Next
			}
		}
	}
	return head
}
