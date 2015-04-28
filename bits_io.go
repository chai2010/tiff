// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

type bitsReader struct {
	buf   []byte
	off   int    // Current offset in buf.
	v     uint32 // Buffer value for reading with arbitrary bit depths.
	nbits uint   // Remaining number of bits in v.
}

func newBitsReader(data []byte) *bitsReader {
	return &bitsReader{
		buf: data,
	}
}

// readBits reads n bits from the internal buffer starting at the current offset.
func (p *bitsReader) ReadBits(n uint) (v uint32, ok bool) {
	for p.nbits < n {
		p.v <<= 8
		if p.off >= len(p.buf) {
			return 0, false
		}
		p.v |= uint32(p.buf[p.off])
		p.off++
		p.nbits += 8
	}
	p.nbits -= n
	rv := p.v >> p.nbits
	p.v &^= rv << p.nbits
	return rv, true
}

// flushBits discards the unread bits in the buffer used by readBits.
// It is used at the end of a line.
func (p *bitsReader) flushBits() {
	p.v = 0
	p.nbits = 0
}
