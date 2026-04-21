package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 0, 1, 0, 0, 0, 0}
	result := [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}

	answer := threeSum(nums)
	fmt.Println("Correct answer: ", result, "-->", answer)
}

func threeSum(nums []int) [][]int {
	var results [][]int

	if len(nums) < 3 {
		return results
	}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l := i + 1
		r := len(nums) - 1

		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum == 0 {
				results = append(results, []int{nums[i], nums[l], nums[r]})

				for l < r && nums[l] == nums[l+1] {
					l++
				}

				for l < r && nums[r] == nums[r-1] {
					r--
				}

				l++
				r--
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return results
}
