package slices_pointers

import "fmt"

func MoveZeroes(nums []int) {
	i, j := 0, 0
	for i < len(nums) && j < len(nums) {
		if nums[i] != 0 {
			nums[j] = nums[i]
			i++
			j++
		} else {
			i++
		}
	}
	for j < len(nums) {
		nums[j] = 0
		j++
	}
	fmt.Println(nums)
}
