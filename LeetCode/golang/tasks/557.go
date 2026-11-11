package tasks

import (
	"strings"
	"unsafe"
)

func ReverseWords557(s string) string {
	str := strings.Split(s, " ")
	res := make([]string, 0, len(s))

	for _, v := range str {
		result := []byte(v)
		l := 0
		r := len(result) - 1
		for l < r {
			result[l], result[r] = result[r], result[l]
			l++
			r--
		}

		res = append(res, unsafe.String(unsafe.SliceData(result), len(result)))
	}
	return strings.Join(res, " ")
}
