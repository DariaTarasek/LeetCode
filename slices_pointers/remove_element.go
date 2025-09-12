package slices_pointers

import "fmt"

func RemoveElement(nums []int, val int) int {
	p1, p2 := 0, len(nums)-1
	k := 0
	for p1 <= p2 {
		if nums[p1] == val {
			if nums[p2] == val {
				p2--
			} else {
				nums[p1], nums[p2] = nums[p2], nums[p1]
				p1++
				p2--
				k++
			}
		} else {
			p1++
			k++
		}
	}
	fmt.Println(nums, k)
	return k
}
