package main

import "fmt"

// Insert merges a new [start, end] interval into a sorted, non-overlapping
// list, keeping it sorted and non-overlapping. Three phases: copy the
// intervals ending before it, absorb everything overlapping, copy the rest.
func Insert(intervals [][2]int, newInterval [2]int) [][2]int {
	out := [][2]int{}
	i, n := 0, len(intervals)
	for i < n && intervals[i][1] < newInterval[0] {
		out = append(out, intervals[i]) // entirely before: untouched
		i++
	}
	for i < n && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++ // overlapping: swallowed into the growing new interval
	}
	out = append(out, newInterval)
	for i < n {
		out = append(out, intervals[i]) // entirely after: untouched
		i++
	}
	return out
}

func main() {
	fmt.Println(Insert([][2]int{{1, 3}, {6, 9}}, [2]int{2, 5}))
	fmt.Println(Insert([][2]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, [2]int{4, 8}))
}
