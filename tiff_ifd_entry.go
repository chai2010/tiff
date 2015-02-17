// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"io"
)

type IFDEntry struct {
	Header   *Header
	Tag      TagType
	DataType DataType
	Count    uint64
	Offset   uint64
	Data     []byte
}

func (p *IFDEntry) Read(r io.Reader) (err error) {
	return
}

func (p *IFDEntry) Bytes() []byte {
	return nil
}

func (p *IFDEntry) String() string {
	return p.Tag.String()
}

func (p *IFDEntry) GetBytes() (data []byte) {
	return
}
func (p *IFDEntry) GetInts() (data []int64) {
	return
}
func (p *IFDEntry) GetUInts() (data []uint64) {
	return
}
func (p *IFDEntry) GetFloats() (data []float64) {
	return
}
func (p *IFDEntry) GetRationals() (data [][2]uint32) {
	return
}

func (p *IFDEntry) SetBytes(data []byte) {
	return
}
func (p *IFDEntry) SetInts(...int64) {
	return
}
func (p *IFDEntry) SetUInts(...uint64) {
	return
}
func (p *IFDEntry) SetFloats(...float64) {
	return
}
func (p *IFDEntry) SetURationals(...[2]uint32) {
	return
}
