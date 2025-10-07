package slices_pointers

func plusOne(digits []int) []int {
	i := len(digits) - 1
	for i >= 0 {
		if digits[i] != 9 {
			digits[i] += 1
			return append(digits[:i], digits[i:]...)
		} else if digits[i] == 9 && i != 0 {
			digits[i] = 0
			i--
		} else if digits[i] == 9 && i == 0 {
			digits[i] = 0
			return append([]int{1}, digits[i:]...)
		}
	}
	return digits
}
