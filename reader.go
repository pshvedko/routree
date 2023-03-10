package routree

import (
	"io"
)

func readPattern(r io.ByteReader) ([]uint16, error) {
	var p []uint16
	for {
		u, err := readDigit(r)
		switch err {
		case nil:
			switch u {
			case 0x8000, 0x4000:
				if len(p) == 0 || u&p[len(p)-1] == u {
					return nil, errIllegalSymbol([...]byte{0, '#', '*'}[u>>14])
				}
				p[len(p)-1] |= u
			default:
				p = append(p, u)
			}
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
		return 0x3FF, nil
	case '#':
		return 0x4000, nil
	case '*':
		return readEnd(r)
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

func readEnd(r io.ByteReader) (uint16, error) {
	c, err := r.ReadByte()
	switch err {
	case nil:
		return 0, errIllegalSymbol(c)
	case io.EOF:
		return 0x8000, nil
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
	var d uint16
	for f <= l {
		d |= f
		f <<= 1
	}
	return d, nil
}

func makeDigit(c byte) (uint16, error) {
	if c < '0' || c > '9' {
		return 0, errIllegalSymbol(c)
	}
	return 1 << (c - '0'), nil
}
