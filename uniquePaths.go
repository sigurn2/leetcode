package main

//func uniquePaths(m int, n int) int {
//	if m == 1 {
//		return 1
//	}
//	if n == 1 {
//		return 1
//	}
//	return uniquePaths(m-1, n) + uniquePaths(m, n-1) + 2
//}
func uniquePaths(m int, n int) int {
	mat := make([][]int, m)
	for i, _ := range mat {
		mat[i] = make([]int, n)
		mat[i][0] = 1
	}
	for j := 0; j < n; j++ {
		mat[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			mat[i][j] = mat[i-1][j] + mat[i][j-1]
		}
	}
	return mat[m-1][n-1]
}
