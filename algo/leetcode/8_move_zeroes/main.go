package main

func moveZeroes(nums []int) []int {
	anchor := 0
	for explorer := 0; explorer < len(nums); explorer++ {
		if nums[explorer] != 0 {
			if anchor != explorer {
				temp := nums[explorer]
				nums[explorer] = nums[anchor]
				nums[anchor] = temp
			}
			anchor++
		}
	}
	return nums
}
