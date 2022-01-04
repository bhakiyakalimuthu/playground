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

func binarySearchRange(intervals []int, newInterval int) int {
	left, right := 0, len(intervals)-1
	mid := 0
	for {
		if left <= right {
			mid = left + (right-left)/2
			if intervals[mid] == newInterval {
				break
			}
			if newInterval < intervals[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			break
		}

	}
	return intervals[mid]
}

func binarySearchSearchNestedSlice(intervals [][]int, newInterval []int) [][]int {
	left, right := 0, len(intervals)-1
	mid := 0
	for {
		if left <= right {
			mid = left + (right-left)/2
			if intervals[mid][0] == newInterval[0] {
				break
			}
			if newInterval[0] < intervals[mid][0] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			break
		}
	}
	insertAt := mid
	if mid == len(intervals)-1 {
		if newInterval[0] > intervals[mid][0] {
			insertAt = mid + 1
		}
	} else if newInterval[0] > intervals[mid][0] {
		insertAt = mid + 1
	}
	intervals = append(intervals[:mid+1], intervals[mid:]...)
	intervals[insertAt] = newInterval
	return intervals
}
