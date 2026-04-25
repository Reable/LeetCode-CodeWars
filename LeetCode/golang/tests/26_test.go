package tests

import (
	"reflect"
	"testing"

	"golang/tasks"
)

func TestRemoveDuplicates26(t *testing.T) {
	tests := []struct {
		name       string
		input      []int
		expectedK  int
		expected   []int // только первые k элементов
	}{
		{
			name:      "пример 1",
			input:     []int{1, 1, 2},
			expectedK: 2,
			expected:  []int{1, 2},
		},
		{
			name:      "пример 2",
			input:     []int{0,0,1,1,1,2,2,3,3,4},
			expectedK: 5,
			expected:  []int{0,1,2,3,4},
		},
		{
			name:      "пустой массив",
			input:     []int{},
			expectedK: 0,
			expected:  []int{},
		},
		{
			name:      "один элемент",
			input:     []int{5},
			expectedK: 1,
			expected:  []int{5},
		},
		{
			name:      "без дубликатов",
			input:     []int{1,2,3,4},
			expectedK: 4,
			expected:  []int{1,2,3,4},
		},
		{
			name:      "все одинаковые",
			input:     []int{7,7,7,7},
			expectedK: 1,
			expected:  []int{7},
		},
		{
			name:      "дубликаты в начале",
			input:     []int{1,1,1,2,3},
			expectedK: 3,
			expected:  []int{1,2,3},
		},
		{
			name:      "дубликаты в конце",
			input:     []int{1,2,3,3,3},
			expectedK: 3,
			expected:  []int{1,2,3},
		},
		{
			name:      "отрицательные числа",
			input:     []int{-3,-3,-2,-1,-1},
			expectedK: 3,
			expected:  []int{-3,-2,-1},
		},
		{
			name:      "чередующиеся дубликаты",
			input:     []int{1,1,2,2,3,3},
			expectedK: 3,
			expected:  []int{1,2,3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.input))
			copy(nums, tt.input)

			k := tasks.RemoveDuplicates26(nums)

			if k != tt.expectedK {
				t.Errorf("k: got %d, want %d", k, tt.expectedK)
			}

			if !reflect.DeepEqual(nums[:k], tt.expected) {
				t.Errorf("array: got %v, want %v", nums[:k], tt.expected)
			}
		})
	}
}