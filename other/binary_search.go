package other

func Search(nums []int, target int) int {
	return searchHelper(nums, target, 0)
}

func searchHelper(nums []int, target int, offset int) int {
	if len(nums) == 0 {
		return -1
	}
	i := len(nums) / 2
	if target == nums[i] {
		return offset + i
	} else if target > nums[i] {
		return searchHelper(nums[i+1:], target, offset+i+1)
	} else {
		return searchHelper(nums[:i], target, offset)
	}
}
