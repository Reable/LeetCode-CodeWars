package tasks

func CountKConstraintSubstrings3258(s string, k int) int {
	l := 0
	r := 0
	count := 0
	count0 := 0
	count1 := 0

	for r < len(s) {

		if s[r] == '0' {
			count0++
		} else {
			count1++
		}

		for l < len(s) && count0 > k && count1 > k {
			if s[l] == '1' {
				count1--
			} else {
				count0--
			}
			l++
		}

		count += r - l + 1
		r++
	}

	return count
}
