package tasks

import (
	"math"
)

func MinimumSumSubarray3364(nums []int, l int, r int) int {
	n := len(nums)

	prefix := make([]int, n+1)

	for i := range n {
		prefix[i+1] = prefix[i] + nums[i]
	}

	result := math.MaxInt

	for j := range n {
		minI := max(0, j-r+1)
		maxI := j - l + 1

		for i := minI; i <= maxI; i++ {
			currentSum := prefix[j+1] - prefix[i]
			if currentSum > 0 && currentSum < result {
				result = currentSum
			}
		}
	}

	if result == math.MaxInt {
		return -1
	}

	return result
}
