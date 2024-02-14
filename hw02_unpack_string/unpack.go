package hw02unpackstring

import (
	"errors"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(encodedString string) (string, error) {
	if encodedString == "" {
		return "", nil
	}

	decoder := NewStringUnpacker()

	for _, r := range encodedString {
		if err := decoder.processRune(r); err != nil {
			return "", err
		}
	}

	return decoder.toString(), nil
}
