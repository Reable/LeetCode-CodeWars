package tasks

import "fmt"

func FindTheDifference389_v1(s string, t string) byte {
	temp := make(map[rune]int)

	for _, c := range s {
		temp[c]++
	}
	for _, c := range t {
		_, ok := temp[c]
		if !ok || temp[c] == 0 {
			return byte(c)
		}
		temp[c]--
	}

	return t[0]
}

func FindTheDifference389_v2(s string, t string) string {
	temp := make(map[rune]int)
	for _, c := range s {
		temp[c]++
	}
	for _, c := range t {
		temp[c]--
	}
	fmt.Println(temp)
	for v, count := range temp {
		if count == -1 {
			return string(v)
		}
	}

	return string(t[0])
}
