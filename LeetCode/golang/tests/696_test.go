package tests

import (
	"golang/tasks"
	"testing"
)

func TestCountBinarySubstrings696(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "00110011",
			want: 6,
		},
		{
			name: "Example 2",
			s:    "10101",
			want: 4,
		},
		{
			name: "Single zero",
			s:    "0",
			want: 0,
		},
		{
			name: "Single one",
			s:    "1",
			want: 0,
		},
		{
			name: "Two different characters",
			s:    "01",
			want: 1,
		},
		{
			name: "Two equal characters",
			s:    "00",
			want: 0,
		},
		{
			name: "All zeros",
			s:    "00000",
			want: 0,
		},
		{
			name: "All ones",
			s:    "11111",
			want: 0,
		},
		{
			name: "Unequal groups",
			s:    "00111",
			want: 2,
		},
		{
			name: "Multiple groups",
			s:    "000111000",
			want: 6,
		},
		{
			name: "Alternating characters",
			s:    "01010101",
			want: 7,
		},
		{
			name: "Long first group",
			s:    "000011",
			want: 2,
		},
		{
			name: "Long second group",
			s:    "001111",
			want: 2,
		},
		{
			name: "Single valid substring",
			s:    "01111",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.CountBinarySubstrings696(tt.s)

			if got != tt.want {
				t.Errorf(
					"CountBinarySubstrings696(%q) = %d, want %d",
					tt.s,
					got,
					tt.want,
				)
			}
		})
	}
}
