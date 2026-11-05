package tasks

import (
	"unsafe"
)

func ReverseStr541_v1(s string, k int) string {
	b := []byte(s)
	length := len(b)
	temp := k
	l := 0
	r := k - 1

	for r < length {
		for l < r {
			b[l], b[r] = b[r], b[l]
			l++
			r--
		}
		l = temp + k
		r = l + k - 1

		temp = l + k
	}

	r = length - 1
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}

	return unsafe.String(unsafe.SliceData(b), len(b))
}

func ReverseStr541_v2(s string, k int) string {
	b := []byte(s)

	for s := 0; s < len(b); s += 2 * k {
		l := s
		r := min(s+k, len(b)) - 1

		for l < r {
			b[l], b[r] = b[r], b[l]
			l++
			r--
		}
	}

	return unsafe.String(unsafe.SliceData(b), len(b))
}
