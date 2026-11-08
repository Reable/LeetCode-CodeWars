package tests

import (
	"golang/tasks"
	"testing"
)

func TestCompress443(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
		length   int
	}{
		{
			name:     "Example 1",
			input:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected: []byte{'a', '2', 'b', '2', 'c', '3'},
			length:   6,
		},
		{
			name:     "Single character",
			input:    []byte{'a'},
			expected: []byte{'a'},
			length:   1,
		},
		{
			name:     "All unique characters",
			input:    []byte{'a', 'b', 'c'},
			expected: []byte{'a', 'b', 'c'},
			length:   3,
		},
		{
			name:     "One group",
			input:    []byte{'a', 'a', 'a', 'a'},
			expected: []byte{'a', '4'},
			length:   2,
		},
		{
			name:     "Count greater than 9",
			input:    []byte{'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a', 'a'},
			expected: []byte{'a', '1', '2'},
			length:   3,
		},
		{
			name:     "Multiple groups",
			input:    []byte{'a', 'a', 'b', 'c', 'c', 'c', 'c'},
			expected: []byte{'a', '2', 'b', 'c', '4'},
			length:   5,
		},
		{
			name:     "Empty slice",
			input:    []byte{},
			expected: []byte{},
			length:   0,
		},
		{
			name:     "Count 10",
			input:    []byte("aaaaaaaaaa"),
			expected: []byte{'a', '1', '0'},
			length:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]byte, len(tt.input))
			copy(input, tt.input)

			gotLength := tasks.Compress443(input)

			if gotLength != tt.length {
				t.Errorf("expected length %d, got %d", tt.length, gotLength)
			}

			got := input[:gotLength]

			if string(got) != string(tt.expected) {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
