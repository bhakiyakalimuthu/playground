package main

import (
	"testing"
)

func Test_maxSumOfSubArray(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"success case 1": {
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6,
		},
		"success case 2": {
			nums: []int{-2, -1, -3, -4, -1, 1, -5, -4},
			want: 1,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := maxSumOfSubArray(testCase.nums); got != testCase.want {
				t.Errorf("want : %v not equal to got : %v", testCase.want, got)
			}
		})

	}

}
