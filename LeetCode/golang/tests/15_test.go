package tests

import (
	"golang/tasks"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name: "пример 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name:     "нет решений",
			nums:     []int{1, 2, 3},
			expected: [][]int{},
		},
		{
			name:     "все нули",
			nums:     []int{0, 0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name: "смешанные числа",
			nums: []int{-2, 0, 1, 1, 2},
			expected: [][]int{
				{-2, 0, 2},
				{-2, 1, 1},
			},
		},
		{
			name:     "меньше 3 элементов",
			nums:     []int{1, 2},
			expected: [][]int{},
		},
		{
			name: "дубликаты в начале",
			nums: []int{-1, -1, -1, 2, 2},
			expected: [][]int{
				{-1, -1, 2},
			},
		},
		{
			name: "с нулями и отрицательными",
			nums: []int{-4, -1, -1, 0, 1, 2},
			expected: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name:     "одинаковые элементы кроме одного",
			nums:     []int{3, 3, -6},
			expected: [][]int{{-6, 3, 3}},
		},
		{
			name:     "пустой массив",
			nums:     []int{},
			expected: [][]int{},
		},
		{
			name: "сложный случай",
			nums: []int{-2, -3, 0, 0, 2, 2, 3},
			expected: [][]int{
				{-3, 0, 3},
				{-2, 0, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tasks.ThreeSum15(tt.nums)

			if !equal3Sum(got, tt.expected) {
				t.Errorf("nums=%v got=%v expected=%v", tt.nums, got, tt.expected)
			}
		})
	}
}

// нормализуем ответ (порядок троек не важен)
func equal3Sum(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	normalize := func(arr [][]int) {
		for _, v := range arr {
			sort.Ints(v)
		}
		sort.Slice(arr, func(i, j int) bool {
			for k := 0; k < 3; k++ {
				if arr[i][k] != arr[j][k] {
					return arr[i][k] < arr[j][k]
				}
			}
			return false
		})
	}

	normalize(a)
	normalize(b)

	for i := range a {
		for j := 0; j < 3; j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}
