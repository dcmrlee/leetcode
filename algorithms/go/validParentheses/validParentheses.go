/**********************************************************************************
* Given a string containing just the characters '(', ')', '{', '}', '[' and ']',
* 	determine if the input string is valid.
*
* The brackets must close in the correct order, "()" and "()[]{}" are all valid
*	but "(]" and "([)]" are not.
**********************************************************************************/

package main

import (
	"fmt"
)

func isValid(s string) bool {

	if len(s) == 0 {
		return true
	}

	stackSize := int(len(s) / 2)

	if stackSize == 0 {
		return false
	}

	stack := make([]byte, stackSize+1)
	stackTop := 0

	for i := 0; i < len(s); i++ {
		if stackTop > stackSize {
			return false
		}
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack[stackTop] = s[i]
			stackTop++
		} else {
			if stackTop == 0 {
				return false
			}
			if (s[i] == ')' && stack[stackTop-1] == '(') ||
				(s[i] == ']' && stack[stackTop-1] == '[') ||
				(s[i] == '}' && stack[stackTop-1] == '{') {
				stackTop--
				stack[stackTop] = 0
			} else {
				return false
			}
		}
	}

	if stackTop == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("({[]})"))
	fmt.Println(isValid("(())(]"))
	fmt.Println(isValid("(("))
}
