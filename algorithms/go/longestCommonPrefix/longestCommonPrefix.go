/**********************************************************************************
*
* Write a function to find the longest common prefix string amongst an array of strings.
*
**********************************************************************************/

package main

import (
	"fmt"
)

func commonPrefix(str1, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	if len1 == 0 || len2 == 0 {
		return ""
	}

	i := 0
	j := 0

	for {
		if i >= len(str1) {
			return str1
		}

		if j >= len(str2) {
			return str2
		}

		if str1[i] != str2[j] {
			break
		}
		i++
		j++
	}
	return str1[0:i]
}

func longestCommonPrefix(str []string) string {

	if len(str) == 0 {
		return ""
	}

	result := str[0]

	for i := 1; i < len(str); i++ {
		result = commonPrefix(result, str[i])
	}

	return result
}

func main() {
	str := []string{"abc", "abcd", "ab"}

	fmt.Printf("%s\n", longestCommonPrefix(str))
}
