package tests

import (
	"golang/tasks"
	"testing"
)

func TestIsLongPressedName925(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		typed    string
		expected bool
	}{
		{
			name:     "пример 1",
			input:    "alex",
			typed:    "aaleex",
			expected: true,
		},
		{
			name:     "пример 2",
			input:    "saeed",
			typed:    "ssaaedd",
			expected: false,
		},
		{
			name:     "дополнительные повторы",
			input:    "alex",
			typed:    "aaleeeex",
			expected: true,
		},
		{
			name:     "лишний символ в конце",
			input:    "alex",
			typed:    "aaleexa",
			expected: false,
		},
		{
			name:     "точное совпадение",
			input:    "alex",
			typed:    "alex",
			expected: true,
		},
		{
			name:     "пустое имя",
			input:    "",
			typed:    "alex",
			expected: false,
		},
		{
			name:     "пустая введённая строка",
			input:    "alex",
			typed:    "",
			expected: false,
		},
		{
			name:     "обе строки пустые",
			input:    "",
			typed:    "",
			expected: false,
		},
		{
			name:     "typed короче name",
			input:    "alex",
			typed:    "ale",
			expected: false,
		},
		{
			name:     "повтор первого символа",
			input:    "alex",
			typed:    "aaalex",
			expected: true,
		},
		{
			name:     "повтор последнего символа",
			input:    "alex",
			typed:    "alexxx",
			expected: true,
		},
		{
			name:     "пропущен символ имени",
			input:    "alex",
			typed:    "aex",
			expected: false,
		},
		{
			name:     "лишний символ в середине",
			input:    "alex",
			typed:    "allexx",
			expected: true,
		},
		{
			name:     "повтор не того символа",
			input:    "alex",
			typed:    "aallx",
			expected: false,
		},
		{
			name:     "один символ с повторами",
			input:    "a",
			typed:    "aaaa",
			expected: true,
		},
		{
			name:     "один символ без совпадения",
			input:    "a",
			typed:    "b",
			expected: false,
		},
		{
			name:     "повторяющиеся символы в имени",
			input:    "heeellooo",
			typed:    "heeellooooo",
			expected: true,
		},
		{
			name:     "неверный порядок символов",
			input:    "abc",
			typed:    "acb",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tasks.IsLongPressedName925(tt.input, tt.typed)

			if result != tt.expected {
				t.Errorf(
					"name=%q typed=%q: got %v, want %v",
					tt.input,
					tt.typed,
					result,
					tt.expected,
				)
			}
		})
	}
}
