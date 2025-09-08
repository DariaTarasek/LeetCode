package strings_slices

func TwoSum(nums []int, target int) []int {
	var ans []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ans = append(ans, i, j)
				return ans
			}
		}
	}
	return ans
}
