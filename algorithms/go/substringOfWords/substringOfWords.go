/*
 * You are given a string, s, and a list of words, words, that are all of the same length.
 * Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once
 *      and without any intervening characters.
 *
 * For example, given:
 * s: "barfoothefoobarman"
 * words: ["foo", "bar"]
 * You should return the indices: [0,9]. (order does not matter)
 */
package main

import (
	"fmt"
)

func makeStrToFixedLengthWords(s string, length int) []string {
	wordsSize := int(len(s) / length)
	words := make([]string, wordsSize)

	var wordsIdx int
	for wordsIdx = 0; wordsIdx < wordsSize; wordsIdx++ {
		abc := wordsIdx
		//fmt.Println(abc*length)
		word := make([]byte, length)
		j := 0
		for i := abc * length; i < abc*length+length; i++ {
			word[j] = s[i]
			j++
		}
		//fmt.Println(string(word))
		words[wordsIdx] = string(word)
	}

	//fmt.Println(len(words))
	return words
}

func isSameWords(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if len(a) == 0 && len(b) == 0 {
		return true
	}
	str := a[0]
	for i := 0; i < len(b); i++ {
		if str == b[i] {
			newb := append(b[:i], b[i+1:]...)
			return isSameWords(a[1:], newb)
		}
	}
	return false
}

func doFindSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	wordsLen := len(words)
	inputWords := makeStrToFixedLengthWords(s, wordLen)
	isSameWords(inputWords[0:len(words)], words)

	var result []int
	for i := 0; i <= len(inputWords)-wordsLen; {
		if isSameWords(inputWords[i:i+wordsLen], words) {
			result = append(result, i*wordLen)
			i = i + 2
		} else {
			i++
		}
	}

	return result
}

func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	var result []int

	for i := 0; i < wordLen; i++ {
		ret := doFindSubstring(s[i:], words)
		if len(ret) != 0 {
			for j := 0; j < len(ret); j++ {
				result = append(result, ret[j]+i)
			}
		}
	}
	return result
}

func main() {
	s := "barfoothefoobarman"
	words := []string{"foo", "bar"}

	fmt.Println(findSubstring(s, words))
}
