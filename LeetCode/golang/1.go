package main

import "fmt"

func test1() {
	nums := [][]int{
		{2, 7, 11, 15},
		{3, 2, 4},
		{3, 3},
	}

	targets := []int{9, 6, 6}

	i := 0

	for i < len(nums) {
		result := twoSum1WithMap(nums[i], targets[i])
		fmt.Println(result)

		i++
	}
}

func twoSum1WithMap(nums []int, target int) []int {
	temp := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		val := target - nums[i]

		_, hasV := temp[val]

		if hasV {
			return []int{temp[val], i}
		} else {
			temp[nums[i]] = i
		}
	}

	return []int{-1, -1}
}

// Only for sorted array
func twoSum1WithTwoPointer(nums []int, target int) []int {

	l, r := 0, len(nums)-1

	for l < r {
		temp := nums[l] + nums[r]

		if target == temp {
			return []int{l, r}
		}

		if temp > target {
			r--
		} else {
			l++
		}

	}

	return []int{-1, -1}
}
