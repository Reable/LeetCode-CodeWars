package main

import (
	"fmt"
)

func main() {
	strings := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
	}
	results := []int{3, 1, 3}

	i := 0

	for i < len(strings) {
		result := lengthOfLongestSubstring(strings[i])
		fmt.Println("Correct answer: ", result == results[i], result)

		i++
	}
}

func lengthOfLongestSubstring(s string) int {
	left := 0
	maxLen := 0

	temp := make(map[byte]int, len(s))

	for right := 0; right < len(s); right++ {
		char := s[right]

		if i, has := temp[char]; has && i >= left {
			left = i + 1
		}

		temp[char] = right

		curLen := right - left + 1
		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}
