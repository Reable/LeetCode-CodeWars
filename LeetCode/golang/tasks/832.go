package tasks

func FlipAndInvertImage832(image [][]int) [][]int {

	for i := range image {
		l := 0
		r := len(image[i]) - 1

		for l <= r {
			image[i][l], image[i][r] = 1-image[i][r], 1-image[i][l]
			l++
			r--
		}
	}

	return image
}
