package tasks

func IsPrefixString1961(s string, words []string) bool {
	p := 0

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			if p >= len(s) || s[p] != word[i] {
				return false
			}
			p++
		}
		if p == len(s) {
			return true
		}
	}
	return false
}
