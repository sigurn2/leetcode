package leetcode

import "strings"

func generateTheString(n int) string {

	if n%2 == 1 {
		ans := strings.Repeat("a", n)
		return ans
	} else {
		ans := strings.Repeat("a", n-1)

		return ans + "b"
	}
}
