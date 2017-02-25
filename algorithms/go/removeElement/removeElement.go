/**********************************************************************************
* Given an array and a value, remove all instances of that value in place
* and return the new length.
* Do not allocate extra space for another array,
* you must do this in place with constant memory.
* The order of elements can be changed.
* It doesn't matter what you leave beyond the new length.
*
* Example:
* Given input array nums = [3,2,2,3], val = 3
* Your function should return length = 2,
* with the first two elements of nums being 2.
**********************************************************************************/

package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[pos] = nums[i]
			pos++
		}
	}
	return pos
}

func main() {
	nums := []int{2}
	fmt.Println(removeElement(nums, 3))
	fmt.Println(nums)
}
