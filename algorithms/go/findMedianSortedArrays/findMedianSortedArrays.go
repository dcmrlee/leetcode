/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

Example 1:
nums1 = [1, 3]
nums2 = [2]
The median is 2.0

Example 2:
nums1 = [1, 2]
nums2 = [3, 4]
The median is (2 + 3)/2 = 2.5

*/

package main

import "fmt"

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Printf("%v\n", findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 == 0 && len2 == 0 {
		return 0.0
	}
	if len1 == 0 {
		m2, _ := findMedianSortedArray(nums2)
		return m2
	}
	if len2 == 0 {
		m1, _ := findMedianSortedArray(nums1)
		return m1
	}
	if len1 == 1 && len2 == 1 {
		return (float64(nums1[0]) + float64(nums2[0])) / 2
	}
	if nums1[0] < nums2[0] {
		if len1 == 1 {
			return findMedianSortedArrays([]int{}, nums2[:len2-1])
		}
		nums1 = nums1[1:]
	} else {
		if len2 == 1 {
			return findMedianSortedArrays(nums1[:len1-1], []int{})
		}
		nums2 = nums2[1:]
	}

	len1 = len(nums1)
	len2 = len(nums2)
	if nums1[len1-1] < nums2[len2-1] {
		if len(nums2) == 1 {
			return findMedianSortedArrays(nums1, []int{})
		}
		nums2 = nums2[:len2-1]
	} else {
		if len(nums1) == 1 {
			return findMedianSortedArrays([]int{}, nums2)
		}
		nums1 = nums1[:len1-1]
	}
	return findMedianSortedArrays(nums1, nums2)
}

func findMedianSortedArray(nums []int) (float64, int) {
	idx := int(len(nums) / 2)
	if len(nums)%2 == 0 {
		return (float64(nums[idx-1]) + float64(nums[idx])) / 2, idx
	}
	return float64(nums[idx]), idx
}
