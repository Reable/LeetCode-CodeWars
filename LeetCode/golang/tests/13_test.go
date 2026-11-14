package tests

import (
	"golang/tasks"
	"testing"
)

func TestRomanToInt13(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "single symbol",
			input: "I",
			want:  1,
		},
		{
			name:  "repeated symbols",
			input: "III",
			want:  3,
		},
		{
			name:  "subtraction IV",
			input: "IV",
			want:  4,
		},
		{
			name:  "subtraction IX",
			input: "IX",
			want:  9,
		},
		{
			name:  "addition VI",
			input: "VI",
			want:  6,
		},
		{
			name:  "mixed symbols LVIII",
			input: "LVIII",
			want:  58,
		},
		{
			name:  "complex number MCMXCIV",
			input: "MCMXCIV",
			want:  1994,
		},
		{
			name:  "largest common example",
			input: "MMMCMXCIX",
			want:  3999,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.RomanToInt13(tt.input)

			if got != tt.want {
				t.Errorf(
					"RomanToInt13(%q) = %d; want %d",
					tt.input,
					got,
					tt.want,
				)
			}
		})
	}
}
