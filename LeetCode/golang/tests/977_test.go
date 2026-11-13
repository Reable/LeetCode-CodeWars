package tests

import (
	"golang/tasks"
	"reflect"
	"testing"
)

func TestSortedSquares977(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "Example 2",
			nums: []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
		{
			name: "Empty array",
			nums: []int{},
			want: []int{},
		},
		{
			name: "Single positive number",
			nums: []int{5},
			want: []int{25},
		},
		{
			name: "Single negative number",
			nums: []int{-5},
			want: []int{25},
		},
		{
			name: "All positive numbers",
			nums: []int{1, 2, 3, 4},
			want: []int{1, 4, 9, 16},
		},
		{
			name: "All negative numbers",
			nums: []int{-5, -4, -3, -2},
			want: []int{4, 9, 16, 25},
		},
		{
			name: "All zeros",
			nums: []int{0, 0, 0},
			want: []int{0, 0, 0},
		},
		{
			name: "Mixed numbers with zero",
			nums: []int{-3, 0, 2},
			want: []int{0, 4, 9},
		},
		{
			name: "Equal absolute values",
			nums: []int{-3, 3},
			want: []int{9, 9},
		},
		{
			name: "Duplicate numbers",
			nums: []int{-2, -2, 2, 2},
			want: []int{4, 4, 4, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.SortedSquares977(tt.nums)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(
					"SortedSquares977(%v) = %v, want %v",
					tt.nums,
					got,
					tt.want,
				)
			}
		})
	}
}
