/**********************************************************************************
* Given a linked list, swap every two adjacent nodes and return its head.
*
* For example,
* Given 1->2->3->4, you should return the list as 2->1->4->3.
*
* Your algorithm should use only constant space.
* You may not modify the values in the list, only nodes itself can be changed.
**********************************************************************************/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) printList() {
	list := head
	for list != nil {
		fmt.Printf(" %d", list.Val)
		list = list.Next
	}
	fmt.Println()
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	var count int
	var prev *ListNode
	var first *ListNode
	var last *ListNode

	if head == nil {
		return head
	}

	first = head
	last = head.Next
	count++

	for last != nil {
		if count == 1 {
			tmp := last.Next
			if prev == nil {
				prev = last
				head = last
			} else {
				prev.Next = last
			}
			last.Next = first
			first.Next = tmp
			prev = first
			first = tmp
			last = first
			count = 0
		}

		if first == nil {
			break
		}
		last = first.Next
		count++
	}
	return head
}

func main() {
	//head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}
	head = swapPairs(head)
	head.printList()
}
