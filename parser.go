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
	for i := range rr {
		p, err = readPattern(&rr[i])
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

func splitPattern(r io.ByteReader) ([]bytes.Buffer, error) {
	var rr, pp, bb []bytes.Buffer
	bb = []bytes.Buffer{{}}
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
				rr, err = joinBuffer(rr, bb)
				if err != nil {
					return nil, err
				}
				bb = []bytes.Buffer{{}}
			default:
				for i := range bb {
					bb[i].WriteByte(c)
				}
			}
		case io.EOF:
			return joinBuffer(rr, bb)
		default:
			return nil, err
		}
	}
}

func makeBuffer(bb, pp []bytes.Buffer) ([]bytes.Buffer, error) {
	var qq []bytes.Buffer
	for _, b := range bb {
		for _, p := range pp {
			var q bytes.Buffer
			_, err := q.Write(b.Bytes())
			if err != nil {
				return nil, err
			}
			_, err = q.Write(p.Bytes())
			if err != nil {
				return nil, err
			}
			qq = append(qq, q)
		}
	}
	return qq, nil
}

func joinBuffer(rr, bb []bytes.Buffer) ([]bytes.Buffer, error) {
	for _, b := range bb {
		if b.Len() == 0 {
			return nil, io.ErrUnexpectedEOF
		}
		rr = append(rr, b)
	}
	return rr, nil
}
