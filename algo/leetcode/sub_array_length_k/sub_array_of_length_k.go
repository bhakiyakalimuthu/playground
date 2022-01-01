package sub_array_length_k

//For example, given array = [10, 5, 2, 7, 8, 7] and k = 3, we should get: [10, 7, 8, 8], since:
//
//10 = max(10, 5, 2)
//7 = max(5, 2, 7)
//8 = max(2, 7, 8)
//8 = max(7, 8, 7)

func maxValue(arr []int, k int) []int {
	s := make([]int, 0)
	max, lp := 0, 0

	for rp := 0; rp < len(arr); {
		if (rp - lp) != k {
			if arr[rp] > max {
				max = arr[rp]
			}
			rp++
		} else {
			lp++
			rp = lp
			s = append(s, max)
			max = 0
		}
	}
	s = append(s, max)
	return s
}
