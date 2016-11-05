/********************************************************************************** 
* 
* Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). 
* n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). 
* 
* Find two lines, which together with x-axis forms a container, such that the container contains the most water.
* 
* Note: You may not slant the container.
* 
*               
**********************************************************************************/


package main

import (
	"fmt"
)

func min(a int, b int) int {
	var ret int = a
	if (b < a) {
		ret = b
	}
	return ret
}

func maxArea(height []int) int {
	maxarea := 0
	left := 0
	right := len(height) - 1

	for {
		if left >= right {
			break
		}
		area := (right - left) * min(height[left], height[right])
		if area > maxarea {
			maxarea = area
		}
		
		if height[left] < height[right] {
			for {
				left++
				if left >= right || height[left] > height[left-1] {
					break
				}
			}
		} else {
			for {
				right--
				if left >= right || height[right] > height[right+1] {
					break
				}
			}
		}
	}
	
	return maxarea
}

func main() {
	a := []int{4,3,7,5,8}
	b := maxArea(a)
	fmt.Println(b)
}

