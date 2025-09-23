package slices_pointers

/*
[[1,3], [4,6], [8,10]] ||| [2,5]
*/

func InsertInterval(intervals [][]int, newInterval []int) [][]int {
	var res [][]int

	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	index := searchPlaceToInsert(intervals, newInterval[0], 0)
	i := 0
	n := len(intervals)

	for i < index && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = max(intervals[i][1], newInterval[1])
		i++
	}
	res = append(res, newInterval)

	for i < n {
		res = append(res, intervals[i])
		i++
	}

	return res
}

func searchPlaceToInsert(intervals [][]int, newInterval int, s int) int {
	mid := len(intervals) / 2
	if len(intervals) == 0 {
		return s + mid
	}
	if intervals[mid][0] >= newInterval {
		return searchPlaceToInsert(intervals[:mid], newInterval, s)
	} else {
		return searchPlaceToInsert(intervals[mid+1:], newInterval, mid+s+1)
	}
}
