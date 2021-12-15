package main

func main() {
}

func twoSum(numbs []int, target int) (int, int) {
	for i, _ := range numbs {
		if i != len(numbs)-1 {
			for j := i + 1; j < len(numbs); j++ {
				if numbs[i]+numbs[j] == target {
					return i, j
				}
			}
		}
	}
	return 0, 0
}
