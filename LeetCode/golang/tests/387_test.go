package tests

import (
	"golang/tasks"
	"testing"
)

func TestFirstUniqChar387(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "first character is unique",
			s:    "leetcode",
			want: 0,
		},
		{
			name: "unique character in the middle",
			s:    "loveleetcode",
			want: 2,
		},
		{
			name: "no unique characters",
			s:    "aabb",
			want: -1,
		},
		{
			name: "single character",
			s:    "a",
			want: 0,
		},
		{
			name: "all characters are unique",
			s:    "abcde",
			want: 0,
		},
		{
			name: "unique character at the end",
			s:    "aabbc",
			want: 4,
		},
		{
			name: "empty string",
			s:    "",
			want: -1,
		},
		{
			name: "multiple unique characters",
			s:    "abacbd",
			want: 3,
		},
		{
			name: "all characters repeated",
			s:    "aabbccddeeff",
			want: -1,
		},
		{
			name: "unique character between duplicates",
			s:    "aabbccd",
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV1 := tasks.FirstUniqChar387_v1(tt.s)
			gotV2 := tasks.FirstUniqChar387_v2(tt.s)

			if gotV1 != tt.want {
				t.Errorf(
					"v1: FirstUniqChar387_v1(%q) = %d; want %d",
					tt.s, gotV1, tt.want,
				)
			}

			if gotV2 != tt.want {
				t.Errorf(
					"v2: FirstUniqChar387_v2(%q) = %d; want %d",
					tt.s, gotV2, tt.want,
				)
			}

			if gotV1 != gotV2 {
				t.Errorf(
					"implementations differ: v1 = %d, v2 = %d",
					gotV1, gotV2,
				)
			}
		})
	}
}
