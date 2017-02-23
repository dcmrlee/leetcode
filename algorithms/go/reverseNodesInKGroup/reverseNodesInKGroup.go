/**********************************************************************************
* Given a linked list, reverse the nodes of a linked list k at a time and
*  return its modified list.
*
* k is a positive integer and is less than or equal to the length of the linked list.
*  If the number of nodes is not a multiple of k then left-out nodes in the end should
*  remain as it is.
*
* You may not alter the values in the nodes, only nodes itself may be changed.
* Only constant memory is allowed.
*
* For example
* Given this linked list: 1->2->3->4->5
* For k = 2, you should return: 2->1->4->3->5
* For k = 3, you should return: 3->2->1->4->5
***********************************************************************************/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (list *ListNode) Print() {
	l := list
	for l != nil {
		fmt.Printf("%d ", l.Val)
		l = l.Next
	}
	fmt.Println()
}

func (head *ListNode) Reverse() *ListNode {
	if head == nil {
		return nil
	}
	var fakeHead, p, q *ListNode
	fakeHead = &ListNode{0, head}
	p = head

	for p != nil {
		q = p.Next
		p.Next = fakeHead
		fakeHead = p
		p = q
	}
	head.Next = nil
	return fakeHead
}

func reverseByK(head *ListNode, k int) *ListNode {
	var fakeHead, start *ListNode
	fakeHead = &ListNode{0, head}

	end := fakeHead
	for end != nil && k > 0 {
		end = end.Next
		k--
	}

	if k < 0 || end == nil {
		// not enough
		return head
	}

	start = head
	for start != end {
		tmp := start.Next
		start.Next = fakeHead
		fakeHead = start
		start = tmp
	}

	head.Next = end.Next
	end.Next = fakeHead

	return end
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}
	var fakeHead, p *ListNode
	var count int

	fakeHead = &ListNode{0, head}
	p = fakeHead

	for p != nil {
		if p == fakeHead {
			fakeHead.Next = reverseByK(fakeHead.Next, k)
		} else {
			p.Next = reverseByK(p.Next, k)
		}
		fakeHead.Print()
		for count = k; count > 0 && p != nil; count-- {
			p = p.Next
		}
		fmt.Println(p)
		fmt.Println(fakeHead)
	}

	return fakeHead.Next
}

func main() {
	//list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	list := &ListNode{1, nil}
	//list = list.Reverse()
	list = reverseKGroup(list, 1)
	//list = reverseByK(list, 3)
	list.Print()
}
