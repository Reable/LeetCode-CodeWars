package tests

import (
	"golang/tasks"
	"testing"
)

func TestReverseWords557(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "обычная строка",
			input: "Let's take LeetCode contest",
			want:  "s'teL ekat edoCteeL tsetnoc",
		},
		{
			name:  "одно слово",
			input: "hello",
			want:  "olleh",
		},
		{
			name:  "пустая строка",
			input: "",
			want:  "",
		},
		{
			name:  "слово из одного символа",
			input: "a",
			want:  "a",
		},
		{
			name:  "слова разной длины",
			input: "Go is awesome",
			want:  "oG si emosewa",
		},
		{
			name:  "несколько пробелов между словами",
			input: "hello  world",
			want:  "olleh  dlrow",
		},
		{
			name:  "слово из двух символов",
			input: "ab cd",
			want:  "ba dc",
		},
		{
			name:  "спецсимволы",
			input: "123 !@#",
			want:  "321 #@!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.ReverseWords557(tt.input)

			if got != tt.want {
				t.Errorf(
					"ReverseWords557(%q) = %q, want %q",
					tt.input,
					got,
					tt.want,
				)
			}
		})
	}
}
