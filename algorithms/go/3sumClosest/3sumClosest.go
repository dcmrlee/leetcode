/**********************************************************************************
* Given an array S of n integers, find three integers in S such that the sum is
* closest to a given number, target. Return the sum of the three integers.
* You may assume that each input would have exactly one solution.
*
* For example, given array S = {-1 2 1 -4}, and target = 1.
* The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
**********************************************************************************/

package main

import (
	"fmt"
	"math"
	"sort"
)

func simpleAbs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)

	closestNum := math.MaxInt32
	result := 0
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
			sum := a + b + c
			if sum == target {
				return sum
			} else {
				diff := simpleAbs(sum - target)

				if diff < closestNum {
					closestNum = diff
					result = sum
				}
				if sum-target < 0 {
					start++
				} else {
					end--
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Printf("%d\n", threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Printf("%d\n", threeSumClosest([]int{0, 2, 1, -3}, 1))
	fmt.Printf("%d\n", threeSumClosest([]int{1, 1, -1, -1, 3}, -1))
}
