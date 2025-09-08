package maps_frequencies

func MajorityElement(nums []int) int {
	candidate, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			candidate = v
			count += 1
		} else {
			if v == candidate {
				count += 1
			} else {
				count -= 1
			}
		}
	}
	return candidate
}
