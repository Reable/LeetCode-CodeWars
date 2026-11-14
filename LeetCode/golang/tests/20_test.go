package tests

import (
	"golang/tasks"
	"testing"
)

func TestIsValid20(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "empty string",
			input: "",
			want:  true,
		},
		{
			name:  "single pair",
			input: "()",
			want:  true,
		},
		{
			name:  "multiple pairs",
			input: "()[]{}",
			want:  true,
		},
		{
			name:  "nested brackets",
			input: "{[]}",
			want:  true,
		},
		{
			name:  "deeply nested brackets",
			input: "{[()]}",
			want:  true,
		},
		{
			name:  "wrong bracket type",
			input: "(]",
			want:  false,
		},
		{
			name:  "incorrect order",
			input: "([)]",
			want:  false,
		},
		{
			name:  "unclosed opening bracket",
			input: "(((",
			want:  false,
		},
		{
			name:  "closing bracket first",
			input: ")",
			want:  false,
		},
		{
			name:  "extra closing bracket",
			input: "()[]{}]",
			want:  false,
		},
		{
			name:  "unclosed nested bracket",
			input: "{[(",
			want:  false,
		},
		{
			name:  "mixed incorrect sequence",
			input: "({[)]}",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.IsValid20(tt.input)

			if got != tt.want {
				t.Errorf(
					"IsValid20(%q) = %v, want %v",
					tt.input,
					got,
					tt.want,
				)
			}
		})
	}
}
