package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func MajorityElement(nums []int) int {

	m := make(map[int]int, 0)

	max := 0
	key := 0
	for _, val := range nums {

		if _, ok := m[val]; !ok {
			m[val] = 1
		} else {
			m[val] += 1
		}
		if m[val] > max {
			max = m[val]
			key = val
		}
	}
	return key
}
