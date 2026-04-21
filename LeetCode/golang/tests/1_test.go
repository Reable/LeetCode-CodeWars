package tests

import (
	"reflect"
	"testing"

	"golang/tasks"
)

func TestTwoSum1WithMap(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "обычный случай",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "порядок неважен",
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			name:     "два одинаковых числа",
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
		{
			name:     "отрицательные числа",
			nums:     []int{-1, -2, -3, -4, -5},
			target:   -8,
			expected: []int{2, 4},
		},
		{
			name:     "решение в конце",
			nums:     []int{1, 2, 3, 4, 6},
			target:   10,
			expected: []int{3, 4},
		},
		{
			name:     "нет решения",
			nums:     []int{1, 2, 3},
			target:   7,
			expected: []int{-1, -1},
		},
		{
			name:     "один элемент",
			nums:     []int{5},
			target:   5,
			expected: []int{-1, -1},
		},
		{
			name:     "пустой массив",
			nums:     []int{},
			target:   0,
			expected: []int{-1, -1},
		},
		{
			name:     "ноль в массиве",
			nums:     []int{0, 4, 3, 0},
			target:   0,
			expected: []int{0, 3},
		},
		{
			name:     "большие числа",
			nums:     []int{1000000, 500000, 500000},
			target:   1000000,
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tasks.TwoSum1WithMap(tt.nums, tt.target)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
