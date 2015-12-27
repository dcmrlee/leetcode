/****************************************************************************************
* You are given two linked lists representing two non-negative numbers. 
* The digits are stored in reverse order and each of their nodes contain a single digit. 
* Add the two numbers and return it as a linked list.

* Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
* Output: 7 -> 0 -> 8
* 
* 342 + 465 = 807
*****************************************************************************************/

package main

import (
    "fmt"
)

type ListNode struct {
    val int
    next *ListNode
}

type action func(*ListNode)

func walkList(l *ListNode, fn action) {
    for ; l != nil; l = l.next {
        fn(l)
    }
}

func printList(l *ListNode) {
    fmt.Print(l.val)
}

/*
*  Return: carry number, a single digit
*/
func addTwo(num1 int, num2 int) (int, int) {
    return (num1+num2) / int(10), (num1+num2) % int(10)
}

func (l *ListNode) getNodeValue() int {
    ret := 0
    if l != nil {
        ret = l.val
    }
    return ret
}

func (l *ListNode) moveNext() (*ListNode) {
    if l != nil {
       return l.next
    } else {
        return nil
    }
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (*ListNode) {
    var result *ListNode
    t := &result
    carry := 0
    for l1 != nil || l2 != nil {
        a := l1.getNodeValue()
        b := l2.getNodeValue()

        newNode := new(ListNode)
        carry, newNode.val = addTwo(a, b+carry)
        *t = newNode
        t = &newNode.next

        l1 = l1.moveNext()
        l2 = l2.moveNext()
    }
    if carry > 0 {
        newNode := new(ListNode)
        newNode.val = carry
        *t = newNode
    }
    return result
}

func main() {
    l1 := &ListNode{2,&ListNode{4,&ListNode{5,nil}}} 
    l2 := &ListNode{5,&ListNode{6,&ListNode{4,&ListNode{9,nil}}}}
    walkList(l1, printList)
    fmt.Println()
    walkList(l2, printList)
    fmt.Println()
    res := addTwoNumbers(l1, l2)
    walkList(res, printList)
    fmt.Println()
}
