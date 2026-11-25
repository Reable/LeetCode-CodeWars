package tests

import (
	"golang/tasks"
	"testing"
)

func TestMajorityElement169(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "majority element in the middle",
			nums:     []int{3, 2, 3},
			expected: 3,
		},
		{
			name:     "majority element at the beginning",
			nums:     []int{2, 2, 1, 1, 1, 2, 2},
			expected: 2,
		},
		{
			name:     "single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "all elements are equal",
			nums:     []int{5, 5, 5, 5},
			expected: 5,
		},
		{
			name:     "no majority element",
			nums:     []int{1, 2, 3, 4},
			expected: -1,
		},
		{
			name:     "majority element at the end",
			nums:     []int{1, 2, 2, 2, 2},
			expected: 2,
		},
		{
			name:     "negative numbers",
			nums:     []int{-1, -1, -1, 2, 3},
			expected: -1,
		},
		{
			name:     "exactly half is not majority",
			nums:     []int{1, 1, 2, 2},
			expected: -1,
		},
		{
			name:     "majority element with negative value",
			nums:     []int{-5, -5, -5, 2, 3},
			expected: -5,
		},
		{
			name:     "majority element after candidate changes",
			nums:     []int{3, 3, 4, 2, 3, 3, 5},
			expected: 3,
		},
		{
			name:     "empty array",
			nums:     []int{},
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.MajorityElement169(tt.nums)

			if got != tt.expected {
				t.Errorf(
					"MajorityElement169(%v) = %d; want %d",
					tt.nums,
					got,
					tt.expected,
				)
			}
		})
	}
}
