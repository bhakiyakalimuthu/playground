package leetcode

import (
	"reflect"
	"testing"
)

func Test_binarySearchInsert(t *testing.T) {
	tests := map[string]struct {
		nums []int
		item int
		want []int
	}{
		"success case 1": {
			nums: []int{1, 5, 9, 17, 18, 25, 60, 99},
			item: 22,
			want: []int{1, 5, 9, 17, 18, 22, 25, 60, 99},
		},
		"success case 2": {
			nums: []int{1, 3, 4, 20, 25, 30},
			item: 3,
			want: []int{1, 3, 3, 4, 20, 25, 30},
		},
		"success case 2 - if not found less": {
			nums: []int{1, 3, 4, 20, 25, 30},
			item: 35,
			want: []int{1, 3, 4, 20, 25, 30, 35},
		},
		"success not found": {
			nums: []int{-11, -5, 0},
			item: -99,
			want: []int{-99, -11, -5, 0},
		},
		"success not found with one element": {
			nums: []int{2},
			item: 1,
			want: []int{1, 2},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := binarySearchInsert(testCase.nums, testCase.item); reflect.DeepEqual(got, testCase.want) {
				t.Logf("want : %v equal to got : %v", testCase.want, got)
			} else {
				t.Errorf("want : %v not equal to got : %v", testCase.want, got)
			}
		})

	}

}
