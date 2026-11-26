package tests

import (
	"golang/tasks"
	"testing"
)

func TestContainsDuplicate217(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{
			name:     "contains duplicate",
			nums:     []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "no duplicates",
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "multiple duplicates",
			nums:     []int{1, 1, 2, 2, 3, 3},
			expected: true,
		},
		{
			name:     "empty array",
			nums:     []int{},
			expected: false,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: false,
		},
		{
			name:     "all elements are equal",
			nums:     []int{5, 5, 5, 5},
			expected: true,
		},
		{
			name:     "negative numbers with duplicate",
			nums:     []int{-1, -2, -3, -1},
			expected: true,
		},
		{
			name:     "negative numbers without duplicate",
			nums:     []int{-1, -2, -3, -4},
			expected: false,
		},
		{
			name:     "zero is duplicated",
			nums:     []int{0, 1, 2, 0},
			expected: true,
		},
		{
			name:     "duplicate at the end",
			nums:     []int{1, 2, 3, 4, 5, 5},
			expected: true,
		},
		{
			name:     "duplicate at the beginning",
			nums:     []int{7, 7, 1, 2, 3},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.ContainsDuplicate217(tt.nums)

			if got != tt.expected {
				t.Errorf(
					"ContainsDuplicate217(%v) = %v; want %v",
					tt.nums,
					got,
					tt.expected,
				)
			}
		})
	}
}
