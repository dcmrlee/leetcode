/*
Given a string, find the length of the longest substring without repeating characters.
Examples:
Given "abcabcbb", the answer is "abc", which the length is 3.
Given "bbbbb", the answer is "b", with the length of 1.
Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/
package main

import "fmt"
import "strings"

func main() {
	s := "dvdf"
	fmt.Printf("max length: %d\n", lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	var maxLen int
	var maxStartPos int
	var maxStr string
	if len(s) == 0 {
		return maxLen
	}
	for i, c := range s {
		pos := strings.IndexByte(maxStr, byte(c))
		if pos == -1 {
			curLen := i - maxStartPos + 1
			if curLen > maxLen {
				maxLen = curLen
			}
		} else {
			maxStartPos += pos + 1
		}
		maxStr = s[maxStartPos : i+1]
	}
	return maxLen
}
