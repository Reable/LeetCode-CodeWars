package tests

import (
	"golang/tasks"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string // допускаем несколько правильных ответов
	}{
		{
			name:     "пример 1",
			input:    "babad",
			expected: []string{"bab", "aba"},
		},
		{
			name:     "пример 2",
			input:    "cbbd",
			expected: []string{"bb"},
		},
		{
			name:     "один символ",
			input:    "a",
			expected: []string{"a"},
		},
		{
			name:     "пустая строка",
			input:    "",
			expected: []string{""},
		},
		{
			name:     "все одинаковые",
			input:    "aaaa",
			expected: []string{"aaaa"},
		},
		{
			name:     "нет длинных палиндромов",
			input:    "abc",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "палиндром в центре",
			input:    "abccba",
			expected: []string{"abccba"},
		},
		{
			name:     "чётная длина",
			input:    "abba",
			expected: []string{"abba"},
		},
		{
			name:     "нечётная длина",
			input:    "racecar",
			expected: []string{"racecar"},
		},
		{
			name:     "сложный случай",
			input:    "forgeeksskeegfor",
			expected: []string{"geeksskeeg"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tasks.LongestPalindrome(tt.input)

			if !contains(tt.expected, result) {
				t.Errorf("input %q: got %q, want one of %v",
					tt.input, result, tt.expected)
			}
		})
	}
}

func contains(arr []string, target string) bool {
	for _, v := range arr {
		if v == target {
			return true
		}
	}
	return false
}
