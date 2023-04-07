package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {

	tests := map[string]struct {
		in   [][]int
		want [][]int
	}{
		"success case1": {
			in:   [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		"success case2": {
			in:   [][]int{{1, 4}, {4, 5}},
			want: [][]int{{1, 5}},
		},
	}

	for tn, tc := range tests {
		t.Run(tn, func(t *testing.T) {
			got := mergeIntervals(tc.in)

			if !reflect.DeepEqual(got, tc.in) {
				t.Errorf("want:%v got:%v", tc.want, got)
			}

		})

	}
}
