package tests

import (
	"golang/tasks"
	"testing"
)

func TestMinimumRecolors2379(t *testing.T) {
	tests := []struct {
		name     string
		blocks   string
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			blocks:   "WBBWWBBWBW",
			k:        7,
			expected: 3,
		},
		{
			name:     "Example 2",
			blocks:   "WBWBBBW",
			k:        2,
			expected: 0,
		},
		{
			name:     "All white",
			blocks:   "WWWW",
			k:        3,
			expected: 3,
		},
		{
			name:     "All black",
			blocks:   "BBBB",
			k:        3,
			expected: 0,
		},
		{
			name:     "Single white",
			blocks:   "BWB",
			k:        1,
			expected: 0,
		},
		{
			name:     "Single white window",
			blocks:   "WBWB",
			k:        1,
			expected: 0,
		},
		{
			name:     "Window covers entire string",
			blocks:   "WWBB",
			k:        4,
			expected: 2,
		},
		{
			name:     "Minimum window at the end",
			blocks:   "WWWWBBBB",
			k:        4,
			expected: 0,
		},
		{
			name:     "Minimum window at the beginning",
			blocks:   "BBBBWWWW",
			k:        4,
			expected: 0,
		},
		{
			name:     "Minimum window in the middle",
			blocks:   "WWBBBBWW",
			k:        4,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.MinimumRecolors2379(tt.blocks, tt.k)

			if got != tt.expected {
				t.Errorf(
					"MinimumRecolors2379(%q, %d) = %d; want %d",
					tt.blocks,
					tt.k,
					got,
					tt.expected,
				)
			}
		})
	}
}
