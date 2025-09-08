package maps_frequencies

import "slices"

func MajorityElement2(nums []int) []int {
	numsCount := make(map[int]int)
	var res []int
	for _, v := range nums {
		numsCount[v] += 1
		if numsCount[v] > len(nums)/3 && !slices.Contains(res, v) {
			res = append(res, v)
		}
	}
	return res
}
