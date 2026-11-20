package tests

import (
	"golang/tasks"
	"testing"
)

func TestMaximumLengthSubstring3090_v2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "Example 1",
			input: "bcbbbcba",
			want:  4,
		},
		{
			name:  "All same characters",
			input: "aaaaa",
			want:  2,
		},
		{
			name:  "All unique characters",
			input: "abcdef",
			want:  6,
		},
		{
			name:  "Exactly two occurrences",
			input: "aabbcc",
			want:  6,
		},
		{
			name:  "Three occurrences",
			input: "aaab",
			want:  3,
		},
		{
			name:  "Multiple repeated characters",
			input: "aabccba",
			want:  6,
		},
		{
			name:  "Single character",
			input: "a",
			want:  1,
		},
		{
			name:  "Two characters",
			input: "ab",
			want:  2,
		},
		{
			name:  "Empty string",
			input: "",
			want:  0,
		},
		{
			name:  "Longest substring in the middle",
			input: "aaaabcdef",
			want:  7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.MaximumLengthSubstring3090_v2(tt.input)

			if got != tt.want {
				t.Errorf(
					"MaximumLengthSubstring3090_v2(%q) = %d; want %d",
					tt.input,
					got,
					tt.want,
				)
			}
		})
	}
}
