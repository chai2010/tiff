// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio

import (
	"io"
)

// Buffer buffers an io.Buffer to satisfy io.ReaderAt.
type Buffer struct {
	r   io.Reader
	buf []byte
}

func NewReaderAt(r io.Reader) io.ReaderAt {
	if ra, ok := r.(io.ReaderAt); ok {
		return ra
	}
	return &Buffer{
		r:   r,
		buf: make([]byte, 0, 1024),
	}
}

func NewReaderAtFromBuf(buf []byte) *Buffer {
	return &Buffer{buf: buf}
}

func (*Buffer) Read([]byte) (int, error) {
	panic("unimplemented")
}

func (b *Buffer) ReadAt(p []byte, off int64) (int, error) {
	o := int(off)
	end := o + len(p)
	if int64(end) != off+int64(len(p)) {
		return 0, io.ErrUnexpectedEOF
	}

	err := b.fill(end)
	return copy(p, b.buf[o:end]), err
}

// Slice returns a slice of the underlying Buffer. The slice contains
// n bytes starting at offset off.
func (b *Buffer) Slice(off, n int) ([]byte, error) {
	end := off + n
	if err := b.fill(end); err != nil {
		return nil, err
	}
	return b.buf[off:end], nil
}

func (b *Buffer) fill(end int) error {
	m := len(b.buf)
	if end > m {
		if end > cap(b.buf) {
			newcap := 1024
			for newcap < end {
				newcap *= 2
			}
			newbuf := make([]byte, end, newcap)
			copy(newbuf, b.buf)
			b.buf = newbuf
		} else {
			b.buf = b.buf[:end]
		}
		if n, err := io.ReadFull(b.r, b.buf[m:end]); err != nil {
			end = m + n
			b.buf = b.buf[:end]
			return err
		}
	}
	return nil
}
