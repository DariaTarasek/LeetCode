package strings_slices

import (
	"slices"
	"strings"
)

func ConvertToTitle(columnNumber int) string {
	var res []string
	for columnNumber > 0 {
		columnNumber--
		ostatok := columnNumber % 26
		res = append(res, string(rune(ostatok)+'A'))
		columnNumber /= 26
	}
	slices.Reverse(res)
	return strings.Join(res, "")
}
