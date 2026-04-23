package tasks

func IsPalindrome9(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}

	reverse := 0

	for x > reverse {
		reverse = reverse*10 + x%10
		x /= 10
	}

	return x == reverse || x == reverse/10
}
