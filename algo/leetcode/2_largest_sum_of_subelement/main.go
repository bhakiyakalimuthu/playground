package main

// nums - number list
// cons - number of consecutive element
func largestSumOfConsecutiveElement(nums []int, cons int) int {

	//[5, 7, 1, 4, 3]
	//	[7, 1, 4, 3, 6]
	//		[1, 4, 3, 6, 2]
	//			[4, 3, 6, 2, 9]
	//				[3, 6, 2, 9, 2]
	max, sum := 0, 0
	for i, j := 0, cons; j < len(nums); i, j = i+1, j+1 {
		sum += nums[i]
		for x := i + 1; x < j; x++ {
			sum += nums[x]
		}
		if sum > max {
			max = sum
		}
		sum = 0

	}
	return max
}

func maxSumSlidingWindow(nums []int, k int) []int {
	if k < 1 || k > len(nums) {
		return nil
	}
	max, sum := 0, 0
	out := make([]int, 0, k)

	for i, j := 0, k-1; j < len(nums); i, j = i+1, j+1 {
		in := make([]int, 0, k)
		sum = nums[i]
		in = append(in, nums[i])
		for x := i + 1; x <= j; x++ {
			sum += nums[x]
			in = append(in, nums[x])
		}
		if sum > max {
			max = sum
			out = in
		}
		sum = 0
		in = in[:0]

	}
	return out
}

func maxSlidingWindow(nums []int, k int) []int {
	if k < 1 || k > len(nums) {
		return nil
	}

	max := 0
	out := make([]int, 0)
	for i, j := 0, k-1; j < len(nums); i, j = i+1, j+1 {
		for x := i; x <= j; x++ {
			if nums[x] > max {
				max = nums[x]
			}
		}
		out = append(out, max)
	}
	return out
}
