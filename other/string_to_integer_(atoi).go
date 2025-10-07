package other

import "strings"

func myAtoi(s string) int {
	const (
		intMax = 1<<31 - 1
		intMin = -1 << 31
	)

	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	sign := 1
	i := 0
	n := len(s)

	if s[i] == '+' {
		i++
	} else if s[i] == '-' {
		sign = -1
		i++
	}

	num := 0

	for i < n {
		r := rune(s[i])
		if r < '0' || r > '9' {
			break
		}
		digit := int(r - '0')

		if num > (intMax-digit)/10 {
			if sign == 1 {
				return intMax
			}
			return intMin
		}

		num = num*10 + digit
		i++
	}

	return num * sign
}
