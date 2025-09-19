package other

func SearchInsert(nums []int, target int) int {
	return searchHelp(nums, target, 0)
}

func searchHelp(nums []int, target, offset int) int {
	if len(nums) == 0 {
		return offset
	}
	mid := len(nums) / 2
	//fmt.Println(nums)
	if nums[mid] == target {
		return offset + mid
	} else if nums[mid] > target {
		return searchHelp(nums[:mid], target, offset)
	} else {
		return searchHelp(nums[mid+1:], target, offset+mid+1)
	}
}
