package _52_max_prod_sub_arry

func maxProduct(nums []int) int {

	minFn := func(x []int) int {

		min := x[0]
		for _, n := range x {
			if n < min {
				min = n
			}
		}
		return min
	}
	maxFn := func(x []int) int {

		max := x[0]
		for _, n := range x {
			if n > max {
				max = n
			}
		}
		return max
	}
	if len(nums) == 1 {
		return nums[0]
	}
	result, curMax, curMin := 0, 1, 1
	for _, n := range nums {
		if n == 0 {
			curMax, curMin = 0, 0
			continue
		}
		tmp := curMax * n
		curMax = maxFn([]int{tmp, curMin * n, n})
		curMin = minFn([]int{tmp, curMin * n, n})
		result = maxFn([]int{result, curMax})

	}
	return result
}
