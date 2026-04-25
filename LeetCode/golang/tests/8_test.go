package tests

import (
	"golang/tasks"
	"testing"
)

func TestMyAtoi8(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "простое число",
			input:    "42",
			expected: 42,
		},
		{
			name:     "с пробелами",
			input:    "   42",
			expected: 42,
		},
		{
			name:     "отрицательное число",
			input:    "   -42",
			expected: -42,
		},
		{
			name:     "плюс",
			input:    "+123",
			expected: 123,
		},
		{
			name:     "с текстом после",
			input:    "4193 with words",
			expected: 4193,
		},
		{
			name:     "с текстом перед",
			input:    "words and 987",
			expected: 0,
		},
		{
			name:     "смешанный знак",
			input:    "+-12",
			expected: 0,
		},
		{
			name:     "только знак",
			input:    "-",
			expected: 0,
		},
		{
			name:     "пустая строка",
			input:    "",
			expected: 0,
		},
		{
			name:     "нули впереди",
			input:    "00000123",
			expected: 123,
		},
		{
			name:     "нули + отрицательное",
			input:    "-00000123",
			expected: -123,
		},
		{
			name:     "переполнение вверх",
			input:    "91283472332",
			expected: 2147483647, // MaxInt32
		},
		{
			name:     "переполнение вниз",
			input:    "-91283472332",
			expected: -2147483648, // MinInt32
		},
		{
			name:     "граница max",
			input:    "2147483647",
			expected: 2147483647,
		},
		{
			name:     "граница min",
			input:    "-2147483648",
			expected: -2147483648,
		},
		{
			name:     "чуть больше max",
			input:    "2147483648",
			expected: 2147483647,
		},
		{
			name:     "чуть меньше min",
			input:    "-2147483649",
			expected: -2147483648,
		},
		{
			name:     "число с пробелом внутри",
			input:    "123 456",
			expected: 123,
		},
		{
			name:     "буква сразу после знака",
			input:    "-a123",
			expected: 0,
		},
		{
			name:     "пробел после знака",
			input:    "- 123",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tasks.MyAtoi8(tt.input)

			if result != tt.expected {
				t.Errorf("input=%q: got %d, want %d",
					tt.input, result, tt.expected)
			}
		})
	}
}
