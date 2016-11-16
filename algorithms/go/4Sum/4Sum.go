/**********************************************************************************
* Given an array S of n integers, are there elements a, b, c, and d in S
* such that a + b + c + d = target? Find all unique quadruplets in the array
* which gives the sum of target.
*
* Note: The solution set must not contain duplicate quadruplets.
*
* For example, given array S = [1, 0, -1, 0, -2, 2], and target = 0.
* A solution set is:
* [
* 	[-1,  0, 0, 1],
*	[-2, -1, 1, 2],
* 	[-2,  0, 0, 2]
* ]
* [-2, -1, 0, 0, 1, 2]
*********************************************************************************/

package main

import (
	"fmt"
	"sort"
)

func addNumToInts(arr [][]int, num int) [][]int {
	// For example
	// arr = [[1,2], [3, 4]]
	// num = 5
	// after running should return
	// [[1, 2, 5], [3, 4, 5]]
	if len(arr) == 0 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = append(arr[i], num)
	}

	return arr
}

func mergeTwoInts(a, b [][]int) [][]int {
	len1 := len(a)
	len2 := len(b)
	if len1 == 0 {
		return b
	}
	if len2 == 0 {
		return a
	}
	newSlice := make([][]int, len1+len2)
	copy(newSlice[0:len1], a)
	copy(newSlice[len1:], b)

	return newSlice
}

func appendIntArray(result [][]int, item []int) [][]int {
	// make sure result not contain duplicate triplets
	// item reprsent a triplet
	resultLen := len(result)
	sort.Ints(item)
	newSlice := make([][]int, resultLen+1)
	copy(newSlice, result)
	result = newSlice
	result[resultLen] = item
	return result
}

func twoSum(nums []int, j int, target int) [][]int {
	var result [][]int
	left := j + 1
	right := len(nums) - 1

	for left < right {
		if nums[left]+nums[right] == target {
			result = appendIntArray(result, []int{nums[left], nums[right]})
			// skip the duplication
			for left < len(nums)-1 && nums[left] == nums[left+1] {
				left++
			}
			// skip the duplication
			for right > 0 && nums[right] == nums[right-1] {
				right--
			}
			left++
			right--
		} else if nums[left]+nums[right] > target {
			// skip the duplication
			for right > 0 && nums[right] == nums[right-1] {
				right--
			}
			right--
		} else {
			// skip the duplication
			for left < len(nums)-1 && nums[left] == nums[left+1] {
				left++
			}
			left++
		}
	}

	return result
}

func threeSum(nums []int, i int, target int) [][]int {
	var result [][]int

	for j := i + 1; j < len(nums)-2; j++ {
		if j > i+1 && nums[j] == nums[j-1] {
			// skip the duplication
			continue
		}
		twoSumRes := twoSum(nums, j, target-nums[j])
		tmpRes := addNumToInts(twoSumRes, nums[j])
		result = mergeTwoInts(result, tmpRes)
	}

	return result
}

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)

	var result [][]int
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			// skip the duplication
			continue
		}
		threeSumRes := threeSum(nums, i, target-nums[i])
		tmpRes := addNumToInts(threeSumRes, nums[i])
		result = mergeTwoInts(result, tmpRes)
	}

	return result
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(nums, 0))
}
