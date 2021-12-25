package __longest_substring_withoug_repeating_char

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {

	tests := []struct {
		name string
		arg  string
		want int
	}{
		{
			name: "success case1",
			arg:  "dvdf",
			want: 3,
		},
		{
			name: "success case2",
			arg:  "abcabcbb",
			want: 3,
		},
		{
			name: "success case3",
			arg:  "bbbbb",
			want: 1,
		},
		{
			name: "success case3",
			arg:  "pwwkew",
			want: 3,
		},
		{
			name: "success case3",
			arg:  " ",
			want: 1,
		},
		{
			name: "success case3",
			arg:  "aabaab!bb",
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.arg); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
