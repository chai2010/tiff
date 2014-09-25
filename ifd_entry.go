// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"io"
)

type TagType uint16

type IFDEntry struct {
	IFD      *IFD
	Tag      TagType
	DataType DataType
	Count    uint64
	Offset   uint64
	Data     []byte
}

func (p *IFDEntry) ReadValue(r io.ReaderAt) (err error) {
	// if data size < sizeof(offset)
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

func (p *IFDEntry) SetInt8s(...int8) {
	return
}
func (p *IFDEntry) SetInt16s(...int16) {
	return
}
func (p *IFDEntry) SetInt32s(...int32) {
	return
}
func (p *IFDEntry) SetInt64s(...int64) {
	return
}
func (p *IFDEntry) SetRationals(...[2]int32) {
	return
}

func (p *IFDEntry) SetUInt8s(...uint8) {
	return
}
func (p *IFDEntry) SetUInt16s(...uint16) {
	return
}
func (p *IFDEntry) SetUInt32s(...uint32) {
	return
}
func (p *IFDEntry) SetUInt64s(...uint64) {
	return
}
func (p *IFDEntry) SetURationals(...[2]uint32) {
	return
}

func (p *IFDEntry) SetFloat32s(...float32) {
	return
}
func (p *IFDEntry) SetFloat64s(...float64) {
	return
}

func (p *IFDEntry) SetBytes(data []byte) {
	return
}
func (p *IFDEntry) SetString(data string) {
	return
}
