package tests

import (
	"golang/tasks"
	"testing"
)

func TestMaxDistance2078(t *testing.T) {
	tests := []struct {
		name     string
		colors   []int
		expected int
	}{
		{
			name:     "пример 1",
			colors:   []int{1, 1, 1, 6, 1, 1, 1},
			expected: 3,
		},
		{
			name:     "пример 2",
			colors:   []int{1, 8, 3, 8, 3},
			expected: 4,
		},
		{
			name:     "все одинаковые",
			colors:   []int{5, 5, 5, 5},
			expected: 0,
		},
		{
			name:     "два элемента разные",
			colors:   []int{1, 2},
			expected: 1,
		},
		{
			name:     "два элемента одинаковые",
			colors:   []int{3, 3},
			expected: 0,
		},
		{
			name:     "максимум на концах",
			colors:   []int{1, 2, 3, 4, 5, 6},
			expected: 5,
		},
		{
			name:     "обратный порядок",
			colors:   []int{6, 5, 4, 3, 2, 1},
			expected: 5,
		},
		{
			name:     "один элемент",
			colors:   []int{1},
			expected: 0,
		},
		{
			name:     "пустой массив",
			colors:   []int{},
			expected: 0,
		},
		{
			name:     "повторяющиеся группы",
			colors:   []int{1, 2, 1, 2, 1, 2},
			expected: 5,
		},
		{
			name:     "максимум внутри",
			colors:   []int{7, 7, 7, 1, 7, 7, 2},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tasks.MaxDistance2078(tt.colors)

			if result != tt.expected {
				t.Errorf("colors=%v: got %d, want %d",
					tt.colors, result, tt.expected)
			}
		})
	}
}
