package tasks

import (
	"strings"
)

func WordPattern290(pattern string, s string) bool {
	words := strings.Fields(s)

	if len(pattern) != len(words) {
		return false
	}

	patternToWord := make(map[byte]string)
	wordToPattern := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		char := pattern[i]
		word := words[i]

		if val, ok := patternToWord[char]; ok {
			if val != word {
				return false
			}
		} else {
			patternToWord[char] = word
		}

		if val, ok := wordToPattern[word]; ok {
			if val != char {
				return false
			}
		} else {
			wordToPattern[word] = char
		}
	}

	return true
}
