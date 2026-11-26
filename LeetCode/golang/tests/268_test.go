package tests

import (
	"golang/tasks"
	"testing"
)

func TestMissingNumber268(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "missing number in the middle",
			nums:     []int{3, 0, 1},
			expected: 2,
		},
		{
			name:     "missing number at the end",
			nums:     []int{0, 1},
			expected: 2,
		},
		{
			name:     "missing number at the beginning",
			nums:     []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "single element zero",
			nums:     []int{0},
			expected: 1,
		},
		{
			name:     "single element one",
			nums:     []int{1},
			expected: 0,
		},
		{
			name:     "empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "missing number in larger array",
			nums:     []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			expected: 8,
		},
		{
			name:     "missing zero",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "missing maximum number",
			nums:     []int{0, 1, 2, 3, 4},
			expected: 5,
		},
		{
			name:     "missing number near the beginning",
			nums:     []int{0, 1, 3, 4, 5},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.MissingNumber268(tt.nums)

			if got != tt.expected {
				t.Errorf(
					"MissingNumber268(%v) = %d; want %d",
					tt.nums,
					got,
					tt.expected,
				)
			}
		})
	}
}
