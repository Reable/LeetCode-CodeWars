package tests

import (
	"golang/tasks"
	"testing"
)

func TestFirstPalindrome2108(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  string
	}{
		{
			name:  "First example",
			words: []string{"abc", "car", "ada", "racecar", "cool"},
			want:  "ada",
		},
		{
			name:  "Second example",
			words: []string{"notapalindrome", "racecar"},
			want:  "racecar",
		},
		{
			name:  "No palindrome",
			words: []string{"def", "ghi"},
			want:  "",
		},
		{
			name:  "Single character",
			words: []string{"a", "abc"},
			want:  "a",
		},
		{
			name:  "Two equal characters",
			words: []string{"aa", "abc"},
			want:  "aa",
		},
		{
			name:  "Two different characters",
			words: []string{"ab", "cd"},
			want:  "",
		},
		{
			name:  "Odd length palindrome",
			words: []string{"abcba"},
			want:  "abcba",
		},
		{
			name:  "Even length palindrome",
			words: []string{"abccba"},
			want:  "abccba",
		},
		{
			name:  "Empty array",
			words: []string{},
			want:  "",
		},
		{
			name:  "Empty string",
			words: []string{""},
			want:  "",
		},
		{
			name:  "First palindrome wins",
			words: []string{"racecar", "ada", "level"},
			want:  "racecar",
		},
		{
			name:  "Palindrome at the end",
			words: []string{"abc", "def", "level"},
			want:  "level",
		},
		{
			name:  "Palindrome with repeated characters",
			words: []string{"aaaa", "abc"},
			want:  "aaaa",
		},
		{
			name:  "Palindrome with one middle character",
			words: []string{"abacaba"},
			want:  "abacaba",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"/v1", func(t *testing.T) {
			got := tasks.FirstPalindrome2108_v1(tt.words)

			if got != tt.want {
				t.Errorf("v1: got %q, want %q", got, tt.want)
			}
		})

		t.Run(tt.name+"/v2", func(t *testing.T) {
			got := tasks.FirstPalindrome2108_v2(tt.words)

			if got != tt.want {
				t.Errorf("v2: got %q, want %q", got, tt.want)
			}
		})
	}
}
