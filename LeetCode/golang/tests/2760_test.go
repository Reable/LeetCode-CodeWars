package tests

import (
	"golang/tasks"
	"testing"
)

func TestLongestAlternatingSubarray2760(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		threshold int
		expected  int
	}{
		{
			name:      "Example 1",
			nums:      []int{3, 2, 5, 4},
			threshold: 5,
			expected:  3,
		},
		{
			name:      "Example 2",
			nums:      []int{1, 2},
			threshold: 2,
			expected:  1,
		},
		{
			name:      "Example 3",
			nums:      []int{2, 3, 4, 5},
			threshold: 4,
			expected:  3,
		},
		{
			name:      "All even",
			nums:      []int{2, 4, 6, 8},
			threshold: 10,
			expected:  1,
		},
		{
			name:      "All odd",
			nums:      []int{1, 3, 5, 7},
			threshold: 10,
			expected:  0,
		},
		{
			name:      "Perfect alternating",
			nums:      []int{2, 3, 4, 5, 6, 7},
			threshold: 10,
			expected:  6,
		},
		{
			name:      "Threshold breaks sequence",
			nums:      []int{2, 3, 8, 5, 6},
			threshold: 5,
			expected:  2,
		},
		{
			name:      "Single even",
			nums:      []int{4},
			threshold: 5,
			expected:  1,
		},
		{
			name:      "Single odd",
			nums:      []int{3},
			threshold: 5,
			expected:  0,
		},
		{
			name:      "First element exceeds threshold",
			nums:      []int{10, 2, 3, 4},
			threshold: 5,
			expected:  3,
		},
		{
			name:      "Alternating starts in the middle",
			nums:      []int{1, 1, 2, 3, 4, 5},
			threshold: 5,
			expected:  4,
		},
		{
			name:      "Zero values",
			nums:      []int{0, 1, 2, 3},
			threshold: 3,
			expected:  4,
		},
		{
			name:      "No valid elements",
			nums:      []int{7, 9, 11},
			threshold: 5,
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.LongestAlternatingSubarray2760(
				tt.nums,
				tt.threshold,
			)

			if got != tt.expected {
				t.Errorf(
					"LongestAlternatingSubarray2760(%v, %d) = %d; want %d",
					tt.nums,
					tt.threshold,
					got,
					tt.expected,
				)
			}
		})
	}
}
