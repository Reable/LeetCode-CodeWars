package tasks

import (
	"math"
)

func FirstUniqChar387_v1(s string) int {
	tempCount := make(map[rune]int)
	tempIdx := make(map[rune]int)

	for i, v := range s {
		tempCount[v]++
		tempIdx[v] = i
	}

	res := math.MaxInt

	for word, count := range tempCount {
		if count == 1 && tempIdx[word] < res {
			res = tempIdx[word]
		}
	}

	if res == math.MaxInt {
		return -1
	}

	return res
}

func FirstUniqChar387_v2(s string) int {
	bitmap := [26]int{}
	for _, v := range s {
		bitmap[v-'a']++
	}
	for i, v := range s {
		if bitmap[v-'a'] == 1 {
			return i
		}
	}
	return -1

}
