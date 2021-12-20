package __longest_substring_withoug_repeating_char

func lengthOfLongestSubstring(s string) int {

	charMap := make(map[rune]struct{})
	result := 0
	temp := make([]rune, 0)

	for _, val := range s {
		if _, ok := charMap[val]; !ok {
			charMap[val] = struct{}{}
			temp = append(temp, val)
		} else {
			for index, item := range temp {
				delete(charMap, item)
				temp = temp[index+1:]
				if item == val {
					break
				}
			}
			temp = append(temp, val)
			charMap[val] = struct{}{}
			if len(temp) > result {
				result = len(temp)
			}
		}
	}
	if len(temp) > result {
		result = len(temp)
	}
	return result
}
