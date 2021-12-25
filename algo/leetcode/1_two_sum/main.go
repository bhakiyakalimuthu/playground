package main

func main() {
}

func twoSum(nums []int, target int) (int, int) {
	//for i, _ := range numbs {
	//	if i != len(numbs)-1 {
	//		for j := i + 1; j < len(numbs); j++ {
	//			if numbs[i]+numbs[j] == target {
	//				return i, j
	//			}
	//		}
	//	}
	//}
	//return 0, 0

	m := make(map[int]int, len(nums))

	for index, element := range nums {
		sub := target - element

		if val, ok := m[sub]; ok {
			return val, index
		}
		m[element] = index
	}
	return 0, 0

}
