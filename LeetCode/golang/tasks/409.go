package tasks

func LongestPalindrome409(s string) int {
	temp := make(map[byte]int)
	for i := range s {
		temp[s[i]]++
	}

	res := 0
	hasOdd := false
	for _, count := range temp {
		if count%2 == 0 {
			res += count
		} else {
			res += count - 1
			hasOdd = true
		}
	}
	if hasOdd {
		res++
	}

	return res
}
