package routree

import (
	"io"
)

func readPattern(r io.ByteReader) ([]uint16, error) {
	var p []uint16
	for {
		d, err := readDigit(r)
		switch err {
		case nil:
			p = append(p, d)
		case io.EOF:
			return p, nil
		default:
			return nil, err
		}
	}
}

func readDigit(r io.ByteReader) (uint16, error) {
	c, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return makeDigit(c)
	case '.':
		return 0x3ff, nil
	case '*':
		return readEOF(r)
	case '[':
		return readDigitFirst(r)
	default:
		return 0, errIllegalSymbol(c)
	}
}

func readDigitFirst(r io.ByteReader) (uint16, error) {
	c, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return readDigitNext(r, c)
	default:
		return 0, errIllegalSymbol(c)
	}
}

func readDigitNext(r io.ByteReader, a byte) (uint16, error) {
	c, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	switch c {
	case ']':
		return makeDigit(a)
	case '-':
		return readDigitLast(r, a)
	case '|':
		return joinDigit(r, a)
	default:
		return 0, errIllegalSymbol(c)
	}
}

func readDigitLast(r io.ByteReader, a byte) (uint16, error) {
	c, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return readDigitRange(r, a, c)
	default:
		return 0, errIllegalSymbol(c)
	}
}

func readDigitRange(r io.ByteReader, a, b byte) (uint16, error) {
	c, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	switch c {
	case ']':
		return makeDigitRange(a, b)
	case '|':
		return joinDigitRange(r, a, b)
	default:
		return 0, errIllegalSymbol(c)
	}
}

func readEOF(r io.ByteReader) (uint16, error) {
	c, err := r.ReadByte()
	switch err {
	case nil:
		return 0, errIllegalSymbol(c)
	case io.EOF:
		return 0, nil
	default:
		return 0, err
	}
}

func joinDigitRange(r io.ByteReader, a, b byte) (uint16, error) {
	f, err := readDigitFirst(r)
	if err != nil {
		return 0, err
	}
	l, err := makeDigitRange(a, b)
	if err != nil {
		return 0, err
	}
	return f | l, nil
}

func joinDigit(r io.ByteReader, a byte) (uint16, error) {
	f, err := readDigitFirst(r)
	if err != nil {
		return 0, err
	}
	l, err := makeDigit(a)
	if err != nil {
		return 0, err
	}
	return f | l, nil
}

func makeDigitRange(a, b byte) (uint16, error) {
	f, err := makeDigit(a)
	if err != nil {
		return 0, err
	}
	l, err := makeDigit(b)
	if err != nil {
		return 0, err
	}
	if f >= l {
		return 0, errIllegalRange(a, b)
	}
	var digit uint16
	for f < l {
		digit |= f
		f <<= 1
	}
	return digit | l, nil
}

func makeDigit(c byte) (uint16, error) {
	return 1 << (c - '0'), nil
}
