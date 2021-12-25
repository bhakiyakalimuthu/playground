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

func closureQuickSort() {
	swapFn := func(n []int, x, y int) []int {
		temp := n[x]
		n[x] = n[y]
		n[y] = temp
		return n
	}
	var s func(in []int, start, end int) []int
	s = func(in []int, start, end int) []int {
		if start < end {
			pIndex := func(arr []int, x, y int) int {
				pivot := arr[y]
				pIndex := x
				for i := x; i < len(arr); i++ {
					if arr[i] <= pivot {
						swapFn(arr, i, pIndex)
						pIndex++
					}
				}
				swapFn(arr, y, pIndex)
				return pIndex
			}(in, start, end)
			s(in, start, pIndex-1)
			s(in, pIndex+1, end)
		}
		return in
	}

	sortFn := func(arr []int) []int {

		return s(arr, 0, len(arr)-1)

	}
	sortFn([]int{})
}
