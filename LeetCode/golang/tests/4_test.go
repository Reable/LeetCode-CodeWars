package tests

import (
	"golang/tasks"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			name:     "пример 1",
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2.0,
		},
		{
			name:     "пример 2",
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.5,
		},
		{
			name:     "оба нули",
			nums1:    []int{0, 0},
			nums2:    []int{0, 0},
			expected: 0.0,
		},
		{
			name:     "один пустой",
			nums1:    []int{},
			nums2:    []int{1},
			expected: 1.0,
		},
		{
			name:     "второй пустой",
			nums1:    []int{2},
			nums2:    []int{},
			expected: 2.0,
		},
		{
			name:     "разные размеры",
			nums1:    []int{1, 3, 5},
			nums2:    []int{2, 4, 6, 7},
			expected: 4.0,
		},
		{
			name:     "отрицательные числа",
			nums1:    []int{-5, -3, -1},
			nums2:    []int{-2, 0, 2},
			expected: -1.5,
		},
		{
			name:     "один элемент в каждом",
			nums1:    []int{1},
			nums2:    []int{2},
			expected: 1.5,
		},
		{
			name:     "все одинаковые",
			nums1:    []int{2, 2, 2},
			nums2:    []int{2, 2},
			expected: 2.0,
		},
		{
			name:     "длинные массивы",
			nums1:    []int{1, 2, 3, 4, 5},
			nums2:    []int{6, 7, 8, 9, 10},
			expected: 5.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tasks.FindMedianSortedArrays(tt.nums1, tt.nums2)

			if result != tt.expected {
				t.Errorf("nums1=%v nums2=%v: got %v, want %v",
					tt.nums1, tt.nums2, result, tt.expected)
			}
		})
	}
}
