package tasks

func CheckString2124(s string) bool {
	iA := -1
	iB := -1

	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			if iA < i {
				iA = i
			}
		} else {
			if iB == -1 {
				iB = i
			}
		}
	}

	if iA > iB && iA != -1 && iB != -1 {
		return false
	}

	return true
}
