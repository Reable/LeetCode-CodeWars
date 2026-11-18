package tests

import (
	"golang/tasks"
	"testing"
)

func TestDivisorSubstrings2269(t *testing.T) {
	tests := []struct {
		name string
		num  int
		k    int
		want int
	}{
		{
			name: "Example 1",
			num:  240,
			k:    2,
			want: 2,
		},
		{
			name: "Example 2",
			num:  430043,
			k:    2,
			want: 2,
		},
		{
			name: "Example 3",
			num:  30003,
			k:    3,
			want: 1,
		},
		{
			name: "Single digit",
			num:  7,
			k:    1,
			want: 1,
		},
		{
			name: "K greater than number length",
			num:  12,
			k:    3,
			want: 0,
		},
		{
			name: "Zero substring",
			num:  101,
			k:    1,
			want: 2,
		},
		{
			name: "All zeros except one",
			num:  1000,
			k:    2,
			want: 1,
		},
		{
			name: "K equals number length",
			num:  12,
			k:    2,
			want: 1,
		},
		{
			name: "No divisors",
			num:  1234,
			k:    2,
			want: 0,
		},
		{
			name: "Repeated divisors",
			num:  111,
			k:    1,
			want: 3,
		},
		{
			name: "Last window is divisor",
			num:  120,
			k:    2,
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.DivisorSubstrings2269(tt.num, tt.k)

			if got != tt.want {
				t.Errorf(
					"DivisorSubstrings2269(%d, %d) = %d; want %d",
					tt.num,
					tt.k,
					got,
					tt.want,
				)
			}
		})
	}
}
