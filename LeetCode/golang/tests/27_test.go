package tests

import (
	"golang/tasks"
	"sort"
	"testing"
)

func TestRemoveElement27(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		val       int
		expectedK int
		expected  []int
	}{
		{
			name:      "пример 1",
			nums:      []int{3, 2, 2, 3},
			val:       3,
			expectedK: 2,
			expected:  []int{2, 2},
		},
		{
			name:      "пример 2",
			nums:      []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:       2,
			expectedK: 5,
			expected:  []int{0, 1, 3, 0, 4},
		},
		{
			name:      "нет удалений",
			nums:      []int{1, 2, 3},
			val:       4,
			expectedK: 3,
			expected:  []int{1, 2, 3},
		},
		{
			name:      "все элементы удаляются",
			nums:      []int{5, 5, 5},
			val:       5,
			expectedK: 0,
			expected:  []int{},
		},
		{
			name:      "один элемент удаляется",
			nums:      []int{1},
			val:       1,
			expectedK: 0,
			expected:  []int{},
		},
		{
			name:      "один элемент остаётся",
			nums:      []int{1},
			val:       2,
			expectedK: 1,
			expected:  []int{1},
		},
		{
			name:      "пустой массив",
			nums:      []int{},
			val:       1,
			expectedK: 0,
			expected:  []int{},
		},
		{
			name:      "чередование",
			nums:      []int{1, 2, 1, 2, 1, 2},
			val:       1,
			expectedK: 3,
			expected:  []int{2, 2, 2},
		},
		{
			name:      "отрицательные числа",
			nums:      []int{-1, -2, -1, -3},
			val:       -1,
			expectedK: 2,
			expected:  []int{-2, -3},
		},
		{
			name:      "повторы и порядок не важен",
			nums:      []int{4, 1, 4, 2, 4, 3},
			val:       4,
			expectedK: 3,
			expected:  []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)

			k := tasks.RemoveElement27(nums, tt.val)

			if k != tt.expectedK {
				t.Errorf("k: got %d, want %d", k, tt.expectedK)
			}

			got := nums[:k]

			sort.Ints(got)
			expected := make([]int, len(tt.expected))
			copy(expected, tt.expected)
			sort.Ints(expected)

			if len(got) != len(expected) {
				t.Fatalf("length mismatch: got %v, want %v", got, expected)
			}

			for i := range got {
				if got[i] != expected[i] {
					t.Errorf("elements mismatch: got %v, want %v", got, expected)
					break
				}
			}
		})
	}
}
