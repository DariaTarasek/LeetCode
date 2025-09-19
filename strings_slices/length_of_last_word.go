package strings_slices

func LengthOfLastWord(s string) int {
	i := len(s) - 1
	count := 0
	flag := 0
	for i >= 0 {
		if s[i] == ' ' && flag == 0 {
			i--
		} else if s[i] != ' ' {
			count += 1
			flag = 1
			i--
		} else {
			return count
		}
	}
	return count
}
