package strings_slices

func SlidingWindow(s string) int {
	left, right := 0, 0
	substringsLetters := make(map[rune]bool)
	var maxLength int
	for left < len(s) && right < len(s) {
		if _, ok := substringsLetters[rune(s[right])]; ok == false {
			substringsLetters[rune(s[right])] = true
			if right-left+1 > maxLength {
				maxLength = right - left + 1
			}
			right += 1

		} else {
			delete(substringsLetters, rune(s[left]))
			left += 1
			if right-left+1 > maxLength {
				maxLength = right - left + 1
			}

		}
	}
	return maxLength
}
