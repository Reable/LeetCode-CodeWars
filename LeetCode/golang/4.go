package main

import (
	"fmt"
)

func test4() {

	nums1 := [][]int{
		[]int{1, 3}, []int{1, 2},
		[]int{0, 0},
	}
	nums2 := [][]int{
		[]int{2}, []int{3, 4},
		[]int{0, 0},
	}

	results := []float64{2.0, 2.5, 0.0}

	for i := 0; i < len(nums1); i++ {
		result := findMedianSortedArrays(nums1[i], nums2[i])
		fmt.Println("Correct answer: ", result == results[i], result)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	arr := mergeSort(nums1, nums2)
	lengthArr := len(arr) - 1

	var median float64

	if lengthArr%2 == 0 {
		median = float64(arr[lengthArr/2])
	} else {
		l := float64(arr[lengthArr/2])
		r := float64(arr[lengthArr/2+1])
		median = (l + r) / 2
	}

	return median
}

func mergeSort(nums1, nums2 []int) []int {

	p1 := 0
	p2 := 0

	var result []int

	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			result = append(result, nums1[p1])
			result = append(result, nums2[p2])
			p1++
			p2++

		} else if nums1[p1] > nums2[p2] {
			result = append(result, nums2[p2])
			p2++
		} else {
			result = append(result, nums1[p1])
			p1++
		}
	}

	for p1 < len(nums1) {
		result = append(result, nums1[p1])
		p1++
	}

	for p2 < len(nums2) {
		result = append(result, nums2[p2])
		p2++
	}

	return result
}
