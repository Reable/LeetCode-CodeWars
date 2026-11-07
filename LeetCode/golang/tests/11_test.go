package tests

import (
	"golang/tasks"
	"testing"
)

func TestMaxArea11(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "базовый пример из условия",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "минимальная длина (2 элемента)",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "два элемента разной высоты",
			height:   []int{1, 2},
			expected: 1,
		},
		{
			name:     "все элементы одинаковые",
			height:   []int{4, 4, 4, 4, 4},
			expected: 16, // крайние 4 и 4: 4 * (4 - 0) = 16
		},
		{
			name:     "строго возрастающий массив",
			height:   []int{1, 2, 3, 4, 5},
			expected: 6, // 2 и 5: 2 * (4 - 1) = 6 или 3 и 5: 3 * (4 - 2) = 6
		},
		{
			name:     "строго убывающий массив",
			height:   []int{5, 4, 3, 2, 1},
			expected: 6, // 5 и 2: 2 * (3 - 0) = 6
		},
		{
			name:     "высокие линии по краям",
			height:   []int{10, 1, 1, 1, 1, 10},
			expected: 50, // 10 * (5 - 0) = 50
		},
		{
			name:     "высокие линии в самом центре",
			height:   []int{1, 2, 100, 100, 2, 1},
			expected: 100, // 100 * (3 - 2) = 100
		},
		{
			name:     "пирамида (пик в центре)",
			height:   []int{1, 3, 5, 7, 5, 3, 1},
			expected: 12, // 3 на индексах 1 и 5: 3 * (5 - 1) = 12
		},
		{
			name:     "массив с нулями",
			height:   []int{0, 2, 0},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tasks.MaxArea11(tt.height)

			if result != tt.expected {
				t.Errorf("input %v: got %d, want %d",
					tt.height, result, tt.expected)
			}
		})
	}
}
