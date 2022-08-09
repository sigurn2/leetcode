package main

func minStartValue(nums []int) int {

	sum := 0
	min := 0
	for _, j := range nums {
		sum = sum + j
		if sum < min {
			min = sum
		}
	}
	return min*-1 + 1
}
