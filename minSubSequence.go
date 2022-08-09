package main

func minSubsequence(nums []int) []int {
	sum := 0
	for _, j := range nums {
		sum += j
	}
	curr := 0
	var ret []int
	for {
		if curr > sum-curr {
			break
		}
		n := findMax(nums)
		curr += n
		ret = append(ret, n)
	}
	return ret
}

func findMax(nums []int) int {
	max := 0
	loc := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			loc = i
		}
	}
	nums[loc] = 0
	return max
}
