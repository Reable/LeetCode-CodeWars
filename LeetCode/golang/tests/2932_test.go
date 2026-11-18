package tests

import (
	"golang/tasks"
	"testing"
)

func TestMaximumStrongPairXor2932(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 4, 5},
			want: 7,
		},
		{
			name: "Example 2",
			nums: []int{1, 1, 10, 3, 9},
			want: 3,
		},
		{
			name: "Example 3",
			nums: []int{5, 6, 25, 30},
			want: 7,
		},
		{
			name: "Single element",
			nums: []int{5},
			want: 0,
		},
		{
			name: "All equal elements",
			nums: []int{7, 7, 7, 7},
			want: 0,
		},
		{
			name: "Two elements",
			nums: []int{1, 2},
			want: 3,
		},
		{
			name: "No valid distinct pair",
			nums: []int{1, 10},
			want: 0,
		},
		{
			name: "Maximum XOR with zero",
			nums: []int{0, 0, 1},
			want: 0,
		},
		{
			name: "Unsorted array",
			nums: []int{10, 1, 5, 3, 9},
			want: 15,
		},
		{
			name: "Powers of two",
			nums: []int{1, 2, 4, 8, 16},
			want: 24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.MaximumStrongPairXor2932(tt.nums)

			if got != tt.want {
				t.Errorf(
					"MaximumStrongPairXor2932(%v) = %d; want %d",
					tt.nums,
					got,
					tt.want,
				)
			}
		})
	}
}
