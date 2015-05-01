// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type Header struct {
	ByteOrder binary.ByteOrder
	TiffType  TiffType
	FirstIFD  int64
}

func NewHeader(isBigTiff bool, firstIFD int64) *Header {
	if isBigTiff {
		return &Header{
			ByteOrder: binary.LittleEndian,
			TiffType:  TiffType_BigTIFF,
			FirstIFD:  firstIFD,
		}
	} else {
		return &Header{
			ByteOrder: binary.LittleEndian,
			TiffType:  TiffType_ClassicTIFF,
			FirstIFD:  firstIFD,
		}
	}
}

func ReadHeader(r io.Reader) (header *Header, err error) {
	var rs io.ReadSeeker
	if rs, _ = r.(io.ReadSeeker); rs == nil {
		seekReader := openSeekioReader(r, 0)
		defer seekReader.Close()
		rs = seekReader
	}
	if _, err = rs.Seek(0, 0); err != nil {
		return
	}

	// read classic TIFF header
	var data [8]byte
	if _, err = rs.Read(data[:8]); err != nil {
		return
	}
	header = new(Header)

	// byte order
	switch {
	case data[0] == 'I' && data[1] == 'I':
		header.ByteOrder = binary.LittleEndian
	case data[0] == 'M' && data[1] == 'M':
		header.ByteOrder = binary.BigEndian
	default:
		err = fmt.Errorf("tiff: ReadHeader, bad order: %v", data[:2])
		return
	}

	// version: ClassicTIFF or BigTIFF
	header.TiffType = TiffType(header.ByteOrder.Uint16(data[2:4]))
	if v := header.TiffType; v != TiffType_ClassicTIFF && v != TiffType_BigTIFF {
		err = fmt.Errorf("tiff: ReadHeader, bad version: %v", v)
		return
	}

	// offset
	switch header.TiffType {
	case TiffType_ClassicTIFF:
		header.FirstIFD = int64(header.ByteOrder.Uint32(data[4:8]))
	case TiffType_BigTIFF:
		byte46 := header.ByteOrder.Uint16(data[4:6])
		byte68 := header.ByteOrder.Uint16(data[6:8])
		if byte46 != 8 || byte68 != 0 {
			err = fmt.Errorf("tiff: ReadHeader, bad offset: %v", data[4:8])
			return
		}
		if _, err = rs.Read(data[:8]); err != nil {
			return
		}
		header.FirstIFD = int64(header.ByteOrder.Uint64(data[0:8]))
	}
	if header.FirstIFD < int64(header.HeadSize()) {
		err = fmt.Errorf("tiff: ReadHeader, bad offset: %v", data[4:8])
		return
	}

	return
}

func (p *Header) Bytes() []byte {
	if !p.Valid() {
		return nil
	}

	var d [16]byte
	switch p.ByteOrder {
	case binary.LittleEndian:
		d[0], d[1] = 'I', 'I'
	case binary.BigEndian:
		d[0], d[1] = 'M', 'M'
	}

	if p.TiffType == TiffType_ClassicTIFF {
		p.ByteOrder.PutUint16(d[2:4], uint16(p.TiffType))
		p.ByteOrder.PutUint32(d[4:8], uint32(p.FirstIFD))
		return d[:16]
	} else {
		p.ByteOrder.PutUint16(d[2:4], uint16(p.TiffType))
		p.ByteOrder.PutUint16(d[4:6], 8)
		p.ByteOrder.PutUint16(d[6:8], 0)
		p.ByteOrder.PutUint64(d[8:], uint64(p.FirstIFD))
		return d[:16]
	}
}

func (p *Header) Valid() bool {
	if p == nil {
		return false
	}
	if x := p.ByteOrder; x != binary.LittleEndian && x != binary.BigEndian {
		return false
	}
	if !p.TiffType.Valid() {
		return false
	}
	if p.FirstIFD < int64(p.HeadSize()) {
		return false
	}
	if p.TiffType == TiffType_ClassicTIFF {
		if p.FirstIFD > math.MaxUint32 {
			return false
		}
	}
	return true
}

func (p *Header) HeadSize() int {
	if p.TiffType == TiffType_ClassicTIFF {
		return 8
	}
	if p.TiffType == TiffType_BigTIFF {
		return 16
	}
	return 0
}

func (p *Header) IsBigTiff() bool {
	return p.TiffType == TiffType_BigTIFF
}

func (p *Header) String() string {
	orderName := "Unknown"
	switch p.ByteOrder {
	case binary.LittleEndian:
		orderName = "LittleEndian"
	case binary.BigEndian:
		orderName = "BigEndian"
	}
	return fmt.Sprintf(
		`tiff.Header{
	ByteOrder: %s,
	TiffType: %v,
	Offset: %#08x,
}`,
		orderName, p.TiffType, p.FirstIFD,
	)
}
