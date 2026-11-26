package tests

import (
	"golang/tasks"
	"testing"
)

func TestIsIsomorphic205(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "basic isomorphic strings",
			s:        "egg",
			t:        "add",
			expected: true,
		},
		{
			name:     "not isomorphic - different mapping",
			s:        "foo",
			t:        "bar",
			expected: false,
		},
		{
			name:     "not isomorphic - same target character",
			s:        "ab",
			t:        "cc",
			expected: false,
		},
		{
			name:     "same strings",
			s:        "paper",
			t:        "title",
			expected: true,
		},
		{
			name:     "different lengths",
			s:        "abc",
			t:        "ab",
			expected: false,
		},
		{
			name:     "single character",
			s:        "a",
			t:        "b",
			expected: true,
		},
		{
			name:     "empty strings",
			s:        "",
			t:        "",
			expected: true,
		},
		{
			name:     "all same characters",
			s:        "aaaa",
			t:        "bbbb",
			expected: true,
		},
		{
			name:     "different repetition patterns",
			s:        "aabb",
			t:        "abab",
			expected: false,
		},
		{
			name:     "complex valid mapping",
			s:        "abcabc",
			t:        "defdef",
			expected: true,
		},
		{
			name:     "reverse mapping conflict",
			s:        "ab",
			t:        "ca",
			expected: true,
		},
		{
			name:     "longer invalid mapping",
			s:        "badc",
			t:        "baba",
			expected: false,
		},
		{
			name:     "unicode characters",
			s:        "你好你好",
			t:        "世界世界",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.IsIsomorphic205(tt.s, tt.t)

			if got != tt.expected {
				t.Errorf(
					"IsIsomorphic205(%q, %q) = %v; want %v",
					tt.s,
					tt.t,
					got,
					tt.expected,
				)
			}
		})
	}
}
