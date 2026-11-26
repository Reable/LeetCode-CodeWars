package tests

import (
	"golang/tasks"
	"testing"
)

func TestCanAliceWin3360(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "less than 10",
			n:    9,
			want: false,
		},
		{
			name: "exactly 10",
			n:    10,
			want: true,
		},
		{
			name: "alice wins with 11",
			n:    11,
			want: true,
		},
		{
			name: "alice wins with 18",
			n:    18,
			want: true,
		},
		{
			name: "bob wins with 19",
			n:    19,
			want: false,
		},
		{
			name: "alice wins with 27",
			n:    27,
			want: true,
		},
		{
			name: "alice wins with 28",
			n:    28,
			want: true,
		},
		{
			name: "bob wins with 37",
			n:    37,
			want: false,
		},
		{
			name: "bob wins with 38",
			n:    38,
			want: false,
		},
		{
			name: "large number",
			n:    100,
			want: true,
		},
		{
			name: "very large number",
			n:    1000,
			want: true,
		},
		{
			name: "zero",
			n:    0,
			want: false,
		},
		{
			name: "negative number",
			n:    -10,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.CanAliceWin3360(tt.n)

			if got != tt.want {
				t.Errorf(
					"CanAliceWin3360(%d) = %v, want %v",
					tt.n,
					got,
					tt.want,
				)
			}
		})
	}
}
