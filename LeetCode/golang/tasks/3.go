package tasks

func LengthOfLongestSubstring(s string) int {
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
