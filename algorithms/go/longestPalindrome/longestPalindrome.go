/*
* Given a string S, find the longest palindromic substring in S.
* You may assume that the maximum length of S is 1000, and there exists one unique longest palindromic substring.
 */

package main

import "fmt"

func main() {
	var s string = "abccbd"
	fmt.Printf("%v\n", longestPalindrome(s))
}

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	var maxStr string
	var tmpStr string
	var strLen int = len(s)
	for idx := 0; idx < (strLen - 1); idx++ {
		tmpStr = findPalindrome(s, idx, idx)
		if len(tmpStr) > len(maxStr) {
			maxStr = tmpStr
		}
		tmpStr = findPalindrome(s, idx, idx+1)
		if len(tmpStr) > len(maxStr) {
			maxStr = tmpStr
		}
	}
	return maxStr
}

func findPalindrome(s string, left int, right int) string {
	strLen := len(s)
	l := left
	r := right
	for l >= 0 && r < strLen && s[l] == s[r] {
		l--
		r++
	}
	return s[(l + 1):r]
}
