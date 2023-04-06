package main

import "fmt"

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	out := mergeIntervals(intervals)
	fmt.Printf("\noutput %+v", out)
}

func mergeIntervals(intervals [][]int) [][]int {
	out := make([][]int, 0, len(intervals))
	last := make([]int, 2)
	for i, current := range intervals {
		if i == 0 {
			last = []int{current[0], current[1]}
			fmt.Printf("\ncurrent %+v", last)
			continue
		}
		if last[1] == current[0] || last[1] > current[0] {
			fmt.Println("inside condition")
			last = []int{last[0], current[1]}
		} else {

			fmt.Println("no condition")
			last = current
		}
		fmt.Printf("\nlast %+v", last)
		out = append(out, last)
	}
	return out
}
