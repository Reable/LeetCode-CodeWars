package tasks

import "fmt"

func IsIsomorphic205(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mappsST := make(map[rune]rune)

	for i := range s {
		v1 := rune(s[i])
		v2 := rune(t[i])
		if value, ok := mappsST[v1]; ok {
			if value != v2 {
				return false
			}
		} else {
			mappsST[v1] = v2
		}
	}

	mappsTS := make(map[rune]rune)
	for i := range t {
		v1 := rune(t[i])
		v2 := rune(s[i])
		if value, ok := mappsTS[v1]; ok {
			if value != v2 {
				return false
			}
		} else {
			mappsTS[v1] = v2
		}
	}

	fmt.Println(mappsST, mappsTS)
	for i := range mappsST {
		v := mappsST[i]
		if i != mappsTS[v] {
			return false
		}
	}

	return true
}
