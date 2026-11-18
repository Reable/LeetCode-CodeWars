package tasks

import (
	"strconv"
)

func DivisorSubstrings2269(num int, k int) int {
	strFull := strconv.Itoa(num)

	if k > len(strFull) {
		return 0
	}

	count := 0

	for l := 0; l <= len(strFull)-k; l++ {
		n, _ := strconv.Atoi(strFull[l : l+k])
		if n != 0 && num%n == 0 {
			count++
		}
	}

	return count
}
