package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func binarySearch(nums []int, item int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for {
		if left <= right {
			mid = (left + right) / 2
			if nums[mid] == item {
				return mid
			}
			if item < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			return -1
		}

	}
}
