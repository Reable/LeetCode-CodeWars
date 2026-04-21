package tests

import (
	"golang/tasks"
	"testing"
)

func TestConvert1(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		numRows  int
		expected string
	}{
		{
			name:     "пример 1",
			s:        "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			name:     "пример 2",
			s:        "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},
		{
			name:     "numRows = 1",
			s:        "AB",
			numRows:  1,
			expected: "AB",
		},
		{
			name:     "numRows = 2",
			s:        "ABCDEF",
			numRows:  2,
			expected: "ACEBDF",
		},
		{
			name:     "один символ",
			s:        "A",
			numRows:  3,
			expected: "A",
		},
		{
			name:     "numRows больше длины строки",
			s:        "ABC",
			numRows:  5,
			expected: "ABC",
		},
		{
			name:     "ровная строка",
			s:        "HELLOWORLD",
			numRows:  3,
			expected: "HOLELWRDLO",
		},
		{
			name:     "пустая строка",
			s:        "",
			numRows:  3,
			expected: "",
		},
		{
			name:     "повторяющиеся символы",
			s:        "AAAAAA",
			numRows:  3,
			expected: "AAAAAA",
		},
		{
			name:     "сложный случай",
			s:        "ABCDEFGHIJKLMN",
			numRows:  4,
			expected: "AGMBFHLNCEIKDJ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tasks.Convert1(tt.s, tt.numRows)

			if result != tt.expected {
				t.Errorf("s=%q numRows=%d: got %q, want %q",
					tt.s, tt.numRows, result, tt.expected)
			}
		})
	}
}
