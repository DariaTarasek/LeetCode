package slices_pointers

import (
	"slices"
)

/* intervals = [[1,3],[2,6],[8,10],[15,18]]
{1, 3}, {2, 6}, {4, 10}, {15, 18}

{1, 4}, {2, 3}

[[4,5],[2,4],[4,6],[3,4],[0,0],[1,1],[3,5],[2,2]]
*/

func MergeIntervals(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		c := a[0] - b[0]
		if c == 0 {
			return a[1] - b[1]
		}
		return c
	})

	res := [][]int{{}}
	res[0] = intervals[0]
	intervals = intervals[1:]
	i := 0
	for len(intervals) > 0 {
		if res[i][1] >= intervals[0][0] {
			edge := max(intervals[0][1], res[i][1])
			res[i] = []int{res[i][0], edge}
			intervals = intervals[1:]
		} else {
			res = append(res, intervals[0])
			intervals = intervals[1:]
			i++
		}
	}
	return res
}

func MergeIntervals1(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}
		return a[0] - b[0]
	})

	res := [][]int{intervals[0]}
	for _, interval := range intervals[1:] {
		last := res[len(res)-1]
		if interval[0] <= last[1] {
			// пересекаются → расширяем
			if interval[1] > last[1] {
				res[len(res)-1][1] = interval[1]
			}
		} else {
			// не пересекаются → добавляем новый интервал
			res = append(res, interval)
		}
	}
	return res
}
