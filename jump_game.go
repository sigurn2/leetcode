package main

func jump(nums []int) int {
	m := 0
	step := 0
	end := 0
	for i := 0; i < len(nums)-1; i++ {

		m = max(i+nums[i], m)

		if i == end {
			end = m
			step++
		}
		if m > len(nums)-1 {
			step++
			break
		}
	}
	return step
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
