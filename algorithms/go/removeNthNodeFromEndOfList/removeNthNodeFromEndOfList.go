/**********************************************************************************
* Given a linked list, remove the nth node from the end of list and return its head.
*
* For example:
* 	Given linked list: 1->2->3->4->5, and n = 2.
* 	After removing the second node from the end, the linked list becomes 1->2->3->5.
*
*
* Note:
* 	Given n will always be valid.
*	Try to do this in one pass.
**********************************************************************************/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 {
		return head
	}

	firstPoint := head
	secondPoint := head
	thirdPoint := head

	for i := 0; i < n; i++ {
		firstPoint = firstPoint.Next
	}

	for firstPoint != nil {
		firstPoint = firstPoint.Next
		thirdPoint = secondPoint
		secondPoint = secondPoint.Next
	}

	if secondPoint.Next == nil {
		// meet the end of the list
		if thirdPoint.Next == nil {
			return nil
		}
		thirdPoint.Next = nil
	} else {
		secondPoint.Val = secondPoint.Next.Val
		ptr := secondPoint.Next.Next
		secondPoint.Next = ptr
	}

	// free prevPoint
	return head
}

func main() {
	head := &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}
	ptr := head
	for ptr != nil {
		fmt.Printf("%d ", ptr.Val)
		ptr = ptr.Next
	}

	fmt.Println()
	head = &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}
	ptr = removeNthFromEnd(head, 2)
	for ptr != nil {
		fmt.Printf("%d ", ptr.Val)
		ptr = ptr.Next
	}

	fmt.Println()
	head = &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}
	ptr = removeNthFromEnd(head, 1)
	for ptr != nil {
		fmt.Printf("%d ", ptr.Val)
		ptr = ptr.Next
	}

	fmt.Println()
	head = &ListNode{1, &ListNode{2, nil}}
	ptr = removeNthFromEnd(head, 2)
	for ptr != nil {
		fmt.Printf("%d ", ptr.Val)
		ptr = ptr.Next
	}

	fmt.Println()
	head = &ListNode{1, nil}
	ptr = removeNthFromEnd(head, 1)
	for ptr != nil {
		fmt.Printf("%d ", ptr.Val)
		ptr = ptr.Next
	}
}
