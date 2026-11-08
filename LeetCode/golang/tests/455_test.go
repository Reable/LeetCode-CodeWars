package tests

import (
	"golang/tasks"
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	tests := []struct {
		name     string
		g        []int
		s        []int
		expected int
	}{
		{
			name:     "базовый пример 1",
			g:        []int{1, 2, 3},
			s:        []int{1, 1},
			expected: 1,
		},
		{
			name:     "базовый пример 2",
			g:        []int{1, 2},
			s:        []int{1, 2, 3},
			expected: 2,
		},
		{
			name:     "пустой список детей",
			g:        []int{},
			s:        []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "пустой список печенья",
			g:        []int{1, 2, 3},
			s:        []int{},
			expected: 0,
		},
		{
			name:     "оба слайса пустые",
			g:        []int{},
			s:        []int{},
			expected: 0,
		},
		{
			name:     "одно печенье подходит одному ребёнку",
			g:        []int{5},
			s:        []int{5},
			expected: 1,
		},
		{
			name:     "печенье меньше потребности ребёнка",
			g:        []int{5},
			s:        []int{3},
			expected: 0,
		},
		{
			name:     "печенье больше потребности ребёнка",
			g:        []int{3},
			s:        []int{5},
			expected: 1,
		},
		{
			name:     "повторяющиеся значения",
			g:        []int{2, 2, 2},
			s:        []int{2, 2},
			expected: 2,
		},
		{
			name:     "много печенья, мало детей",
			g:        []int{1, 2},
			s:        []int{1, 2, 3, 4, 5},
			expected: 2,
		},
		{
			name:     "много детей, мало печенья",
			g:        []int{1, 2, 3, 4, 5},
			s:        []int{1, 2},
			expected: 2,
		},
		{
			name:     "массивы в обратном порядке",
			g:        []int{5, 4, 3, 2, 1},
			s:        []int{5, 4, 3, 2, 1},
			expected: 5,
		},
		{
			name:     "ни одно печенье не подходит",
			g:        []int{5, 6, 7},
			s:        []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "все дети могут быть довольны",
			g:        []int{1, 2, 3, 4},
			s:        []int{2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "печенья подходят не всем детям",
			g:        []int{1, 3, 5, 7},
			s:        []int{2, 4, 6},
			expected: 3,
		},
		{
			name:     "большие значения",
			g:        []int{100, 200, 300},
			s:        []int{150, 250, 350},
			expected: 3,
		},
		{
			name:     "нулевые значения",
			g:        []int{0, 1, 2},
			s:        []int{0, 1},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tasks.FindContentChildren455(tt.g, tt.s)

			if result != tt.expected {
				t.Errorf(
					"g = %v, s = %v: got %d, want %d",
					tt.g,
					tt.s,
					result,
					tt.expected,
				)
			}
		})
	}
}
