package main

import "sort"

func main() {
	merge(nil)
}

////Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
////Output: [[1,6],[8,10],[15,18]]
////Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
////Example 2:
////
////Input: intervals = [[1,4],[4,5]]
////Output: [[1,5]]
//Explanation: Intervals [1,4] and [4,5] are considered overlapping.

type intervals [][]int

func merge(arr [][]int) [][]int {
	//s := make([][]int, 0)
	//s = append(s, []int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18})
	//i := [][]int{{13, 18}, {1, 3}, {5, 10}, {2, 6}}
	//sort.Sort(intervals(i))
	//fmt.Println(i)
	//s := []int{2, 6, 3, 8, 1}
	//sort.Sort(intervals(s))

	sort.Sort(intervals(arr))
	result := make([][]int, 0)
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			result = append(result, arr[i])
			break
		}
		if arr[i][1] >= arr[i+1][0] {
			result = append(result, []int{arr[i][0], arr[i+1][1]})
			i = i + 1
		} else {
			result = append(result, arr[i])
		}
	}
	return result
}

func (i intervals) Len() int {
	return len(i)
}

func (i intervals) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}
func (i intervals) Less(a, b int) bool {
	return i[a][0] < i[b][0]
}
