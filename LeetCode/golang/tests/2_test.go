package tests

import (
	"golang/tasks"
	"reflect"
	"testing"
)

// helper: массив -> список
func buildList(nums []int) *tasks.ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &tasks.ListNode{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &tasks.ListNode{Val: nums[i]}
		current = current.Next
	}

	return head
}

// helper: список -> массив
func listToSlice(node *tasks.ListNode) []int {
	var result []int

	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}

	return result
}

func TestAddTwoNumbers2(t *testing.T) {
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		{
			name:     "пример из условия",
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		{
			name:     "разная длина",
			l1:       []int{1, 8},
			l2:       []int{0},
			expected: []int{1, 8},
		},
		{
			name:     "перенос в новый разряд",
			l1:       []int{9, 9, 9},
			l2:       []int{1},
			expected: []int{0, 0, 0, 1},
		},
		{
			name:     "оба нули",
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			name:     "один список пустой",
			l1:       []int{},
			l2:       []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "оба пустые",
			l1:       []int{},
			l2:       []int{},
			expected: []int{0},
		},
		{
			name:     "длинный перенос",
			l1:       []int{9, 9, 9, 9},
			l2:       []int{9, 9, 9},
			expected: []int{8, 9, 9, 0, 1},
		},
		{
			name:     "без переноса",
			l1:       []int{1, 2, 3},
			l2:       []int{4, 5, 6},
			expected: []int{5, 7, 9},
		},
		{
			name:     "разные длины с переносом",
			l1:       []int{2, 4, 9},
			l2:       []int{5, 6, 4, 9},
			expected: []int{7, 0, 4, 0, 1},
		},
		{
			name:     "один элемент с переносом",
			l1:       []int{5},
			l2:       []int{5},
			expected: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := buildList(tt.l1)
			l2 := buildList(tt.l2)

			result := tasks.AddTwoNumbers2(l1, l2)
			got := listToSlice(result)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
