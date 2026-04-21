package tests

import (
	"golang/tasks"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "пример 1",
			input:    "abcabcbb",
			expected: 3,
		},
		{
			name:     "все одинаковые",
			input:    "bbbbb",
			expected: 1,
		},
		{
			name:     "пример 2",
			input:    "pwwkew",
			expected: 3,
		},
		{
			name:     "пустая строка",
			input:    "",
			expected: 0,
		},
		{
			name:     "один символ",
			input:    "a",
			expected: 1,
		},
		{
			name:     "все уникальные",
			input:    "abcdef",
			expected: 6,
		},
		{
			name:     "повтор в середине",
			input:    "abccdef",
			expected: 4, // "cdef"
		},
		{
			name:     "повтор в начале",
			input:    "aab",
			expected: 2,
		},
		{
			name:     "повтор с пересечением",
			input:    "abba",
			expected: 2,
		},
		{
			name:     "сложный случай",
			input:    "dvdf",
			expected: 3, // "vdf"
		},
		{
			name:     "с пробелами",
			input:    "a b c a b",
			expected: 3,
		},
		{
			name:     "цифры и символы",
			input:    "123@123!",
			expected: 5, // "@123!"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tasks.LengthOfLongestSubstring(tt.input)

			if result != tt.expected {
				t.Errorf("input %q: got %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
