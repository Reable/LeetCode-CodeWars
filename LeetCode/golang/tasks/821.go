package tasks

import (
	"math"
)

func ShortestToChar821(s string, c byte) []int {
	result := make([]int, 0, len(s))
	ids := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			ids = append(ids, i)
		}
	}

	for i := 0; i < len(s); i++ {
		temp := math.MaxInt

		for j := 0; j < len(ids); j++ {
			sum := int(math.Abs(float64(ids[j]) - float64(i)))
			if temp > sum {
				temp = sum
			}
		}
		result = append(result, temp)
	}

	return result
}
