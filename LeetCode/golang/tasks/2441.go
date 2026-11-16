package tasks

import (
	"math"
	"slices"
)

func FindMaxK2441_v1(nums []int) int {
	temp := make(map[int]int, len(nums))

	for _, n := range nums {
		temp[n] = n
	}

	maxInt := math.MinInt

	for _, n := range nums {
		if n > 0 {
			value, ok := temp[-n]
			if ok && n > maxInt {
				maxInt = value
			}
		} else if n < 0 {
			value, ok := temp[-n]
			if ok && value > maxInt {
				maxInt = value
			}
		}
	}

	if maxInt != math.MinInt {
		return maxInt
	}

	return -1
}

func FindMaxK2441_v2(nums []int) int {
	pos := make([]bool, 1001)
	result := 1

	for _, n := range nums {
		if n > 0 {
			pos[n] = true
		}
	}

	for _, n := range nums {
		if n < 0 && pos[-n] && n < result {
			result = n
		}
	}

	return -result
}

func FindMaxK2441_v3(nums []int) int {
	slices.Sort(nums)
	l := 0
	r := len(nums) - 1

	for l < r {
		if nums[r] == -nums[l] {
			return nums[r]
		} else if -nums[l] < nums[r] {
			r--
		} else {
			l++
		}
	}

	return -1
}
