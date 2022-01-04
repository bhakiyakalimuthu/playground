package _7_insert_interval

import (
	"reflect"
	"testing"
)

//Example 1:
//
//Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
//Output: [[1,5],[6,9]]
//Example 2:
//
//Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
//Output: [[1,2],[3,10],[12,16]]
//Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
func Test_insert(t *testing.T) {
	tests := map[string]struct {
		nums [][]int
		item []int
		want [][]int
	}{
		"success case 1": {
			nums: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			item: []int{4, 8},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		"success case 2": {
			nums: [][]int{{1, 3}, {6, 9}},
			item: []int{2, 5},
			want: [][]int{{1, 5}, {6, 9}},
		},
		"success case 3": {
			nums: [][]int{},
			item: []int{2, 5},
			want: [][]int{{2, 5}},
		},
		"success case 4": {
			nums: [][]int{{1, 3}, {6, 9}},
			item: []int{8, 12},
			want: [][]int{{1, 3}, {6, 12}},
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := insert(testCase.nums, testCase.item); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("want : %v not equal to got : %v", testCase.want, got)
			}
		})

	}
}
