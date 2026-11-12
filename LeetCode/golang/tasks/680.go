package tasks

func checkPalindrome(text string, l, r int) bool {
	for l < r {
		if text[l] != text[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func ValidPalindrome680(s string) bool {
	if len(s) == 1 {
		return true
	}

	l := 0
	r := len(s) - 1

	for l < r {
		if s[l] != s[r] {
			return checkPalindrome(s, l+1, r) || checkPalindrome(s, l, r-1)
		}

		l++
		r--
	}

	return true
}
