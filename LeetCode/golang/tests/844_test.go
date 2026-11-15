package tests

import (
	"golang/tasks"
	"testing"
)

func TestBackspaceCompare844(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "basic example",
			s:    "ab#c",
			t:    "ad#c",
			want: true,
		},
		{
			name: "different results",
			s:    "a#c",
			t:    "b",
			want: false,
		},
		{
			name: "multiple backspaces",
			s:    "ab##",
			t:    "c#d#",
			want: true,
		},
		{
			name: "backspaces at beginning",
			s:    "###abc",
			t:    "abc",
			want: true,
		},
		{
			name: "all characters deleted",
			s:    "abc###",
			t:    "",
			want: true,
		},
		{
			name: "empty strings",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "one empty string",
			s:    "a",
			t:    "",
			want: false,
		},
		{
			name: "different lengths",
			s:    "abc",
			t:    "ab",
			want: false,
		},
		{
			name: "consecutive backspaces",
			s:    "a##c",
			t:    "c",
			want: true,
		},
		{
			name: "backspace without character",
			s:    "###a",
			t:    "a",
			want: true,
		},
		{
			name: "same strings without backspaces",
			s:    "hello",
			t:    "hello",
			want: true,
		},
		{
			name: "different strings without backspaces",
			s:    "hello",
			t:    "world",
			want: false,
		},
		{
			name: "backspace removes different characters",
			s:    "xywrrmp",
			t:    "xywrrmu#p",
			want: true,
		},
		{
			name: "all backspaces",
			s:    "####",
			t:    "###",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV1 := tasks.BackspaceCompare844_v1(tt.s, tt.t)
			gotV2 := tasks.BackspaceCompare844_v2(tt.s, tt.t)

			if gotV1 != tt.want {
				t.Errorf("v1: got %v, want %v", gotV1, tt.want)
			}

			if gotV2 != tt.want {
				t.Errorf("v2: got %v, want %v", gotV2, tt.want)
			}

			if gotV1 != gotV2 {
				t.Errorf("v1 and v2 differ: v1=%v, v2=%v", gotV1, gotV2)
			}
		})
	}
}
