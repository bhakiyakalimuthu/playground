package main

import (
	"reflect"
	"testing"
)

func Test_largestSumOfConsecutiveElement(t *testing.T) {

	tests := map[string]struct {
		numbs []int
		size  int
		want  int
	}{
		"case 1": {
			numbs: []int{5, 7, 1, 4, 3, 6, 2, 9, 2},
			size:  5,
			want:  24,
		},
		//"case 2": {
		//	numbs:  []int{0, 9, 10, 17, 5, 3},
		//	target: 27,
		//	want:   2,
		//},
		//"case 3": {
		//	numbs:  []int{0, 9, 10, 17, 5, 3},
		//	target: 28,
		//	want:   0,
		//},
		//"case 4": {
		//	numbs:  []int{-1, 9, -10, 10, 5, 3},
		//	target: 0,
		//	want:   2,
		//},
		//"case 5": {
		//	numbs:  []int{-1, 9, -10, 10, 5, 3},
		//	target: 2,
		//	want:   0,
		//},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := largestSumOfConsecutiveElement(testCase.numbs, testCase.size); got != testCase.want {
				t.Errorf("got1 : %v != want : %v", got, testCase.want)
				return
			}
		})
	}

}

func Test_maxSumSlidingWindow(t *testing.T) {

	tests := map[string]struct {
		numbs []int
		size  int
		want  []int
	}{
		"case 1": {
			numbs: []int{5, 7, 1, 4, 3, 6, 2, 9, 2},
			size:  5,
			want:  []int{4, 3, 6, 2, 9},
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := maxSumSlidingWindow(testCase.numbs, testCase.size); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("got1 : %v != want : %v", got, testCase.want)
				return
			}
		})
	}

}

func Test_maxSlidingWindow(t *testing.T) {

	tests := map[string]struct {
		numbs []int
		size  int
		want  []int
	}{
		"case 1": {
			numbs: []int{1, 3, -1, -3, 5, 3, 6, 7},
			size:  3,
			want:  []int{3, 3, 5, 5, 6, 7},
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := maxSlidingWindow(testCase.numbs, testCase.size); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("got1 : %v != want : %v", got, testCase.want)
				return
			}
		})
	}

}
