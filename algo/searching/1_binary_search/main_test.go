package main

import (
	"reflect"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	tests := map[string]struct {
		nums []int
		item int
		want int
	}{
		"success case 1": {
			nums: []int{1, 5, 9, 17, 18, 25, 60, 99},
			item: 18,
			want: 4,
		},
		"success case 2": {
			nums: []int{1, 3, 4, 20, 25, 30},
			item: 20,
			want: 3,
		},
		"success case 3": {
			nums: []int{-11, -5, 0},
			item: 0,
			want: 2,
		},
		"success case 4": {
			nums: []int{-11, -5, 0},
			item: -11,
			want: 0,
		},
		"success not found": {
			nums: []int{-11, -5, 0},
			item: 20,
			want: -1,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := binarySearch(testCase.nums, testCase.item); got != testCase.want {
				t.Errorf("want : %v not equal to got : %v", testCase.want, got)
			}
		})

	}

}

func Test_binarySearchRange(t *testing.T) {
	tests := map[string]struct {
		nums []int
		item int
		want int
	}{
		"success case 1": {
			nums: []int{1, 5, 9, 17, 18, 25, 60, 99},
			item: 20,
			want: 18,
		},
		"success case 2": {
			nums: []int{1, 3, 4, 20, 25, 30},
			item: 21,
			want: 20,
		},
		"success case 3": {
			nums: []int{-11, -5, 0},
			item: 0,
			want: 0,
		},
		"success not found": {
			nums: []int{-11, -5, 0},
			item: 1,
			want: 0,
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := binarySearchRange(testCase.nums, testCase.item); got != testCase.want {
				t.Errorf("want : %v not equal to got : %v", testCase.want, got)
			}
		})

	}
}
func Test_binarySearchSearchNestedSlice(t *testing.T) {
	tests := map[string]struct {
		nums [][]int
		item []int
		want [][]int
	}{
		"success case 1": {
			nums: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			item: []int{4, 8},
			want: [][]int{{1, 2}, {3, 5}, {4, 8}, {6, 7}, {8, 10}, {12, 16}},
		},
		"success case 2": {
			nums: [][]int{{1, 6}, {6, 9}},
			item: []int{4, 8},
			want: [][]int{{1, 6}, {4, 8}, {6, 9}},
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := binarySearchSearchNestedSlice(testCase.nums, testCase.item); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("want : %v not equal to got : %v", testCase.want, got)
			}
		})

	}

}
