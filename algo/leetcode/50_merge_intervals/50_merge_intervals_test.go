package main

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		arr [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "success case 1",
			args: args{arr: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		//[[1,4],[4,5]]
		{
			name: "success case 2",
			args: args{arr: [][]int{{1, 4}, {4, 5}}},
			want: [][]int{{1, 5}},
		},
		{
			name: "success case 3",
			args: args{arr: [][]int{{13, 18}, {1, 3}, {5, 10}, {2, 6}}},
			want: [][]int{{1, 10}, {13, 18}},
		},
		{
			name: "success case 4",
			args: args{arr: [][]int{{1, 3}, {2, 6}, {5, 10}}},
			want: [][]int{{1, 10}},
		},
		{
			name: "success case 5",
			args: args{arr: [][]int{{1, 5}, {2, 6}, {5, 10}}},
			want: [][]int{{1, 10}},
		},
		{
			name: "success case 5",
			args: args{arr: [][]int{{1, 5}, {2, 4}}},
			want: [][]int{{1, 5}},
		},
		{
			name: "success case 6",
			args: args{arr: [][]int{{1, 5}}},
			want: [][]int{{1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
