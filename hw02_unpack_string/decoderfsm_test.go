package hw02unpackstring

import (
	"errors"
	"testing"
)

func TestStringUnpacker(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		err      error
	}{
		{"Простая строка", "abc", "abc", nil},
		{"Строка с цифрами", "a2b3", "aabbb", nil},
		{"Строки с экранированием", `a\3`, "a3", nil},
		{"Смешанные строки", "a\\3b2", "a3bb", nil},
		{"Строка начинающаяся с числа", "3ab", "", ErrInvalidString},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			unpacker := NewStringUnpacker()
			var lastErr error
			for _, r := range test.input {
				if err := unpacker.processRune(r); err != nil {
					lastErr = err
					if test.err == nil {
						t.Fatalf("Неожиданная ошибка для символа '%c' в тесте '%s': %v", r, test.name, lastErr)
					}
					break
				}
			}
			if test.err != nil {
				if !errors.Is(lastErr, test.err) {
					t.Errorf("Неожиданная ошибка %v, в тесте '%s': %v", test.err, test.name, lastErr)
				}
				return
			}
			result := unpacker.toString()
			if result != test.expected {
				t.Errorf("Ожидалось '%s', получено '%s' в тесте '%s'", test.expected, result, test.name)
			}
		})
	}
}
