package tests

import (
	"golang/tasks"
	"testing"
)

func TestNumberOfAlternatingGroups3206(t *testing.T) {
	tests := []struct {
		name   string
		colors []int
		want   int
	}{
		{
			name:   "Example 1",
			colors: []int{0, 1, 0, 0, 1},
			want:   3,
		},
		{
			name:   "All alternating",
			colors: []int{0, 1, 0},
			want:   1,
		},
		{
			name:   "No alternating groups",
			colors: []int{0, 0, 0, 0},
			want:   0,
		},
		{
			name:   "Alternating with circular boundary",
			colors: []int{1, 0, 1, 0},
			want:   4,
		},
		{
			name:   "Single color change",
			colors: []int{0, 1, 1},
			want:   1,
		},
		{
			name:   "Two colors grouped",
			colors: []int{0, 0, 1, 1},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.NumberOfAlternatingGroups3206(tt.colors)

			if got != tt.want {
				t.Errorf(
					"NumberOfAlternatingGroups3206(%v) = %d; want %d",
					tt.colors,
					got,
					tt.want,
				)
			}
		})
	}
}
