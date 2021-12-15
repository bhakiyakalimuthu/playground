package main

import "testing"

func Test_stockProfit(t *testing.T) {

	tests := []struct {
		name  string
		price []int
		want  int
	}{
		{
			name:  "",
			price: []int{7, 1, 5, 3, 6, 4},
			want:  5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stockProfit(tt.price); got != tt.want {
				t.Errorf("stockProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stockProfitOptimized(t *testing.T) {

	tests := []struct {
		name  string
		price []int
		want  int
	}{
		{
			name:  "success case - if there is a profit",
			price: []int{7, 2, 1, 5, 3, 6, 4},
			want:  5,
		},
		{
			name:  "success case - if there is no profit",
			price: []int{8, 5, 5, 2, 1},
			want:  0,
		},
		{
			name:  "success case - if price list contains negatives",
			price: []int{8, 5, 5, -2, -1},
			want:  0,
		},
		{
			name:  "success case - if price list contains zero",
			price: []int{8, 5, 5, 0, -1, 10},
			want:  5,
		},
		{
			name:  "if price list is empty",
			price: []int{},
			want:  0,
		},
		{
			name:  "if price list is nil",
			price: []int{},
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stockProfitOptimized(tt.price); got != tt.want {
				t.Errorf("stockProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
