package routree

import (
	"fmt"
)

type ErrIllegalRange struct {
	a byte
	b byte
}

func (e ErrIllegalRange) Error() string {
	return fmt.Sprintf("illegal range '%c-%c'", e.a, e.b)
}

func errIllegalRange(a, b byte) error {
	return ErrIllegalRange{a: a, b: b}
}

type ErrIllegalSymbol struct {
	c byte
}

func (e ErrIllegalSymbol) Error() string {
	return fmt.Sprintf("illegal symbol `%c`", e.c)
}

func errIllegalSymbol(c byte) error {
	return ErrIllegalSymbol{c: c}
}

func errMandatorySign(c byte) error {
	return fmt.Errorf("mandatory `+` sing: %w", ErrIllegalSymbol{c: c})
}

type ErrFormat int

func (e ErrFormat) Error() string {
	return fmt.Sprintf("illegal format `E%03d`", int(e))
}

func (e ErrFormat) Format(number string) (string, error) {
	switch e {
	case E164:
		if len(number) != 0 && number[0] != '+' {
			return "", errMandatorySign(number[0])
		}
		return number[1:], nil
	default:
		return "", e
	}
}

const E164 ErrFormat = 164
