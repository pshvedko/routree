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

type byteReadWriters []byteReadWriter

func (rw *byteReadWriters) append(bb ...*bytes.Buffer) (byteReadWriters, error) {
	for _, b := range bb {
		if b.Len() == 0 {
			return nil, io.ErrUnexpectedEOF
		}
		*rw = append(*rw, b)
	}
	return *rw, nil
}

type byteReadWriter interface {
	io.ByteReader
	io.ByteWriter
	io.ReaderFrom
	io.Writer
	Bytes() []byte
}

func splitPattern(r io.ByteReader) ([]byteReadWriter, error) {
	bb := []*bytes.Buffer{bytes.NewBuffer([]byte{})}
	var (
		rr, pp byteReadWriters
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
				bb, err = joinBuffer(bb, pp)
				if err != nil {
					return nil, err
				}
			case ',':
				_, err = rr.append(bb...)
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
			return rr.append(bb...)
		default:
			return nil, err
		}
	}
}

func joinBuffer(bb []*bytes.Buffer, pp byteReadWriters) (qq []*bytes.Buffer, err error) {
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
