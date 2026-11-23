package tests

import (
	"golang/tasks"
	"testing"
)

func TestProductOfNumbers1352(t *testing.T) {
	product := tasks.Constructor1352()

	// Add numbers
	product.Add1352(3)
	product.Add1352(0)
	product.Add1352(2)
	product.Add1352(5)
	product.Add1352(4)

	tests := []struct {
		name string
		k    int
		want int
	}{
		{
			name: "last 2 numbers",
			k:    2,
			want: 20,
		},
		{
			name: "last 3 numbers",
			k:    3,
			want: 40,
		},
		{
			name: "last 4 numbers with zero",
			k:    4,
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := product.GetProduct1352(tt.k)

			if got != tt.want {
				t.Errorf("GetProduct1352(%d) = %d, want %d",
					tt.k, got, tt.want)
			}
		})
	}
}

func TestProductOfNumbers1352WithoutZero(t *testing.T) {
	product := tasks.Constructor1352()

	product.Add1352(1)
	product.Add1352(2)
	product.Add1352(3)
	product.Add1352(4)

	tests := []struct {
		k    int
		want int
	}{
		{k: 1, want: 4},
		{k: 2, want: 12},
		{k: 3, want: 24},
		{k: 4, want: 24},
	}

	for _, tt := range tests {
		got := product.GetProduct1352(tt.k)

		if got != tt.want {
			t.Errorf("GetProduct1352(%d) = %d, want %d",
				tt.k, got, tt.want)
		}
	}
}

func TestProductOfNumbers1352AfterZero(t *testing.T) {
	product := tasks.Constructor1352()

	product.Add1352(2)
	product.Add1352(0)
	product.Add1352(3)
	product.Add1352(4)

	tests := []struct {
		k    int
		want int
	}{
		{k: 1, want: 4},
		{k: 2, want: 12},
		{k: 3, want: 0},
		{k: 4, want: 0},
	}

	for _, tt := range tests {
		got := product.GetProduct1352(tt.k)

		if got != tt.want {
			t.Errorf("GetProduct1352(%d) = %d, want %d",
				tt.k, got, tt.want)
		}
	}
}
