package main

import (
	"sort"
)

//todo
func longestSubarray(nums []int, limit int) int {
	sort.Ints(nums)
	i, j := 0, len(nums)-1
	for {
		if abs(nums[i], nums[j]) > limit {

		}
	}
}
func abs(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}
