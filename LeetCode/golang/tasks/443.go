package tasks

import (
	"strconv"
)

func Compress443(chars []byte) int {
	s := 0
	f := 0

	for f < len(chars) {
		temp := chars[f]
		count := 0

		for f < len(chars) && chars[f] == temp {
			f++
			count++
		}

		chars[s] = temp
		s++

		if count > 1 {
			for _, digit := range strconv.Itoa(count) {
				chars[s] = byte(digit)
				s++
			}
		}
	}

	return s
}
