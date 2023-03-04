package routree

import "fmt"

type ErrIllegalRange struct {
	a byte
	b byte
}

func (e ErrIllegalRange) Error() string {
	return fmt.Sprintf("illegal range '%c-%c'", e.a, e.b)
}

func errIllegalSymbol(c byte) error {
	return ErrIllegalSymbol{c: c}
}

type ErrIllegalSymbol struct {
	c byte
}

func (e ErrIllegalSymbol) Error() string {
	return fmt.Sprintf("illegal symbol '%c'", e.c)
}

func errIllegalRange(a byte, b byte) error {
	return ErrIllegalRange{a: a, b: b}
}
