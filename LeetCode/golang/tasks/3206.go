package tasks

func NumberOfAlternatingGroups3206(colors []int) int {
	n := len(colors)
	count := 0

	for i := range colors {
		a, b, c := colors[i], colors[(i+1)%n], colors[(i+2)%n]
		if a != b && b != c {
			count++
		}
	}

	return count
}
