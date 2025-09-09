package slices_pointers

func RemoveDuplicates(nums []int) int {
	i, j := 0, 1
	for i < len(nums)-1 {
		if nums[i] == nums[i+1] {
			i++
		} else {
			nums[j] = nums[i+1]
			j++
			i++
		}
	}
	nums = nums[:j]
	return j
}
