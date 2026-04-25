package tests

import (
	"golang/tasks"
	"testing"
)

func TestIsPalindrome9(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "палиндром",
			input:    121,
			expected: true,
		},
		{
			name:     "не палиндром",
			input:    123,
			expected: false,
		},
		{
			name:     "отрицательное число",
			input:    -121,
			expected: false,
		},
		{
			name:     "ноль",
			input:    0,
			expected: true,
		},
		{
			name:     "число с нулём на конце",
			input:    10,
			expected: false,
		},
		{
			name:     "одна цифра",
			input:    7,
			expected: true,
		},
		{
			name:     "чётная длина",
			input:    1221,
			expected: true,
		},
		{
			name:     "нечётная длина",
			input:    12321,
			expected: true,
		},
		{
			name:     "почти палиндром",
			input:    12331,
			expected: false,
		},
		{
			name:     "большое палиндромное число",
			input:    2147447412,
			expected: true,
		},
		{
			name:     "большое не палиндромное число",
			input:    2147483647,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tasks.IsPalindrome9(tt.input)

			if result != tt.expected {
				t.Errorf("input=%d: got %v, want %v",
					tt.input, result, tt.expected)
			}
		})
	}
}
