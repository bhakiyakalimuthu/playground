package main

import "fmt"

func main() {
	list := []int{10, 12, 11, 5, 6}

	i, j := twoSum(list, 11)
	fmt.Printf("index1: %d, index2: %d", i, j)

	m, n := twoSumEfficient(list, 11)
	fmt.Printf("\nindex1: %d, index2: %d", m, n)
}

func twoSum(list []int, num int) (int, int) {

	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			fmt.Println(i, j)
			if num == (list[i] + list[j]) {
				return i, j
			}
		}
	}
	return 0, 0
}

func twoSumEfficient(list []int, num int) (int, int) {
	m := make(map[int]int, len(list))

	for i, val := range list {
		diff := num - val

		if j, ok := m[diff]; ok {
			if i > j {
				return j, i
			}
			return i, j
		}
		m[val] = i
	}
	return 0, 0

}
