package strings_slices

import "slices"

func GroupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	var ans [][]string
	for _, str := range strs {
		runes := []rune(str)
		slices.Sort(runes)
		anagrams[string(runes)] = append(anagrams[string(runes)], str)
	}
	for _, v := range anagrams {
		ans = append(ans, v)
	}
	return ans
}
