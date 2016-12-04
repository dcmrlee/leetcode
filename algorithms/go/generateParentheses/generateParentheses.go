/**********************************************************************************
* Given n pairs of parentheses, write a function to generate all combinations
* of well-formed parentheses.
*
* For example, given n = 3, a solution set is:
*
* [
*	"((()))",
*	"(()())",
*	"(())()",
*	"()(())",
*	"()()()"
* ]
**********************************************************************************/

package main

import (
	"fmt"
	"sort"
)

func generateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}
	if n == 1 {
		return []string{"()"}
	}
	prevResult := generateParenthesis(n - 1)

	// three conditions
	// ([]) []() ()[]
	var result []string
	for _, val := range prevResult {
		midStr := "(" + val + ")"
		leftStr := val + "()"
		rightStr := "()" + val
		// ()()() will be duplicated
		if leftStr != rightStr {
			result = append(result, leftStr)
		}
		result = append(result, rightStr)
		result = append(result, midStr)
	}
	sort.Strings(result)
	return result
}

func main() {
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(4))
}
