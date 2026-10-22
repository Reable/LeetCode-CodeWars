package tasks

func isLetterByte(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}

func ReverseOnlyLetters917(s string) string {
	arrByte := []byte(s)

	l := 0
	r := len(arrByte) - 1

	var temp byte

	var lLetter bool
	var rLetter bool

	for l < r {

		lLetter = isLetterByte(arrByte[l])
		rLetter = isLetterByte(arrByte[r])

		if lLetter && rLetter {
			temp = arrByte[l]
			arrByte[l] = arrByte[r]
			arrByte[r] = temp
			l++
			r--
		}

		if !lLetter {
			l++
		}

		if !rLetter {
			r--
		}
	}

	return string(arrByte)
}
