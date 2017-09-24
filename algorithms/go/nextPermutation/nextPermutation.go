/******************************************************************************
Implement next permutation, which rearranges numbers into the lexicographically
next greater permutation of numbers.
If such arrangement is not possible, it must rearrange it as the lowest
possible order (ie, sorted in ascending order).
The replacement must be in-place, do not allocate extra memory.

Here are some examples. Inputs are in the left-hand column and
its corresponding outputs are in the right-hand column.
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1
******************************************************************************/
package main

import (
    "fmt"
)

func main() {
    var nums []int

    
    nums = []int {1, 2, 3, 4}
    nextPermutation(nums)
    fmt.Println(nums)
    fmt.Println(_sliceCmp(nums, []int{1,2,4,3}))

    nums = []int {1, 2, 3}
    nextPermutation(nums)
    fmt.Println(nums)
    fmt.Println(_sliceCmp(nums, []int{1,3,2}))

    nums = []int {3, 2, 1}
    nextPermutation(nums)
    fmt.Println(nums)
    fmt.Println(_sliceCmp(nums, []int{1,2,3}))

    nums = []int {1, 1, 5}
    nextPermutation(nums)
    fmt.Println(nums)
    fmt.Println(_sliceCmp(nums, []int{1,5,1}))

    nums = []int {1, 2, 4, 3}
    nextPermutation(nums)
    fmt.Println(nums)
    fmt.Println(_sliceCmp(nums, []int{1,3,2,4}))

    nums = []int {2,3,1}
    nextPermutation(nums)
    fmt.Println(nums)
    fmt.Println(_sliceCmp(nums, []int{3,1,2}))


    nums = []int {5,4,7,5,3,2}
    nextPermutation(nums)
    fmt.Println(nums)
    fmt.Println(_sliceCmp(nums, []int{5,5,2,3,4,7}))
}

func _sliceCmp(a []int, b []int) bool {
    lenA := len(a)
    lenB := len(b)

    if lenA == lenB {
        for i := 0; i < lenA; i++ {
            if a[i] != b[i] {
                return false
            }
        }
        return true
    }
    return false
}

func _mySmallest(nums []int, num int) int {
    // find all the number larger than num in nums as largerNums
    // return the index of smallest number in largerNums
    // if not found, return -1
    idx := -1
    smallest := nums[0]
    for i := 0; i < len(nums); i++ {
        if nums[i] > num {
            if nums[i] <= smallest {
                idx = i
            }
        }
    }
    return idx
}

func _nextPermutation(nums []int) bool {
    size := len(nums)
    if size == 2 {
        if nums[0] < nums[1] {
            // find result
            nums[0], nums[1] = nums[1], nums[0]
            return true
        } else {
            return false
        }
    }
    // recursive
    ret := _nextPermutation(nums[1:])
    if ret == true {
        return true
    }
    // If initial case is: 1 2 3 4 5
    // Right here, the case must be [2 5 4 3]
    // swap first and last [3 5 4 2]
    // the sort [5 4 2] to [2 4 5]
    // finial result [3 2 4 5]
    //lastIndex := size - 1
    swapIndex := _mySmallest(nums[1:], nums[0])
    if swapIndex != -1 {
        nums[0], nums[swapIndex+1] = nums[swapIndex+1], nums[0]
        changeOrder(nums[1:])
        return true
    }
    return false
}

func changeOrder(nums []int) {
    // from desc to asc or from asc to desc
    low := 0
    high := len(nums) - 1

    for low <= high {
        nums[low], nums[high] = nums[high], nums[low]
        low++
        high--
    }
}

func nextPermutation(nums []int) {
    if len(nums) <= 1 {
        return
    }
    ret := _nextPermutation(nums)
    if ret == true {
        return
    }
    if nums[0] < nums[1] {
        nums[0], nums[1] = nums[1], nums[0]
        changeOrder(nums[1:])
        return
    }
    changeOrder(nums)
}
