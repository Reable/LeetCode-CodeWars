package tests

import (
	"golang/tasks"
	"testing"
)

func TestCheckIfExist(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want bool
	}{
		{
			name: "Basic example",
			arr:  []int{10, 2, 5, 3},
			want: true,
		},
		{
			name: "Another basic example",
			arr:  []int{7, 1, 14, 11},
			want: true,
		},
		{
			name: "No double exists",
			arr:  []int{3, 1, 7, 11},
			want: false,
		},
		{
			name: "Two zeros",
			arr:  []int{0, 0},
			want: true,
		},
		{
			name: "Single zero",
			arr:  []int{0},
			want: false,
		},
		{
			name: "Zero and non-zero",
			arr:  []int{0, 1},
			want: false,
		},
		{
			name: "Negative numbers",
			arr:  []int{-10, -5, 2, 4},
			want: true,
		},
		{
			name: "Negative double",
			arr:  []int{-2, -4},
			want: true,
		},
		{
			name: "Odd and even numbers",
			arr:  []int{1, 3, 6, 12},
			want: true,
		},
		{
			name: "Empty array",
			arr:  []int{},
			want: false,
		},
		{
			name: "Single element",
			arr:  []int{5},
			want: false,
		},
		{
			name: "Duplicate values but no double",
			arr:  []int{3, 3},
			want: false,
		},
		{
			name: "Double appears before original",
			arr:  []int{8, 4},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.CheckIfExist1346(tt.arr)

			if got != tt.want {
				t.Errorf(
					"checkIfExist(%v) = %v, want %v",
					tt.arr,
					got,
					tt.want,
				)
			}
		})
	}
}
