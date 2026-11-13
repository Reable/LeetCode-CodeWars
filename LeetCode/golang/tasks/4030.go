package tasks

import (
	"fmt"
	"strings"
)

func IsPalindromic4030(s string) bool {
	str := make([]string, len(s))

	for i := 0; i < len(s); i++ {
		str[i] = fmt.Sprintf("%08b", s[i])
	}

	data := strings.Join(str, "")
	l := 0
	r := len(data) - 1

	for l < r {
		if data[l] != data[r] {
			return false
		}

		l++
		r--
	}

	return true
}
