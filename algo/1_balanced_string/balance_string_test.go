package main

import (
	"testing"
)

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name, in string
		want     bool
	}{
		{
			name: "success case 1",
			in:   "{{(([[]]))}}",
			want: true,
		},
		{
			name: "success case 1",
			in:   "{{([])}}",
			want: true,
		},
		{
			name: "failure case 1",
			in:   "{{{(([[]]))}}",
			want: false,
		},
		{
			name: "failure case 2",
			in:   "{{))(}",
			want: false,
		},
		{
			name: "single char",
			in:   "[",
			want: false,
		},
		{
			name: "single char",
			in:   "[)",
			want: false,
		},
		{
			name: "empty string",
			in:   "",
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := isBalanced(test.in); got != test.want {

				t.Errorf("got : %t not equal to want %t", got, test.want)
			}

		})

	}

}
