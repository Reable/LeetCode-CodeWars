package tasks

func LongestPalindrome5(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, maxLen := 0, 1

	for i := 0; i < len(s); i++ {
		l1 := find(s, i, i)
		l2 := find(s, i, i+1)

		current := max(l1, l2)

		if current > maxLen {
			maxLen = current
			start = i - (current-1)/2
		}
	}

	return s[start : start+maxLen]
}

func find(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

func max(l1, l2 int) int {
	if l1 > l2 {
		return l1
	}
	return l2
}
