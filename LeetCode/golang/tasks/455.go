package tasks

import "slices"

func FindContentChildren455(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)

	i, j := 0, 0

	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}

		j++
	}

	return i
}
