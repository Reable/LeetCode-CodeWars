package tests

import (
	"golang/tasks"
	"testing"
)

func TestFindTheDifference389_v1(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want byte
	}{
		{
			name: "basic example",
			s:    "abcd",
			t:    "abcde",
			want: 'e',
		},
		{
			name: "different order",
			s:    "aabbcc",
			t:    "abcbacd",
			want: 'd',
		},
		{
			name: "single character",
			s:    "",
			t:    "a",
			want: 'a',
		},
		{
			name: "repeated characters",
			s:    "aaaa",
			t:    "aaaaa",
			want: 'a',
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.FindTheDifference389_v1(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("FindTheDifference389_v1(%q, %q) = %q, want %q",
					tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func TestFindTheDifference389_v2(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{
			name: "basic example",
			s:    "abcd",
			t:    "abcde",
			want: "e",
		},
		{
			name: "different order",
			s:    "aabbcc",
			t:    "abcbacd",
			want: "d",
		},
		{
			name: "single character",
			s:    "",
			t:    "a",
			want: "a",
		},
		{
			name: "repeated characters",
			s:    "aaaa",
			t:    "aaaaa",
			want: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.FindTheDifference389_v2(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("FindTheDifference389_v2(%q, %q) = %q, want %q",
					tt.s, tt.t, got, tt.want)
			}
		})
	}
}
