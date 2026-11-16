package tasks

import "unsafe"

func ReversePrefix2000(word string, ch byte) string {
	findRightIdx := -1
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			findRightIdx = i
			break
		}
	}

	if findRightIdx < 1 {
		return word
	}

	l := 0
	r := findRightIdx
	strWord := []byte(word)
	for l < r {
		strWord[l], strWord[r] = strWord[r], strWord[l]
		r--
		l++
	}
	return unsafe.String(unsafe.SliceData(strWord), len(strWord))
}
