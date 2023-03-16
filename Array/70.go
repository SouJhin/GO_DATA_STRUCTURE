package main

import "fmt"

func main() {
	fmt.Printf("climbStairs(44) =====> 🚀🚀🚀 %v\n", climbStairs(44))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		b := 2
		for a, i := 1, 3; i <= n; i++ {
			t := a
			a = b
			b = t + b
		}
		return b
	}
}
