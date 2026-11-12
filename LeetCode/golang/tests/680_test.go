package tests

import (
	"golang/tasks"
	"testing"
)

func TestValidPalindrome680(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "уже палиндром",
			input: "aba",
			want:  true,
		},
		{
			name:  "удаление символа из середины",
			input: "abca",
			want:  true,
		},
		{
			name:  "удаление левого символа при несовпадении",
			input: "ebcbbe",
			want:  true,
		},
		{
			name:  "удаление правого символа при несовпадении",
			input: "deeee",
			want:  true,
		},
		{
			name:  "невозможно получить палиндром",
			input: "abc",
			want:  false,
		},
		{
			name:  "пустая строка",
			input: "",
			want:  true,
		},
		{
			name:  "один символ",
			input: "a",
			want:  true,
		},
		{
			name:  "два одинаковых символа",
			input: "aa",
			want:  true,
		},
		{
			name:  "два разных символа",
			input: "ab",
			want:  true,
		},
		{
			name:  "длинный палиндром",
			input: "racecar",
			want:  true,
		},
		{
			name:  "длинная строка с одним лишним символом",
			input: "aguokepatgbnvfqmgmlcuputjc",
			want:  false,
		},
		{
			name:  "нужно удалить больше одного символа",
			input: "abcdef",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.ValidPalindrome680(tt.input)

			if got != tt.want {
				t.Errorf(
					"ValidPalindrome680(%q) = %v, want %v",
					tt.input,
					got,
					tt.want,
				)
			}
		})
	}
}
