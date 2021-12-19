package leetcode

func binarySearchInsertPosition(nums []int, x int) int {
	left, right := 0, len(nums)-1

	for i := 0; i < len(nums); i++ {
		mid := (left + right) / 2
		if x == nums[mid] {
			return mid
		}
		if x < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
