package tests

import (
	"golang/tasks"
	"testing"
)

func TestMinimumSumSubarray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		l    int
		r    int
		want int
	}{
		{
			name: "basic positive numbers",
			nums: []int{1, 2, 3, 4},
			l:    2,
			r:    3,
			want: 3,
		},
		{
			name: "negative numbers",
			nums: []int{-1, 2, -3, 4},
			l:    1,
			r:    2,
			want: 1,
		},
		{
			name: "all negative numbers",
			nums: []int{-1, -2, -3},
			l:    1,
			r:    3,
			want: -1,
		},
		{
			name: "no positive sum",
			nums: []int{-5, 1, -2},
			l:    2,
			r:    3,
			want: -1,
		},
		{
			name: "single element",
			nums: []int{5},
			l:    1,
			r:    1,
			want: 5,
		},
		{
			name: "minimum length",
			nums: []int{3, -2, 1},
			l:    1,
			r:    1,
			want: 1,
		},
		{
			name: "maximum length",
			nums: []int{2, -1, 3},
			l:    3,
			r:    3,
			want: 4,
		},
		{
			name: "zero sum",
			nums: []int{1, -1, 2},
			l:    2,
			r:    2,
			want: 1,
		},
		{
			name: "repeated numbers",
			nums: []int{2, 2, 2, 2},
			l:    2,
			r:    3,
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.MinimumSumSubarray3364(tt.nums, tt.l, tt.r)

			if got != tt.want {
				t.Errorf(
					"minimumSumSubarray(%v, %d, %d) = %d; want %d",
					tt.nums, tt.l, tt.r, got, tt.want,
				)
			}
		})
	}
}
