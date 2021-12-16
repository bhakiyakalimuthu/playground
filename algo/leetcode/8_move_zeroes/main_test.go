package main

import (
	"reflect"
	"testing"
)

func Test_MajorityElements(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want []int
	}{
		"success case 1": {
			nums: []int{0, 0, 6, 0, 6, 8, 6},
			want: []int{6, 6, 8, 6, 0, 0, 0},
		},
		"success case 2": {
			nums: []int{1, 3, 0, 2, 0, 3, 3},
			want: []int{1, 3, 2, 3, 3, 0, 0},
		},
		"success case 3": {
			nums: []int{1, 0, 1, 2, 0, 3, 3},
			want: []int{1, 1, 2, 3, 3, 0, 0},
		},
		"negative case 1 - no zeros": {
			nums: []int{1, 3, 6, 6},
			want: []int{1, 3, 6, 6},
		},
		"negative case 1 - only zeros": {
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := moveZeroes(testCase.nums); reflect.DeepEqual(got, testCase.want) {
				t.Logf("want : %v equal to got : %v", testCase.want, got)
			} else {
				t.Errorf("want : %v not equal to got : %v", testCase.want, got)
			}
		})

	}

}
