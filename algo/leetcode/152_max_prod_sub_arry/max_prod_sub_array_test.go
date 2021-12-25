package _52_max_prod_sub_arry

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "success case 1",
			nums: []int{-1, -2, -9, -6},
			want: 108,
		},
		{
			name: "success case 2",
			nums: []int{-1},
			want: -1,
		},
		{
			name: "success case 3",
			nums: []int{-2, -5},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
