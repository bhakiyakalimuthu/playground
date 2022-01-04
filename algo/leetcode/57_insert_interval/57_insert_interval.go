package _7_insert_interval

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return append(intervals, newInterval)
	}
	intervals = search(intervals, newInterval)
	out := make([][]int, 0)
	out = append(out, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if out[len(out)-1][1] >= intervals[i][0] {
			item := intervals[i][1]
			if out[len(out)-1][1] > intervals[i][1] {
				item = out[len(out)-1][1]
			}
			out[len(out)-1] = []int{out[len(out)-1][0], item}
		} else {
			out = append(out, intervals[i])
		}
	}
	return out
}

func search(intervals [][]int, newInterval []int) [][]int {

	left, right := 0, len(intervals)-1
	mid := 0
	for {
		if left <= right {
			mid = left + (right-left)/2
			if intervals[mid][0] == newInterval[0] {
				break
			}
			if newInterval[0] < intervals[mid][0] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			break
		}
	}
	insertAt := mid
	if newInterval[0] > intervals[mid][0] {
		insertAt = mid + 1
	}
	intervals = append(intervals[:mid+1], intervals[mid:]...)
	intervals[insertAt] = newInterval
	return intervals
}
