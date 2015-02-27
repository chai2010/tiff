// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"image/color"
)

type colorRGBA struct {
	R, G, B, A uint8
}

func (c colorRGBA) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R)<<8 | uint16(c.R))
	g = uint32(uint16(c.G)<<8 | uint16(c.G))
	b = uint32(uint16(c.B)<<8 | uint16(c.B))
	a = uint32(uint16(c.A)<<8 | uint16(c.A))
	return
}

func colorToRgbaModel(c color.Color) color.Color {
	if c, ok := c.(colorRGBA); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return colorRGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: uint8(a >> 8),
	}
}

type colorRGBA64 struct {
	R, G, B, A uint16
}

func (c colorRGBA64) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = uint32(uint16(c.A))
	return
}

func colorToRgba64Model(c color.Color) color.Color {
	if c, ok := c.(colorRGBA64); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	return colorRGBA64{
		R: uint16(r),
		G: uint16(g),
		B: uint16(b),
		A: uint16(a),
	}
}

type colorRGBA128i struct {
	R, G, B, A int32
}

func (c colorRGBA128i) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = uint32(uint16(c.A))
	return
}

func colorToRgba128iModel(c color.Color) color.Color {
	if c, ok := c.(colorRGBA128i); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorRGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: 0xFFFF,
		}
	case colorGray32f:
		return colorRGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: 0xFFFF,
		}
	case colorGray64i:
		return colorRGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: 0xFFFF,
		}
	case colorGray64f:
		return colorRGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: 0xFFFF,
		}
	case colorGrayA64i:
		return colorRGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: int32(c.A),
		}
	case colorGrayA64f:
		return colorRGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: int32(c.A),
		}
	case colorGrayA128i:
		return colorRGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: int32(c.A),
		}
	case colorGrayA128f:
		return colorRGBA128i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
			A: int32(c.A),
		}
	case colorRGB96i:
		return colorRGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: 0xFFFF,
		}
	case colorRGB96f:
		return colorRGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: 0xFFFF,
		}
	case colorRGB192i:
		return colorRGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: 0xFFFF,
		}
	case colorRGB192f:
		return colorRGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: 0xFFFF,
		}
	case colorRGBA128i:
		return colorRGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: int32(c.A),
		}
	case colorRGBA128f:
		return colorRGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: int32(c.A),
		}
	case colorRGBA256i:
		return colorRGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: int32(c.A),
		}
	case colorRGBA256f:
		return colorRGBA128i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
			A: int32(c.A),
		}
	}
	r, g, b, a := c.RGBA()
	return colorRGBA128i{
		R: int32(r),
		G: int32(g),
		B: int32(b),
		A: int32(a),
	}
}

type colorRGBA128f struct {
	R, G, B, A float32
}

func (c colorRGBA128f) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = uint32(uint16(c.A))
	return
}

func colorToRgba128fModel(c color.Color) color.Color {
	if c, ok := c.(colorRGBA128f); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorRGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: 0xFFFF,
		}
	case colorGray32f:
		return colorRGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: 0xFFFF,
		}
	case colorGray64i:
		return colorRGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: 0xFFFF,
		}
	case colorGray64f:
		return colorRGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: 0xFFFF,
		}
	case colorGrayA64i:
		return colorRGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: float32(c.A),
		}
	case colorGrayA64f:
		return colorRGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: float32(c.A),
		}
	case colorGrayA128i:
		return colorRGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: float32(c.A),
		}
	case colorGrayA128f:
		return colorRGBA128f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
			A: float32(c.A),
		}
	case colorRGB96i:
		return colorRGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: 0xFFFF,
		}
	case colorRGB96f:
		return colorRGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: 0xFFFF,
		}
	case colorRGB192i:
		return colorRGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: 0xFFFF,
		}
	case colorRGB192f:
		return colorRGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: 0xFFFF,
		}
	case colorRGBA128i:
		return colorRGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: float32(c.A),
		}
	case colorRGBA128f:
		return colorRGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: float32(c.A),
		}
	case colorRGBA256i:
		return colorRGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: float32(c.A),
		}
	case colorRGBA256f:
		return colorRGBA128f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
			A: float32(c.A),
		}
	}
	r, g, b, a := c.RGBA()
	return colorRGBA128f{
		R: float32(r),
		G: float32(g),
		B: float32(b),
		A: float32(a),
	}
}

type colorRGBA256i struct {
	R, G, B, A int64
}

func (c colorRGBA256i) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = uint32(uint16(c.A))
	return
}

func colorToRgba256iModel(c color.Color) color.Color {
	if c, ok := c.(colorRGBA256i); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorRGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: 0xFFFF,
		}
	case colorGray32f:
		return colorRGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: 0xFFFF,
		}
	case colorGray64i:
		return colorRGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: 0xFFFF,
		}
	case colorGray64f:
		return colorRGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: 0xFFFF,
		}
	case colorGrayA64i:
		return colorRGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: int64(c.A),
		}
	case colorGrayA64f:
		return colorRGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: int64(c.A),
		}
	case colorGrayA128i:
		return colorRGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: int64(c.A),
		}
	case colorGrayA128f:
		return colorRGBA256i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
			A: int64(c.A),
		}
	case colorRGB96i:
		return colorRGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: 0xFFFF,
		}
	case colorRGB96f:
		return colorRGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: 0xFFFF,
		}
	case colorRGB192i:
		return colorRGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: 0xFFFF,
		}
	case colorRGB192f:
		return colorRGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: 0xFFFF,
		}
	case colorRGBA128i:
		return colorRGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: int64(c.A),
		}
	case colorRGBA128f:
		return colorRGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: int64(c.A),
		}
	case colorRGBA256i:
		return colorRGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: int64(c.A),
		}
	case colorRGBA256f:
		return colorRGBA256i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
			A: int64(c.A),
		}
	}
	r, g, b, a := c.RGBA()
	return colorRGBA256i{
		R: int64(r),
		G: int64(g),
		B: int64(b),
		A: int64(a),
	}
}

type colorRGBA256f struct {
	R, G, B, A float64
}

func (c colorRGBA256f) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = uint32(uint16(c.A))
	return
}

func colorToRgba256fModel(c color.Color) color.Color {
	if c, ok := c.(colorRGBA256f); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorRGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: 0xFFFF,
		}
	case colorGray32f:
		return colorRGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: 0xFFFF,
		}
	case colorGray64i:
		return colorRGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: 0xFFFF,
		}
	case colorGray64f:
		return colorRGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: 0xFFFF,
		}
	case colorGrayA64i:
		return colorRGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: float64(c.A),
		}
	case colorGrayA64f:
		return colorRGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: float64(c.A),
		}
	case colorGrayA128i:
		return colorRGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: float64(c.A),
		}
	case colorGrayA128f:
		return colorRGBA256f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
			A: float64(c.A),
		}
	case colorRGB96i:
		return colorRGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: 0xFFFF,
		}
	case colorRGB96f:
		return colorRGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: 0xFFFF,
		}
	case colorRGB192i:
		return colorRGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: 0xFFFF,
		}
	case colorRGB192f:
		return colorRGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: 0xFFFF,
		}
	case colorRGBA128i:
		return colorRGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: float64(c.A),
		}
	case colorRGBA128f:
		return colorRGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: float64(c.A),
		}
	case colorRGBA256i:
		return colorRGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: float64(c.A),
		}
	case colorRGBA256f:
		return colorRGBA256f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
			A: float64(c.A),
		}
	}
	r, g, b, a := c.RGBA()
	return colorRGBA256f{
		R: float64(r),
		G: float64(g),
		B: float64(b),
		A: float64(a),
	}
}
