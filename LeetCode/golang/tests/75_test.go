package tests

import (
	"golang/tasks"
	"slices"
	"testing"
)

func TestSortColors75(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "LeetCode example 1",
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "LeetCode example 2",
			nums: []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
		{
			name: "empty array",
			nums: []int{},
			want: []int{},
		},
		{
			name: "single element",
			nums: []int{1},
			want: []int{1},
		},
		{
			name: "all zeros",
			nums: []int{0, 0, 0, 0},
			want: []int{0, 0, 0, 0},
		},
		{
			name: "all ones",
			nums: []int{1, 1, 1, 1},
			want: []int{1, 1, 1, 1},
		},
		{
			name: "all twos",
			nums: []int{2, 2, 2, 2},
			want: []int{2, 2, 2, 2},
		},
		{
			name: "already sorted",
			nums: []int{0, 0, 1, 1, 2, 2},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "reverse sorted",
			nums: []int{2, 2, 1, 1, 0, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "duplicates",
			nums: []int{1, 2, 0, 2, 0, 1, 1, 0},
			want: []int{0, 0, 0, 1, 1, 1, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.nums...)

			tasks.SortColors75(nums)

			if !slices.Equal(nums, tt.want) {
				t.Errorf("SortColors75() = %v, want %v", nums, tt.want)
			}
		})
	}
}
