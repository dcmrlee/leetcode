/**********************************************************************************
* Given a digit string, return all possible letter combinations that the
*     number could represent.
*
* A mapping of digit to letters (just like on the telephone buttons)
*
* Input:Digit string "23"
* Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
**********************************************************************************/

package main

import (
	"fmt"
)

func combination(digitMap map[string]string, digits string) []string {

	var result []string
	var resultLen int
	if len(digits) == 1 {
		letters := digitMap[digits]
		resultLen = len(letters)
		result = make([]string, resultLen)
		for i := 0; i < resultLen; i++ {
			result[i] = string(letters[i])
		}
	} else {
		numStr := string(digits[0])

		// Recursion
		lastResult := combination(digitMap, digits[1:])
		lastResultLen := len(lastResult)

		letters := digitMap[numStr]
		lettersLen := len(letters)

		resultLen = lettersLen * lastResultLen
		result = make([]string, resultLen)

		resultIdx := 0
		for i := 0; i < lettersLen; i++ {
			for j := 0; j < lastResultLen; j++ {
				result[resultIdx] = string(letters[i]) + lastResult[j]
				resultIdx++
			}
		}
	}

	return result
}

func letterCombinations(digits string) []string {

	digitMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
		"0": " ",
	}

	if len(digits) == 0 {
		return nil
	}

	return combination(digitMap, digits)
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("5678"))
}
