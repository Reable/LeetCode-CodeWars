package tests

import (
	"golang/tasks"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "basic example",
			s:    "abccccdd",
			want: 7,
		},
		{
			name: "all same characters",
			s:    "ccc",
			want: 3,
		},
		{
			name: "all unique characters",
			s:    "abcdef",
			want: 1,
		},
		{
			name: "all even frequencies",
			s:    "aabbcc",
			want: 6,
		},
		{
			name: "multiple odd frequencies",
			s:    "aabbbcdd",
			want: 7,
		},
		{
			name: "single character",
			s:    "a",
			want: 1,
		},
		{
			name: "empty string",
			s:    "",
			want: 0,
		},
		{
			name: "two different characters",
			s:    "ab",
			want: 1,
		},
		{
			name: "one odd and multiple even",
			s:    "aabbc",
			want: 5,
		},
		{
			name: "many repeated characters",
			s:    "aaaabbbbcccc",
			want: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.LongestPalindrome409(tt.s)

			if got != tt.want {
				t.Errorf(
					"longestPalindrome(%q) = %d, want %d",
					tt.s,
					got,
					tt.want,
				)
			}
		})
	}
}
