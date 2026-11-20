package tasks

func MaximumLengthSubstring3090_v1(s string) int {
	count := 0

	for l := 0; l < len(s); l++ {
		size := 1
		temp := map[byte]int{s[l]: 1}

		r := l + 1

		for r < len(s) {
			if temp[s[r]] == 2 {
				break
			}
			temp[s[r]]++
			size++
			r++
		}
		if size > count {
			count = size
		}
	}

	return count
}

func MaximumLengthSubstring3090_v2(s string) int {
	temp := make(map[byte]int, 26)

	l := 0
	count := 0

	for r := 0; r < len(s); r++ {
		temp[s[r]]++

		for temp[s[r]] > 2 {
			temp[s[l]]--
			l++
		}

		size := r - l + 1

		if size > count {
			count = size
		}
	}

	return count
}
