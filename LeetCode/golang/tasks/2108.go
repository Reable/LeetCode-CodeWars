package tasks

func FirstPalindrome2108_v1(words []string) string {
	for _, word := range words {
		l := 0
		r := len(word) - 1
		for l < r {
			if word[l] != word[r] {
				break
			}
			r--
			l++
		}
		if l >= r {
			return word
		}
	}
	return ""
}

func FirstPalindrome2108_v2(words []string) string {
	for _, word := range words {
		l := 0
		r := len(word) - 1
		for l < r && word[l] == word[r] {
			r--
			l++
		}
		if l >= r {
			return word
		}
	}
	return ""
}
