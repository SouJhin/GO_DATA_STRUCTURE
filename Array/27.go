package main

import "fmt"

func removeElement(nums []int, val int) int {
	res := 0
	for i, length := 0, len(nums); i < length; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	nums = nums[:res]
	return res
}

func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Printf("removeElement(nums, 3) =====> 🚀🚀🚀 %v\n", removeElement(nums, 3))
}
