package hw02unpackstring

import (
	"strconv"
	"unicode"
)

// State - Определение состояний FSM
type State int

const (
	Start State = iota
	Normal
	Escape
)

// StringUnpacker - структура машины для разбора закодированной строки
type StringUnpacker struct {
	state       State
	decoded     []rune
	currentRune rune
}

// NewStringUnpacker - конструктор машины для разбора закодированной строки
func NewStringUnpacker() *StringUnpacker {
	return &StringUnpacker{
		state:   Start,
		decoded: make([]rune, 0),
	}
}

// processRune - обработка символа
func (su *StringUnpacker) processRune(r rune) error {
	switch su.state {
	case Start:
		switch {
		case r == '\\':
			su.state = Escape
			su.currentRune = r
		case !unicode.IsLetter(r):
			return ErrInvalidString
		default:
			su.state = Normal
			su.currentRune = r
		}
	case Normal:
		switch {
		case r == '\\':
			su.decoded = append(su.decoded, su.currentRune)
			su.state = Escape
		case unicode.IsDigit(r):
			cnt, _ := strconv.Atoi(string(r))
			su.writeTimes(su.currentRune, cnt)
			su.state = Start
		default:
			su.decoded = append(su.decoded, su.currentRune)
			su.currentRune = r
		}
	case Escape:
		su.currentRune = r
		su.state = Normal
	default:
		return ErrInvalidString
	}

	return nil
}

// writeTimes - запись символа times раз
func (su *StringUnpacker) writeTimes(symbol rune, times int) {
	for i := 0; i < times; i++ {
		su.decoded = append(su.decoded, symbol)
	}
}

// toString - преобразование результата в строку
func (su *StringUnpacker) toString() string {
	if su.state == Start {
		return string(su.decoded)
	}
	return string(su.decoded) + string(su.currentRune)
}
