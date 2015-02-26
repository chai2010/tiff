// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"image/color"
	"reflect"
)

var (
	_ color.Color = (*Color)(nil)
)

type Color struct {
	// Pixel value in big-endian order format.
	Pix []byte

	// 1:Gray, 2:GrayA, 3:RGB, 4:RGBA
	Channels int
	// Uint8/Uint16/Uint32/Uint64/Int32/Int64/Float32/Float64
	DataType reflect.Kind
}

func (p *Color) RGBA() (r, g, b, a uint32) {
	return
}

func (p *Color) GetInt() []int64 {
	return nil
}
func (p *Color) SetInt(v []int64) {
	return
}

func (p *Color) GetFloat() []float64 {
	return nil
}
func (p *Color) SetFloat(v []float64) {
	return
}

func (p *Color) GetUint8() []byte {
	return nil
}
func (p *Color) SetUint8(v []byte) {
	return
}

func (p *Color) GetUint16() []uint16 {
	return nil
}
func (p *Color) SetUint16(v []uint16) {
	return
}

func (p *Color) GetUint32() []uint32 {
	return nil
}
func (p *Color) SetUint32(v []uint32) {
	return
}

func (p *Color) GetUint64() []uint64 {
	return nil
}
func (p *Color) SetUint64(v []uint64) {
	return
}

func (p *Color) GetInt32() []int32 {
	return nil
}
func (p *Color) SetInt32(v []int32) {
	return
}

func (p *Color) GetInt64() []int64 {
	return nil
}
func (p *Color) SetInt64(v []int64) {
	return
}

func (p *Color) GetFloat32() []float32 {
	return nil
}
func (p *Color) SetFloat32(v []float32) {
	return
}

func (p *Color) GetFloat64() []float64 {
	return nil
}
func (p *Color) SetFloat64(v []float64) {
	return
}

var ColorModel = color.ModelFunc(func(c color.Color) color.Color {
	return c
})
