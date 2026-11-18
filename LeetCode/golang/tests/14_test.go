package tests

import (
	"golang/tasks"
	"testing"
)

func TestLongestCommonPrefix14(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "Example 1",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "Example 2",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "Example 3",
			strs: []string{"interview", "internet", "internal"},
			want: "inter",
		},
		{
			name: "All strings equal",
			strs: []string{"hello", "hello", "hello"},
			want: "hello",
		},
		{
			name: "Single string",
			strs: []string{"hello"},
			want: "hello",
		},
		{
			name: "Empty slice",
			strs: []string{},
			want: "",
		},
		{
			name: "Empty string",
			strs: []string{"", "flower", "flow"},
			want: "",
		},
		{
			name: "One string is prefix of another",
			strs: []string{"abc", "abcd", "abcde"},
			want: "abc",
		},
		{
			name: "Different first characters",
			strs: []string{"apple", "banana", "cherry"},
			want: "",
		},
		{
			name: "Common prefix is one character",
			strs: []string{"car", "cat", "camera"},
			want: "ca",
		},
		{
			name: "Long common prefix",
			strs: []string{"prefix123", "prefix456", "prefix789"},
			want: "prefix",
		},
		{
			name: "Two strings",
			strs: []string{"testing", "test"},
			want: "test",
		},
		{
			name: "Repeated strings",
			strs: []string{"abc", "abc", "abd"},
			want: "ab",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.LongestCommonPrefix14(tt.strs)

			if got != tt.want {
				t.Errorf(
					"LongestCommonPrefix14(%v) = %q; want %q",
					tt.strs,
					got,
					tt.want,
				)
			}
		})
	}
}

func TestLongestCommonPrefix14V2(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "Example 1",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "Example 2",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "Example 3",
			strs: []string{"interview", "internet", "internal"},
			want: "inter",
		},
		{
			name: "All strings equal",
			strs: []string{"hello", "hello", "hello"},
			want: "hello",
		},
		{
			name: "Single string",
			strs: []string{"hello"},
			want: "hello",
		},
		{
			name: "Empty slice",
			strs: []string{},
			want: "",
		},
		{
			name: "Empty string",
			strs: []string{"", "flower", "flow"},
			want: "",
		},
		{
			name: "One string is prefix of another",
			strs: []string{"abc", "abcd", "abcde"},
			want: "abc",
		},
		{
			name: "Different first characters",
			strs: []string{"apple", "banana", "cherry"},
			want: "",
		},
		{
			name: "Common prefix is one character",
			strs: []string{"car", "cat", "camera"},
			want: "ca",
		},
		{
			name: "Long common prefix",
			strs: []string{"prefix123", "prefix456", "prefix789"},
			want: "prefix",
		},
		{
			name: "Two strings",
			strs: []string{"testing", "test"},
			want: "test",
		},
		{
			name: "Repeated strings",
			strs: []string{"abc", "abc", "abd"},
			want: "ab",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.LongestCommonPrefix14_v2(tt.strs)

			if got != tt.want {
				t.Errorf(
					"LongestCommonPrefix14_v2(%v) = %q; want %q",
					tt.strs,
					got,
					tt.want,
				)
			}
		})
	}
}
