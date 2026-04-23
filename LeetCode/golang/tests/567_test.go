package tests

import (
	"golang/tasks"
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "пример true",
			s1:       "ab",
			s2:       "eidbaooo",
			expected: true,
		},
		{
			name:     "пример false",
			s1:       "ab",
			s2:       "eidboaoo",
			expected: false,
		},
		{
			name:     "s1 длиннее s2",
			s1:       "abcd",
			s2:       "abc",
			expected: false,
		},
		{
			name:     "s1 пустая",
			s1:       "",
			s2:       "anything",
			expected: true,
		},
		{
			name:     "оба пустые",
			s1:       "",
			s2:       "",
			expected: true,
		},
		{
			name:     "один символ true",
			s1:       "a",
			s2:       "a",
			expected: true,
		},
		{
			name:     "один символ false",
			s1:       "a",
			s2:       "b",
			expected: false,
		},
		{
			name:     "повторяющиеся символы",
			s1:       "aab",
			s2:       "eidbaaboo",
			expected: true,
		},
		{
			name:     "нет подходящей перестановки",
			s1:       "abc",
			s2:       "ccccbbbbaaaa",
			expected: false,
		},
		{
			name:     "перестановка в начале",
			s1:       "abc",
			s2:       "cbaaaa",
			expected: true,
		},
		{
			name:     "перестановка в конце",
			s1:       "abc",
			s2:       "aaaacba",
			expected: true,
		},
		{
			name:     "всё строка совпадает",
			s1:       "abc",
			s2:       "abc",
			expected: true,
		},
		{
			name:     "перекрывающиеся окна",
			s1:       "ab",
			s2:       "abab",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tasks.CheckInclusion567(tt.s1, tt.s2)

			if result != tt.expected {
				t.Errorf("s1=%q s2=%q: got %v, want %v",
					tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}
