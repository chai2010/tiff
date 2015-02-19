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
func (p *bitsReader) ReadBits(n uint) uint32 {
	for p.nbits < n {
		p.v <<= 8
		p.v |= uint32(p.buf[p.off])
		p.off++
		p.nbits += 8
	}
	p.nbits -= n
	rv := p.v >> p.nbits
	p.v &^= rv << p.nbits
	return rv
}
