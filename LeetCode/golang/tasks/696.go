package tasks

func CountBinarySubstrings696(s string) int {
	count := 0
	prev := 0
	curr := 1

	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			curr++
		} else {
			count += min(curr, prev)
			prev = curr
			curr = 1
		}
	}

	count += min(curr, prev)

	return count
}
