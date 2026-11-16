package tests

import (
	"golang/tasks"
	"testing"
)

func TestReversePrefix2000(t *testing.T) {
	tests := []struct {
		name string
		word string
		ch   byte
		want string
	}{
		{
			name: "First example",
			word: "abcdefd",
			ch:   'd',
			want: "dcbaefd",
		},
		{
			name: "Second example",
			word: "xyxzxe",
			ch:   'z',
			want: "zxyxxe",
		},
		{
			name: "Character not found",
			word: "abcd",
			ch:   'z',
			want: "abcd",
		},
		{
			name: "Character at first position",
			word: "hello",
			ch:   'h',
			want: "hello",
		},
		{
			name: "Character at last position",
			word: "abcde",
			ch:   'e',
			want: "edcba",
		},
		{
			name: "Character at second position",
			word: "abcde",
			ch:   'b',
			want: "bacde",
		},
		{
			name: "Repeated character",
			word: "abacada",
			ch:   'a',
			want: "abacada",
		},
		{
			name: "Empty string",
			word: "",
			ch:   'a',
			want: "",
		},
		{
			name: "Single character",
			word: "a",
			ch:   'a',
			want: "a",
		},
		{
			name: "Single character not found",
			word: "a",
			ch:   'b',
			want: "a",
		},
		{
			name: "Reverse entire word",
			word: "racecar",
			ch:   'r',
			want: "racecar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.ReversePrefix2000(tt.word, tt.ch)

			if got != tt.want {
				t.Errorf(
					"ReversePrefix2000(%q, %q) = %q, want %q",
					tt.word,
					tt.ch,
					got,
					tt.want,
				)
			}
		})
	}
}
