/**********************************************************************************
*
* Merge k sorted linked lists and return it as one sorted list.
* Analyze and describe its complexity.
*
**********************************************************************************/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return lists[0]
}

func main() {
	lists := []*ListNode{
		&ListNode{5, &ListNode{7, nil}},
		&ListNode{3, &ListNode{4, &ListNode{2, nil}}},
		&ListNode{10, &ListNode{9, &ListNode{1, nil}}},
	}
	fmt.Println(mergeKLists(lists))
}
