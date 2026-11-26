package tests

import (
	"golang/tasks"
	"testing"
)

func TestWordPattern290(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		s       string
		want    bool
	}{
		{
			name:    "basic match",
			pattern: "abba",
			s:       "dog cat cat dog",
			want:    true,
		},
		{
			name:    "pattern mismatch",
			pattern: "abba",
			s:       "dog cat cat fish",
			want:    false,
		},
		{
			name:    "same word mapped to different characters",
			pattern: "abba",
			s:       "dog dog dog dog",
			want:    false,
		},
		{
			name:    "different words mapped to same character",
			pattern: "aaaa",
			s:       "dog cat cat dog",
			want:    false,
		},
		{
			name:    "different lengths",
			pattern: "abc",
			s:       "dog cat",
			want:    false,
		},
		{
			name:    "single character",
			pattern: "a",
			s:       "dog",
			want:    true,
		},
		{
			name:    "single character mismatch",
			pattern: "a",
			s:       "dog cat",
			want:    false,
		},
		{
			name:    "all unique characters and words",
			pattern: "abc",
			s:       "dog cat fish",
			want:    true,
		},
		{
			name:    "repeated characters and words",
			pattern: "abab",
			s:       "dog cat dog cat",
			want:    true,
		},
		{
			name:    "empty pattern and string",
			pattern: "",
			s:       "",
			want:    true,
		},
		{
			name:    "empty pattern with non-empty string",
			pattern: "",
			s:       "dog",
			want:    false,
		},
		{
			name:    "same word at different positions",
			pattern: "abcba",
			s:       "one two three two one",
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.WordPattern290(tt.pattern, tt.s)

			if got != tt.want {
				t.Errorf(
					"WordPattern290(%q, %q) = %v, want %v",
					tt.pattern,
					tt.s,
					got,
					tt.want,
				)
			}
		})
	}
}
