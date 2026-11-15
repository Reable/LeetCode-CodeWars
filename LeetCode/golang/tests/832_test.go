package tests

import (
	"golang/tasks"
	"reflect"
	"testing"
)

func TestFlipAndInvertImage832(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		want  [][]int
	}{
		{
			name: "basic example",
			input: [][]int{
				{1, 1, 0},
				{1, 0, 1},
				{0, 0, 0},
			},
			want: [][]int{
				{1, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
		},
		{
			name: "single row",
			input: [][]int{
				{1, 0, 1},
			},
			want: [][]int{
				{0, 1, 0},
			},
		},
		{
			name: "single column",
			input: [][]int{
				{1},
				{0},
				{1},
			},
			want: [][]int{
				{0},
				{1},
				{0},
			},
		},
		{
			name: "single element",
			input: [][]int{
				{1},
			},
			want: [][]int{
				{0},
			},
		},
		{
			name: "all zeros",
			input: [][]int{
				{0, 0},
				{0, 0},
			},
			want: [][]int{
				{1, 1},
				{1, 1},
			},
		},
		{
			name: "all ones",
			input: [][]int{
				{1, 1},
				{1, 1},
			},
			want: [][]int{
				{0, 0},
				{0, 0},
			},
		},
		{
			name:  "empty image",
			input: [][]int{},
			want:  [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.FlipAndInvertImage832(tt.input)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlipAndInvertImage832() = %v, want %v",
					got, tt.want)
			}
		})
	}
}
