package tests

import (
	"golang/tasks"
	"testing"
)

func TestCountKConstraintSubstrings3258(t *testing.T) {
	tests := []struct {
		name string
		s    string
		k    int
		want int
	}{
		{
			name: "Example 1",
			s:    "10101",
			k:    1,
			want: 12,
		},
		{
			name: "Example 2",
			s:    "1010101",
			k:    2,
			want: 25,
		},
		{
			name: "All zeros",
			s:    "00000",
			k:    1,
			want: 15,
		},
		{
			name: "All ones",
			s:    "11111",
			k:    1,
			want: 15,
		},
		{
			name: "Single zero",
			s:    "0",
			k:    0,
			want: 1,
		},
		{
			name: "Single one",
			s:    "1",
			k:    0,
			want: 1,
		},
		{
			name: "Alternating with k zero",
			s:    "01010",
			k:    0,
			want: 5,
		},
		{
			name: "All substrings valid",
			s:    "10101",
			k:    5,
			want: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.CountKConstraintSubstrings3258(tt.s, tt.k)

			if got != tt.want {
				t.Errorf(
					"CountKConstraintSubstrings3258(%q, %d) = %d; want %d",
					tt.s, tt.k, got, tt.want,
				)
			}
		})
	}
}
