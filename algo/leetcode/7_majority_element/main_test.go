package main

import (
	"testing"
)

func Test_MajorityElements(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"success case 1": {
			nums: []int{1, 2, 3, 3, 2, 2},
			want: 2,
		},
		"success case 2": {
			nums: []int{1, 3, 1, 2, 2, 3, 3},
			want: 3,
		},
		"success case 3": {
			nums: []int{1, 3, 1},
			want: 1,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := MajorityElement(testCase.nums); got != testCase.want {
				t.Errorf("want : %v not equal to got : %v", testCase.want, got)
			}
		})

	}

}
