package tests

import (
	"golang/tasks"
	"testing"
)

func TestFindMaxK2441(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "First example",
			nums: []int{-1, 2, -3, 3},
			want: 3,
		},
		{
			name: "Second example",
			nums: []int{-1, 10, 6, 7, -7, 1},
			want: 7,
		},
		{
			name: "No matching pair",
			nums: []int{-10, 8, 6, 7, -2, -3},
			want: -1,
		},
		{
			name: "Multiple matching pairs",
			nums: []int{-1, 1, -2, 2, -3, 3},
			want: 3,
		},
		{
			name: "Only positive numbers",
			nums: []int{1, 2, 3, 4},
			want: -1,
		},
		{
			name: "Only negative numbers",
			nums: []int{-1, -2, -3, -4},
			want: -1,
		},
		{
			name: "Single matching pair",
			nums: []int{-5, 5},
			want: 5,
		},
		{
			name: "Zero in array",
			nums: []int{0, -1, 1},
			want: 1,
		},
		{
			name: "Only zero",
			nums: []int{0},
			want: -1,
		},
		{
			name: "Empty array",
			nums: []int{},
			want: -1,
		},
		{
			name: "Maximum allowed values",
			nums: []int{-1000, 1000},
			want: 1000,
		},
		{
			name: "Minimum matching pair",
			nums: []int{-1, 1, -1000},
			want: 1,
		},
		{
			name: "Unordered matching pairs",
			nums: []int{8, -8, 3, -3, 5, -5},
			want: 8,
		},
		{
			name: "Large positive without pair",
			nums: []int{-2, 2, -5, 1000},
			want: 2,
		},
	}

	functions := []struct {
		name string
		fn   func([]int) int
	}{
		{"v1", tasks.FindMaxK2441_v1},
		{"v2", tasks.FindMaxK2441_v2},
		{"v3", tasks.FindMaxK2441_v3},
	}

	for _, tt := range tests {
		for _, f := range functions {
			t.Run(tt.name+"/"+f.name, func(t *testing.T) {
				// Копируем массив, потому что v3 сортирует его.
				nums := append([]int(nil), tt.nums...)

				got := f.fn(nums)

				if got != tt.want {
					t.Errorf(
						"%s(%v) = %d, want %d",
						f.name,
						tt.nums,
						got,
						tt.want,
					)
				}
			})
		}
	}
}
