package sorting

func quickSort(nums []int) []int {
	return sort(nums, 0, len(nums)-1)
}

func sort(nums []int, start, end int) []int {
	if start < end {
		pIndex := partition(nums, start, end)
		sort(nums, start, pIndex-1)
		sort(nums, pIndex+1, end)

	}
	return nums
}
func partition(nums []int, start, end int) int {
	pivot := nums[end]
	pIndex := start

	for i := start; i < end; i++ {
		if nums[i] <= pivot {
			swap(nums, i, pIndex)
			pIndex++
		}
	}
	swap(nums, end, pIndex)
	return pIndex
}

func swap(nums []int, i, j int) []int {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
	return nums
}