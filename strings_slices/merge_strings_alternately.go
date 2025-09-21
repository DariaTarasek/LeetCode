package strings_slices

func MergeAlternately(word1 string, word2 string) string {
	minLen := min(len(word1), len(word2))
	var res string
	for i := 0; i < minLen; i++ {
		res += string(word1[i])
		res += string(word2[i])
	}
	if len(word1) > len(word2) {
		res += word1[minLen:]
	} else {
		res += word2[minLen:]
	}
	return res
}
