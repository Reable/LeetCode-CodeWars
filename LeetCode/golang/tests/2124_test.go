package tests

import (
	"golang/tasks"
	"testing"
)

func TestCheckString2124(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "empty string",
			s:    "",
			want: true,
		},
		{
			name: "only a",
			s:    "aaaa",
			want: true,
		},
		{
			name: "only b",
			s:    "bbbb",
			want: true,
		},
		{
			name: "a before b",
			s:    "aaabbb",
			want: true,
		},
		{
			name: "single a and b",
			s:    "ab",
			want: true,
		},
		{
			name: "b before a",
			s:    "ba",
			want: false,
		},
		{
			name: "a after b",
			s:    "aababb",
			want: false,
		},
		{
			name: "b in the middle",
			s:    "aababb",
			want: false,
		},
		{
			name: "b followed by multiple a",
			s:    "bbaaa",
			want: false,
		},
		{
			name: "a followed by b",
			s:    "abbbb",
			want: true,
		},
		{
			name: "multiple transitions",
			s:    "aaabbba",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.CheckString2124(tt.s)

			if got != tt.want {
				t.Errorf(
					"CheckString2124(%q) = %v, want %v",
					tt.s,
					got,
					tt.want,
				)
			}
		})
	}
}
