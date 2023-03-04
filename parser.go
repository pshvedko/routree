package routree

import (
	"bytes"
	"io"
)

// Parse ...
func Parse(r io.ByteReader) ([][]uint16, error) {
	rr, err := splitPattern(r)
	if err != nil {
		return nil, err
	}
	var pp [][]uint16
	var p []uint16
	for _, r = range rr {
		p, err = readPattern(r)
		if err != nil {
			return nil, err
		}
		pp = append(pp, p)
	}
	return pp, nil
}

type innerReader struct {
	r io.ByteReader
}

func (r innerReader) ReadByte() (byte, error) {
	c, err := r.r.ReadByte()
	if err != nil {
		return 0, err
	}
	switch c {
	case ')':
		return 0, io.EOF
	default:
		return c, nil
	}
}

type byteReader interface {
	io.ByteReader
	Bytes() []byte
}

func splitPattern(r io.ByteReader) ([]byteReader, error) {
	bb := []*bytes.Buffer{bytes.NewBuffer([]byte{})}
	var (
		rr, pp []byteReader
	)
	for {
		c, err := r.ReadByte()
		switch err {
		case nil:
			switch c {
			case '(':
				pp, err = splitPattern(innerReader{r: r})
				if err != nil {
					return nil, err
				}
				bb, err = makeBuffer(bb, pp)
				if err != nil {
					return nil, err
				}
			case ',':
				rr, err = joinBuffer(rr, bb...)
				if err != nil {
					return nil, err
				}
				bb = []*bytes.Buffer{bytes.NewBuffer([]byte{})}
			default:
				for _, b := range bb {
					b.WriteByte(c)
				}
			}
		case io.EOF:
			return joinBuffer(rr, bb...)
		default:
			return nil, err
		}
	}
}

func makeBuffer(bb []*bytes.Buffer, pp []byteReader) (qq []*bytes.Buffer, err error) {
	for _, b := range bb {
		for _, p := range pp {
			q := bytes.NewBuffer(b.Bytes())
			_, err = q.Write(p.Bytes())
			if err != nil {
				return
			}
			qq = append(qq, q)
		}
	}
	return
}

func joinBuffer(rr []byteReader, bb ...*bytes.Buffer) ([]byteReader, error) {
	for _, b := range bb {
		if b.Len() == 0 {
			return nil, io.ErrUnexpectedEOF
		}
		rr = append(rr, b)
	}
	return rr, nil
}
