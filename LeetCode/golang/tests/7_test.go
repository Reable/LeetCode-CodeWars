package tests

import (
	"golang/tasks"
	"testing"
)

func TestReverse7(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "положительное число",
			input:    123,
			expected: 321,
		},
		{
			name:     "отрицательное число",
			input:    -123,
			expected: -321,
		},
		{
			name:     "число с нулём на конце",
			input:    120,
			expected: 21,
		},
		{
			name:     "ноль",
			input:    0,
			expected: 0,
		},
		{
			name:     "однозначное",
			input:    7,
			expected: 7,
		},
		{
			name:     "переполнение положительное",
			input:    1534236469,
			expected: 0, // выходит за int32
		},
		{
			name:     "переполнение отрицательное",
			input:    -1534236469,
			expected: 0,
		},
		{
			name:     "граница int32 max",
			input:    2147483647,
			expected: 0,
		},
		{
			name:     "граница int32 min",
			input:    -2147483648,
			expected: 0,
		},
		{
			name:     "палиндром",
			input:    1221,
			expected: 1221,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tasks.Reverse(tt.input)

			if result != tt.expected {
				t.Errorf("input=%d: got %d, want %d",
					tt.input, result, tt.expected)
			}
		})
	}
}
