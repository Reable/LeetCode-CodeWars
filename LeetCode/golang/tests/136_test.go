package tests

import (
	"golang/tasks"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1",
			nums: []int{2, 2, 1},
			want: 1,
		},
		{
			name: "example 2",
			nums: []int{4, 1, 2, 1, 2},
			want: 4,
		},
		{
			name: "single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "negative number",
			nums: []int{-1, 2, 2},
			want: -1,
		},
		{
			name: "all negative numbers",
			nums: []int{-3, -2, -2},
			want: -3,
		},
		{
			name: "zero as unique element",
			nums: []int{1, 1, 0},
			want: 0,
		},
		{
			name: "unique element at the end",
			nums: []int{5, 5, 3, 3, 7},
			want: 7,
		},
		{
			name: "unique element at the beginning",
			nums: []int{9, 4, 4, 2, 2},
			want: 9,
		},
		{
			name: "large numbers",
			nums: []int{1000000, 5, 5},
			want: 1000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.SingleNumber136(tt.nums)

			if got != tt.want {
				t.Errorf(
					"singleNumber(%v) = %d, want %d",
					tt.nums,
					got,
					tt.want,
				)
			}
		})
	}
}
