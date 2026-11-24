package tests

import (
	"golang/tasks"
	"testing"
)

func TestMaxLength3411(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "single element",
			nums: []int{5},
			want: 0,
		},
		{
			name: "two coprime numbers",
			nums: []int{2, 3},
			want: 2,
		},
		{
			name: "two numbers with common divisor",
			nums: []int{2, 4},
			want: 2,
		},
		{
			name: "all ones",
			nums: []int{1, 1, 1, 1},
			want: 4,
		},
		{
			name: "mixed ones and primes",
			nums: []int{1, 2, 1, 2, 1, 1, 1},
			want: 5,
		},
		{
			name: "increasing numbers",
			nums: []int{2, 3, 5, 7},
			want: 4,
		},
		{
			name: "repeated numbers",
			nums: []int{3, 3, 3, 3},
			want: 2,
		},
		{
			name: "answer in the middle",
			nums: []int{6, 2, 3, 5, 7, 4},
			want: 4,
		},
		{
			name: "answer at the beginning",
			nums: []int{2, 3, 5, 6, 6},
			want: 3,
		},
		{
			name: "answer at the end",
			nums: []int{6, 6, 2, 3, 5},
			want: 3,
		},
		{
			name: "empty array",
			nums: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.MaxLength3411(tt.nums)

			if got != tt.want {
				t.Errorf(
					"MaxLength3411(%v) = %d, want %d",
					tt.nums,
					got,
					tt.want,
				)
			}
		})
	}
}
