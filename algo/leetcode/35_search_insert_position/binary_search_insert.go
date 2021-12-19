package leetcode

func binarySearchInsert(nums []int, x int) []int {

	left, right := 0, len(nums)-1
	// binary search to find where to insert
	// find only the index
	for i := 0; i < len(nums); i++ {
		mid := (left + right) / 2
		if x == nums[mid] {
			return insertAt(nums, mid, x)
		}
		if x < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return insertAt(nums, left, x)
}

// insert the item to the index
func insertAt(nums []int, index, item int) []int {
	// if item is greater than length then just append
	if index == len(nums) {
		nums = append(nums, item)
		return nums
	}
	nums = append(nums[:index+1], nums[index:]...)
	if nums[index] < item {
		nums[index+1] = item
	} else {
		nums[index] = item
	}

	return nums
}
