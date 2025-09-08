package strings_slices

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	s = reg.ReplaceAllString(s, "")
	if len(s) <= 1 {
		return true
	}
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
