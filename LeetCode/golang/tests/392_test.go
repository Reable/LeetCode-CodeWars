package tests

import (
	"golang/tasks"
	"testing"
)

func TestIsSubsequence392(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "базовый пример 1 (подпоследовательность найдена)",
			s:        "abc",
			t:        "ahbgdc",
			expected: true,
		},
		{
			name:     "базовый пример 2 (символ не найден)",
			s:        "axc",
			t:        "ahbgdc",
			expected: false,
		},
		{
			name:     "пустая s (пустая строка является подпоследовательностью любой)",
			s:        "",
			t:        "ahbgdc",
			expected: true,
		},
		{
			name:     "обе строки пустые",
			s:        "",
			t:        "",
			expected: true,
		},
		{
			name:     "непустая s и пустая t",
			s:        "a",
			t:        "",
			expected: false,
		},
		{
			name:     "полное совпадение строк",
			s:        "abc",
			t:        "abc",
			expected: true,
		},
		{
			name:     "s длиннее, чем t",
			s:        "abcdef",
			t:        "abc",
			expected: false,
		},
		{
			name:     "символы те же, но в неправильном порядке",
			s:        "acb",
			t:        "ahbgdc",
			expected: false,
		},
		{
			name:     "повторяющиеся символы (хватает)",
			s:        "aaaa",
			t:        "bbaaaaacc",
			expected: true,
		},
		{
			name:     "повторяющиеся символы (не хватает одного)",
			s:        "aaaa",
			t:        "bbaaacc", // здесь только три буквы 'a'
			expected: false,
		},
		{
			name:     "одиночный символ в самом конце t",
			s:        "z",
			t:        "abcdefgz",
			expected: true,
		},
		{
			name:     "одиночный символ, которого нет в t",
			s:        "x",
			t:        "abcdefg",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tasks.IsSubsequence392(tt.s, tt.t)

			if result != tt.expected {
				t.Errorf("s = %q, t = %q: got %v, want %v",
					tt.s, tt.t, result, tt.expected)
			}
		})
	}
}
