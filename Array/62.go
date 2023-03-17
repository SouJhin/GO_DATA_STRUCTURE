package main

import "fmt"

func uniquePaths(m int, n int) int {
	nums := make([][]int, m)
	for i := range nums {
		nums[i] = make([]int, n)
		nums[i][0] = 1
	}
	for j := 0; j < n; j++ {
		nums[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			nums[i][j] = nums[i-1][j] + nums[i][j-1]
		}
	}
	return nums[m-1][n-1]
}

func main() {
	fmt.Printf("uniquePaths(3, 9) =====> 🚀🚀🚀 %v\n", uniquePaths(3, 9))
}
