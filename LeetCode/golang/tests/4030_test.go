package tests

import (
	"golang/tasks"
	"testing"
)

func TestIsPalindromic4030(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Empty string",
			s:    "",
			want: true,
		},
		{
			name: "Single character",
			s:    "a",
			want: false,
		},
		{
			name: "Two different characters",
			s:    "ab",
			want: false,
		},
		{
			name: "Palindrome",
			s:    "aba",
			want: false,
		},
		{
			name: "Even length palindrome",
			s:    "abba",
			want: false,
		},
		{
			name: "Non-palindrome",
			s:    "abcd",
			want: false,
		},
		{
			name: "Repeated binary characters",
			s:    "ff",
			want: true,
		},
		{
			name: "Different ASCII characters",
			s:    "ab",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.IsPalindromic4030(tt.s)

			if got != tt.want {
				t.Errorf(
					"IsPalindromic4030(%q) = %v, want %v",
					tt.s,
					got,
					tt.want,
				)
			}
		})
	}
}
