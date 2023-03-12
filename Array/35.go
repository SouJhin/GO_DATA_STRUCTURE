package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-1)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Printf("searchInsert(nums, 5) =====> 🚀🚀🚀 %v\n", searchInsert(nums, 2))
}
