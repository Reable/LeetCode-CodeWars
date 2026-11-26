package tasks

func CanAliceWin3360(n int) bool {
	if n < 10 {
		return false
	}
	n -= 10
	count := 9

	aliceWin := true

	for n-count >= 0 && count >= 0 {
		n -= count
		count--
		aliceWin = !aliceWin
	}

	return aliceWin
}
