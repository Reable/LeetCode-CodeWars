package tests

import (
	"golang/tasks"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
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
			name: "two elements",
			nums: []int{1, 2},
			want: []int{1, 2},
		},
		{
			name: "one majority element",
			nums: []int{3, 2, 3},
			want: []int{3},
		},
		{
			name: "two majority elements",
			nums: []int{1, 1, 1, 3, 3, 2, 2, 2},
			want: []int{1, 2},
		},
		{
			name: "no majority elements",
			nums: []int{1, 2, 3, 4, 5, 6},
			want: []int{},
		},
		{
			name: "all elements are equal",
			nums: []int{7, 7, 7, 7, 7},
			want: []int{7},
		},
		{
			name: "element occurs exactly n/3 times",
			nums: []int{1, 1, 2, 2, 3, 3},
			want: []int{},
		},
		{
			name: "element occurs n/3 plus one times",
			nums: []int{1, 1, 1, 2, 3, 4},
			want: []int{1},
		},
		{
			name: "negative numbers",
			nums: []int{-1, -1, -1, 2, 3, 4},
			want: []int{-1},
		},
		{
			name: "zero as majority element",
			nums: []int{0, 0, 0, 1, 2, 3},
			want: []int{0},
		},
		{
			name: "two candidates with different frequencies",
			nums: []int{1, 1, 1, 1, 2, 2, 2, 3, 4},
			want: []int{1},
		},
		{
			name: "large array with one majority",
			nums: append(
				make([]int, 100),
				make([]int, 0)...,
			),
			want: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.MajorityElement229(tt.nums)

			// Порядок элементов не имеет значения.
			if !equalAsSet(got, tt.want) {
				t.Errorf("majorityElement(%v) = %v, want %v",
					tt.nums, got, tt.want)
			}
		})
	}
}

func equalAsSet(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[int]int)

	for _, v := range a {
		counts[v]++
	}

	for _, v := range b {
		counts[v]--
		if counts[v] < 0 {
			return false
		}
	}

	return true
}
