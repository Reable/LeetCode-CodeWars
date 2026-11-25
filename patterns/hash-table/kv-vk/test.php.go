package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "ABBBCCAABCC"

	counts := make(map[rune]int)
	for _, char := range str {
		counts[char]++
	}

	freqList := make([][]rune, len(str)+1)
	for char, freq := range counts {
		freqList[freq] = append(freqList[freq], char)
	}

	result := strings.Builder{}
	for freq := len(freqList) - 1; freq > 0; freq-- {
		for _, char := range freqList[freq] {
			for i := 0; i < freq; i++ {
				result.WriteRune(char)
			}
		}
	}

	fmt.Println(result.String())
}
