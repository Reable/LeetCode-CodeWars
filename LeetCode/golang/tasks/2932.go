package tasks

import (
	"math"
)

func MaximumStrongPairXor2932(nums []int) int {
	xorMax := math.MinInt

	for i := range nums {
		for j := range nums {
			if abs(nums[i]-nums[j]) <= min(nums[i], nums[j]) {
				xorMax = max(xorMax, nums[i]^nums[j])
			}
		}
	}

	return xorMax
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
