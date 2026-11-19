package tasks

func MinimumRecolors2379(blocks string, k int) int {
	count := 0
	for i := range k {
		if blocks[i] != 'B' {
			count++
		}
	}

	minCount := count
	for r := k; r < len(blocks); r++ {
		if blocks[r-k] != 'B' {
			count--
		}

		if blocks[r] != 'B' {
			count++
		}

		if count < minCount {
			minCount = count
		}
	}

	return minCount
}
