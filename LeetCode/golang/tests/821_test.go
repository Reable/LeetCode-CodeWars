package tests

import (
	"golang/tasks"
	"reflect"
	"testing"
)

func TestShortestToChar821(t *testing.T) {
	tests := []struct {
		name string
		s    string
		c    byte
		want []int
	}{
		{
			name: "basic example",
			s:    "loveleetcode",
			c:    'e',
			want: []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
		{
			name: "character at beginning",
			s:    "abca",
			c:    'a',
			want: []int{0, 1, 1, 0},
		},
		{
			name: "character at end",
			s:    "abcd",
			c:    'd',
			want: []int{3, 2, 1, 0},
		},
		{
			name: "character in middle",
			s:    "abcde",
			c:    'c',
			want: []int{2, 1, 0, 1, 2},
		},
		{
			name: "multiple occurrences",
			s:    "aabaa",
			c:    'a',
			want: []int{0, 0, 1, 0, 0},
		},
		{
			name: "single character",
			s:    "a",
			c:    'a',
			want: []int{0},
		},
		{
			name: "empty string",
			s:    "",
			c:    'a',
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.ShortestToChar821(tt.s, tt.c)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShortestToChar821(%q, %q) = %v, want %v",
					tt.s, tt.c, got, tt.want)
			}
		})
	}
}
