package __longest_substring_withoug_repeating_char

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	result := 0
	left := 0
	for right, ch := range []rune(s) {
		if index, ok := m[ch]; !ok {
			m[ch] = right
		} else {
			if index >= left {
				left = index + 1
			}
			m[ch] = right
		}
		diff := (right - left) + 1
		if diff > result {
			result = diff
		}
	}
	return result
}
