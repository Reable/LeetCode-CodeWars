package tests

import (
	"golang/tasks"
	"testing"
)

func TestReverseStr541(t *testing.T) {
	impls := []struct {
		name string
		fn   func(s string, k int) string
	}{
		{"v1", tasks.ReverseStr541_v1},
		{"v2", tasks.ReverseStr541_v2},
	}

	tests := []struct {
		name     string
		s        string
		k        int
		expected string
	}{
		{"пример 1 из условия", "abcdefg", 2, "bacdfeg"},
		{"пример 2 из условия", "abcd", 2, "bacd"},
		{"пустая строка", "", 2, ""},
		{"один символ", "a", 1, "a"},
		{"k = 1, ничего не меняется", "abcdef", 1, "abcdef"},
		{"k больше длины строки", "ab", 3, "ba"},
		{"длина равна k", "abc", 3, "cba"},
		{"длина равна 2k", "abcdef", 3, "cbadef"},
		{"хвост меньше k разворачивается целиком", "abcdefgh", 3, "cbadefhg"},
		{"хвост из одного символа", "abcde", 2, "bacde"},
		{"хвост ровно k символов", "abcdef", 2, "bacdfe"},
		{"хвост между k и 2k, разворот только первых k", "abcdefg", 3, "cbadefg"},
		{"два блока 2k и хвост из одного символа", "abcdefghi", 2, "bacdfeghi"},
		{"разворот только первой половины блока", "abcdefghij", 5, "edcbafghij"},
		{"хвост из трёх символов при k = 4", "abcdefghijk", 4, "dcbaefghkji"},
		{"длинная строка, несколько блоков", "abcdefghijklmnopqrstuvwxyz", 4, "dcbaefghlkjimnoptsrquvwxzy"},
	}

	for _, impl := range impls {
		for _, tt := range tests {
			t.Run(impl.name+"/"+tt.name, func(t *testing.T) {
				got := impl.fn(tt.s, tt.k)

				if got != tt.expected {
					t.Errorf("s=%q k=%d: got %q, want %q", tt.s, tt.k, got, tt.expected)
				}
			})
		}
	}
}

// Проверяем, что исходная строка не изменилась:
// это важно именно из-за unsafe.String.
func TestReverseStr541_DoesNotMutateInput(t *testing.T) {
	impls := map[string]func(string, int) string{
		"v1": tasks.ReverseStr541_v1,
		"v2": tasks.ReverseStr541_v2,
	}

	for name, fn := range impls {
		t.Run(name, func(t *testing.T) {
			original := "abcdefg"
			input := original

			_ = fn(input, 2)

			if input != original {
				t.Errorf("входная строка изменилась: got %q, want %q", input, original)
			}
		})
	}
}
