package strings_slices

import (
	"strconv"
)

func IsPalindromeNumber(x int) bool {
	var strX string
	if x < 0 {
		return false
	}
	for x > 0 {
		strX += strconv.Itoa(x % 10)
		x /= 10
	}
	i := 0
	j := len(strX) - 1
	for i < j {
		if strX[i] != strX[j] {
			return false
		}
		i++
		j--
	}
	return true
}
