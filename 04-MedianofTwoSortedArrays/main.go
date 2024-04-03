package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	mid := length >> 1
	nums := make([]int, 0)
	nums = append(nums, nums1...)
	nums = append(nums, nums2...)

	sort.Ints(nums)

	if length&1 != 0 {
		// 奇数
		return float64(nums[mid])
	}
	return float64(nums[mid]+nums[mid-1]) / 2
}

func main() {
	res := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(res)
}
