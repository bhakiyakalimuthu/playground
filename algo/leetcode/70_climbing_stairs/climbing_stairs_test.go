package main

import (
	"testing"
)

func Test_findWays(t *testing.T) {
	tests := map[string]struct {
		in   int
		want int
	}{
		"success case 1": {2, 2},
		"success case 2": {3, 3},
		"success case 3": {4, 5},
		"success case 4": {5, 8},
		"success case 5": {35, 8},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			if got := findWays(testCase.in); got != testCase.want {
				t.Errorf("want : %d != got : %d", testCase.want, got)
			} else {
				t.Logf("want : %d == got : %d", testCase.want, got)
			}
		})

	}

}
