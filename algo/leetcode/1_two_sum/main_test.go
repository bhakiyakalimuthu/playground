package main

import "testing"

func TestTwoSums(t *testing.T) {

	tests := map[string]struct {
		numbs  []int
		target int
		want1  int
		want2  int
	}{
		"case 1": {
			numbs:  []int{0, 9, 10, 17, 5, 3},
			target: 5,
			want1:  0,
			want2:  4,
		},
		"case 2": {
			numbs:  []int{0, 9, 10, 17, 5, 3},
			target: 27,
			want1:  2,
			want2:  3,
		},
		"case 3": {
			numbs:  []int{0, 9, 10, 17, 5, 3},
			target: 28,
			want1:  0,
			want2:  0,
		},
		"case 4": {
			numbs:  []int{-1, 9, -10, 10, 5, 3},
			target: 0,
			want1:  2,
			want2:  3,
		},
		"case 5": {
			numbs:  []int{-1, 9, -10, 10, 5, 3},
			target: 2,
			want1:  0,
			want2:  5,
		},
	}
	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			got1, got2 := twoSum(testCase.numbs, testCase.target)
			if got1 != testCase.want1 {
				t.Errorf("got1 : %v != want1 : %v", got1, testCase.want1)
				return
			}
			if got2 != testCase.want2 {
				t.Errorf("got2 : %v != want2 : %v", got2, testCase.want2)
				return
			}
		})
	}

}
