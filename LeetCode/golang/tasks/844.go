package tasks

func BackspaceCompare844_v1(s string, t string) bool {
	resL := []byte{}
	resR := []byte{}

	countLB := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '#' {
			countLB++
		} else {
			if countLB > 0 {
				countLB--
			} else {
				resL = append(resL, s[i])
			}
		}
	}

	countRB := 0
	for i := len(t) - 1; i >= 0; i-- {
		if t[i] == '#' {
			countRB++
		} else {
			if countRB > 0 {
				countRB--
			} else {
				resR = append(resR, t[i])
			}
		}
	}

	return string(resL) == string(resR)
}

func BackspaceCompare844_v2(s string, t string) bool {
	p1 := len(s) - 1
	p2 := len(t) - 1

	for p1 >= 0 || p2 >= 0 {
		skipS := 0
		skipT := 0

		for p1 >= 0 {
			if s[p1] == '#' {
				skipS++
			} else if skipS > 0 {
				skipS--
			} else {
				break
			}
			p1--
		}

		for p2 >= 0 {
			if t[p2] == '#' {
				skipT++
			} else if skipT > 0 {
				skipT--
			} else {
				break
			}
			p2--
		}

		if p1 < 0 || p2 < 0 {
			return p1 == p2
		}

		if s[p1] != t[p2] {
			return false
		}

		p1--
		p2--
	}

	return true
}
