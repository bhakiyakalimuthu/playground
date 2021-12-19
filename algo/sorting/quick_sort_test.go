package sorting

import (
	"reflect"
	"testing"
)

func Test_quickSort(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want []int
	}{
		"success case 1": {
			nums: []int{1, 3, 10, 9, 13, 2, 8},
			want: []int{1, 2, 3, 8, 9, 10, 13},
		},
		"success case 2": {
			nums: []int{7, 2, 1, 6, 8, 5, 3, 4},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		"success case 3": {
			nums: []int{1, 0, 3, 5, 2},
			want: []int{0, 1, 2, 3, 5},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := quickSort(testCase.nums); reflect.DeepEqual(got, testCase.want) {
				t.Logf("want : %v equal to got : %v", testCase.want, got)
			} else {
				t.Errorf("want : %v not equal to got : %v", testCase.want, got)
			}
		})

	}

}
