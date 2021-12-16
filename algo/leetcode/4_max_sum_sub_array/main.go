package main

//Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
//Output: 6
//Explanation: [4,-1,2,1] has the largest sum = 6.

func maxSumOfSubArray(nums []int) int {
	max := 0
	sum := 0

	for i := 0; i < len(nums); i++ {
		if sum+nums[i] > 0 {
			sum += nums[i]
		} else {
			sum = 0
		}

		if sum > max {
			max = sum
		}
	}

	return max

}
