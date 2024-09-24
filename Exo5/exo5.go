package exo

func Ft_max_substring(s string) int {
	lastIndex := make(map[byte]int)
	start := 0
	maxLength := 0

	for i := 0; i < len(s); i++ {
		if idx, found := lastIndex[s[i]]; found && idx >= start {
			start = idx + 1
		}
		lastIndex[s[i]] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	return maxLength
}
