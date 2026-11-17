package tests

import (
	"golang/tasks"
	"testing"
)

func TestIsPrefixString1961(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		words []string
		want  bool
	}{
		{
			name:  "LeetCode example 1",
			s:     "iloveleetcode",
			words: []string{"i", "love", "leetcode", "apples"},
			want:  true,
		},
		{
			name:  "LeetCode example 2",
			s:     "iloveleetcode",
			words: []string{"apples", "i", "love", "leetcode"},
			want:  false,
		},
		{
			name:  "exact match with one word",
			s:     "hello",
			words: []string{"hello"},
			want:  true,
		},
		{
			name:  "exact match with multiple words",
			s:     "helloworld",
			words: []string{"hello", "world"},
			want:  true,
		},
		{
			name:  "string is longer than concatenation",
			s:     "helloworld!",
			words: []string{"hello", "world"},
			want:  false,
		},
		{
			name:  "string is shorter than first word",
			s:     "hi",
			words: []string{"hello", "world"},
			want:  false,
		},
		{
			name:  "first word does not match",
			s:     "hello",
			words: []string{"world", "hello"},
			want:  false,
		},
		{
			name:  "partial match",
			s:     "helloworld",
			words: []string{"hello", "wor"},
			want:  false,
		},
		{
			name:  "extra words after exact match",
			s:     "hello",
			words: []string{"hello", "world"},
			want:  true,
		},
		{
			name:  "empty string",
			s:     "",
			words: []string{"hello"},
			want:  false,
		},
		{
			name:  "empty words",
			s:     "hello",
			words: []string{},
			want:  false,
		},
		{
			name:  "both empty",
			s:     "",
			words: []string{},
			want:  false,
		},
		{
			name:  "single character",
			s:     "a",
			words: []string{"a", "b"},
			want:  true,
		},
		{
			name:  "empty word in the beginning",
			s:     "hello",
			words: []string{"", "hello"},
			want:  true,
		},
		{
			name:  "empty word between words",
			s:     "helloworld",
			words: []string{"hello", "", "world"},
			want:  true,
		},
		{
			name:  "empty word before incomplete match",
			s:     "hello",
			words: []string{"", "hel"},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.IsPrefixString1961(tt.s, tt.words)

			if got != tt.want {
				t.Errorf(
					"IsPrefixString1961(%q, %v) = %v, want %v",
					tt.s,
					tt.words,
					got,
					tt.want,
				)
			}
		})
	}
}
