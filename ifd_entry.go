// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"io"
)

type IFDEntry struct {
	IFD      *IFD
	Tag      uint16
	DataType DataType
	Count    uint32
	Offset   uint64
	Data     []byte
}

func (p *IFDEntry) ReadValue(r io.ReaderAt) (err error) {
	return
}

func (p *IFDEntry) GetInt8s() (data []int8) {
	return
}
func (p *IFDEntry) GetInt16s() (data []int16) {
	return
}
func (p *IFDEntry) GetInt32s() (data []int32) {
	return
}
func (p *IFDEntry) GetInt64s() (data []int64) {
	return
}
func (p *IFDEntry) GetRationals() (data [][2]int32) {
	return
}

func (p *IFDEntry) GetUInt8s() (data []uint8) {
	return
}
func (p *IFDEntry) GetUInt16s() (data []uint16) {
	return
}
func (p *IFDEntry) GetUInt32s() (data []uint32) {
	return
}
func (p *IFDEntry) GetUInt64s() (data []uint64) {
	return
}
func (p *IFDEntry) GetURationals() (data [][2]uint32) {
	return
}

func (p *IFDEntry) GetFloat32s() (data []float32) {
	return
}
func (p *IFDEntry) GetFloat64s() (data []float64) {
	return
}

func (p *IFDEntry) GetBytes() (data []byte) {
	return
}
func (p *IFDEntry) GetString() (data string) {
	return
}
