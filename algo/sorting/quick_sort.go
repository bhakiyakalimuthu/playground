package sorting

func quickSort(nums []int) []int {

	pIndex := partition(nums, 0, len(nums)-1)
	quickSort(nums, 0, pIndex-1)
	quickSort(nums)
	return nums
}

func partition(nums []int, i, j int) int {
	pivot := nums[0]
	result := 0
	for {
		if i <= j {
			if nums[i] > pivot && nums[j] < pivot {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
			if i == j {
				result = i
			}
			i++
			j++
		} else {
			temp := nums[result]
			nums[result] = nums[0]
			nums[0] = temp
			return result
		}
	}
}
