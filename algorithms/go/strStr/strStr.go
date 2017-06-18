package main

import (
	"fmt"
)

func isFind(haystack string, needle string) bool {
	if len(haystack) >= len(needle) {
		if haystack[0] == needle[0] {
			if len(needle) == 1 {
				return true
			} else {
				return isFind(haystack[1:], needle[1:])
			}
		}
	}
	return false
}

func strStr(haystack string, needle string) int {
	pos := -1

	if len(needle) == 0 {
		return 0
	}

	if len(haystack) >= len(needle) {
		for idx := 0; idx < len(haystack); idx++ {
			if needle[0] == haystack[idx] {
				if isFind(haystack[idx:], needle) {
					pos = idx
					break
				}
			}
		}
	}
	return pos
}


func main() {
	a := ""
	b := "a"
	fmt.Println(strStr(a, b))
}