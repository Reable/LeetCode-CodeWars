package tests

import (
	"golang/tasks"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		name       string
		ransomNote string
		magazine   string
		want       bool
	}{
		{
			name:       "basic success",
			ransomNote: "a",
			magazine:   "a",
			want:       true,
		},
		{
			name:       "missing character",
			ransomNote: "aa",
			magazine:   "ab",
			want:       false,
		},
		{
			name:       "not enough characters",
			ransomNote: "ce",
			magazine:   "e",
			want:       false,
		},
		{
			name:       "enough repeated characters",
			ransomNote: "aa",
			magazine:   "aab",
			want:       true,
		},
		{
			name:       "different character order",
			ransomNote: "abc",
			magazine:   "cba",
			want:       true,
		},
		{
			name:       "empty ransom note",
			ransomNote: "",
			magazine:   "abc",
			want:       true,
		},
		{
			name:       "both strings empty",
			ransomNote: "",
			magazine:   "",
			want:       true,
		},
		{
			name:       "empty magazine",
			ransomNote: "a",
			magazine:   "",
			want:       false,
		},
		{
			name:       "multiple repeated characters",
			ransomNote: "aaabb",
			magazine:   "aabbba",
			want:       true,
		},
		{
			name:       "one extra required character",
			ransomNote: "aaa",
			magazine:   "aa",
			want:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.CanConstruct383(tt.ransomNote, tt.magazine)

			if got != tt.want {
				t.Errorf(
					"canConstruct(%q, %q) = %v; want %v",
					tt.ransomNote,
					tt.magazine,
					got,
					tt.want,
				)
			}
		})
	}
}
