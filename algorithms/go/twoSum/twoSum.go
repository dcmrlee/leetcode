/********************************************************************************** 
* 
* Given an array of integers, find two numbers such that they add up to a specific target number.
* 
* The function twoSum should return indices of the two numbers such that they add up to the target, 
* where index1 must be less than index2. Please note that your returned answers (both index1 and index2) 
* are not zero-based.
* 
* You may assume that each input would have exactly one solution.
* 
* Input: numbers={2, 7, 11, 15}, target=9
* Output: index1=1, index2=2
* 
*               
**********************************************************************************/

package main

import (
    "fmt"
)

//
// The implementation as below is bit tricky. but not difficult to understand
//
//  1) Traverse the array one by one
//  2) just put the `target - num[i]`(not `num[i]`) into the map
//     so, when we checking the next num[i], if we found it is exisited in the map.
//     which means we found the second one.
func twoSum(nums []int, target int) []int {
    num2pos := make(map[int]int)
    res := make([]int, 2)
    for i, num := range nums {
        if pos, ok := num2pos[num]; ok {
            res[0] = pos + 1
            res[1] = i + 1
            break
        } else {
            num2pos[target-num] = i
        }
    }
    return res
}

func main() {
    nums := []int{3,2,4}
    target := 6
    result := twoSum(nums, target)
    fmt.Println(result)
}
