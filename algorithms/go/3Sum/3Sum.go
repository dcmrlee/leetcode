/*****************************************************************************************
* Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0?
* Find all unique triplets in the array which gives the sum of zero.
*
* Note: The solution set must not contain duplicate triplets.
* For example, given array S = [-1, 0, 1, 2, -1, -4],
* A solution set is:
* [
* 	[-1, 0, 1],
*	[-1, -1, 2]
* ]
********************************************************************************/

package main

import (
	"fmt"
	"sort"
)

func isSameArray(a []int, sortedB []int) bool {
	len1 := len(a)
	len2 := len(sortedB)

	if len1 != len2 {
		return false
	}
	sort.Ints(a)
	for i := 0; i < len1; i++ {
		if a[i] != sortedB[i] {
			return false
		}
	}

	return true
}

func appendIntArray(result [][]int, item []int) [][]int {
	// make sure result not contain duplicate triplets
	// item reprsent a triplet
	resultLen := len(result)
	sort.Ints(item)
	for i := 0; i < resultLen; i++ {
		if isSameArray(result[i], item) {
			return result
		}
	}
	newSlice := make([][]int, resultLen+1)
	copy(newSlice, result)
	result = newSlice
	result[resultLen] = item
	return result
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		start := i + 1
		end := len(nums) - 1

		for {
			if start >= end {
				break
			}
			b := nums[start]
			c := nums[end]
			if a+b+c == 0 {
				result = appendIntArray(result, []int{a, b, c})
				end = end - 1
			} else if a+b+c > 0 {
				end = end - 1
			} else {
				start = start + 1
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, 4}
	fmt.Println(threeSum(nums))
}
