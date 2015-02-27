// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"image/color"
)

type colorGrayA struct {
	Y, A uint8
}

func (c colorGrayA) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y)<<8 | uint16(c.Y))
	a = uint32(uint16(c.A)<<8 | uint16(c.A))
	return y, y, y, a
}

func colorToGrayAModel(c color.Color) color.Color {
	if c, ok := c.(colorGrayA); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return colorGrayA{
		Y: uint8(y >> 8),
		A: 0xFF,
	}
}

type colorGrayA32 struct {
	Y, A uint16
}

func (c colorGrayA32) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	a = uint32(uint16(c.A))
	return y, y, y, a
}

func colorToGrayA32Model(c color.Color) color.Color {
	if c, ok := c.(colorGrayA32); ok {
		return c
	}
	r, g, b, a := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return colorGrayA32{
		Y: uint16(y),
		A: uint16(a),
	}
}

type colorGrayA64i struct {
	Y, A int32
}

func (c colorGrayA64i) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	a = uint32(uint16(c.A))
	return y, y, y, a
}

func colorToGrayA64iModel(c color.Color) color.Color {
	if c, ok := c.(colorGrayA64i); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorGrayA64i{
			Y: int32(c.Y),
			A: 0xFFFF,
		}
	case colorGray32f:
		return colorGrayA64i{
			Y: int32(c.Y),
			A: 0xFFFF,
		}
	case colorGray64i:
		return colorGrayA64i{
			Y: int32(c.Y),
			A: 0xFFFF,
		}
	case colorGray64f:
		return colorGrayA64i{
			Y: int32(c.Y),
			A: 0xFFFF,
		}
	case colorGrayA64i:
		return colorGrayA64i{
			Y: int32(c.Y),
			A: int32(c.A),
		}
	case colorGrayA64f:
		return colorGrayA64i{
			Y: int32(c.Y),
			A: int32(c.A),
		}
	case colorGrayA128i:
		return colorGrayA64i{
			Y: int32(c.Y),
			A: int32(c.A),
		}
	case colorGrayA128f:
		return colorGrayA64i{
			Y: int32(c.Y),
			A: int32(c.A),
		}
	case colorRGB96i:
		return colorGrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: 0xFFFF,
		}
	case colorRGB96f:
		return colorGrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: 0xFFFF,
		}
	case colorRGB192i:
		return colorGrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: 0xFFFF,
		}
	case colorRGB192f:
		return colorGrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: 0xFFFF,
		}
	case colorRGBA128i:
		return colorGrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: int32(c.A),
		}
	case colorRGBA128f:
		return colorGrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: int32(c.A),
		}
	case colorRGBA256i:
		return colorGrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: int32(c.A),
		}
	case colorRGBA256f:
		return colorGrayA64i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
			A: int32(c.A),
		}
	}
	r, g, b, a := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return colorGrayA64i{
		Y: int32(y),
		A: int32(a),
	}
}

type colorGrayA64f struct {
	Y, A float32
}

func (c colorGrayA64f) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	a = uint32(uint16(c.A))
	return y, y, y, a
}

func colorToGrayA64fModel(c color.Color) color.Color {
	if c, ok := c.(colorGrayA64f); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorGrayA64f{
			Y: float32(c.Y),
			A: 0xFFFF,
		}
	case colorGray32f:
		return colorGrayA64f{
			Y: float32(c.Y),
			A: 0xFFFF,
		}
	case colorGray64i:
		return colorGrayA64f{
			Y: float32(c.Y),
			A: 0xFFFF,
		}
	case colorGray64f:
		return colorGrayA64f{
			Y: float32(c.Y),
			A: 0xFFFF,
		}
	case colorGrayA64i:
		return colorGrayA64f{
			Y: float32(c.Y),
			A: float32(c.A),
		}
	case colorGrayA64f:
		return colorGrayA64f{
			Y: float32(c.Y),
			A: float32(c.A),
		}
	case colorGrayA128i:
		return colorGrayA64f{
			Y: float32(c.Y),
			A: float32(c.A),
		}
	case colorGrayA128f:
		return colorGrayA64f{
			Y: float32(c.Y),
			A: float32(c.A),
		}
	case colorRGB96i:
		return colorGrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: 0xFFFF,
		}
	case colorRGB96f:
		return colorGrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: 0xFFFF,
		}
	case colorRGB192i:
		return colorGrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: 0xFFFF,
		}
	case colorRGB192f:
		return colorGrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: 0xFFFF,
		}
	case colorRGBA128i:
		return colorGrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: float32(c.A),
		}
	case colorRGBA128f:
		return colorGrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: float32(c.A),
		}
	case colorRGBA256i:
		return colorGrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: float32(c.A),
		}
	case colorRGBA256f:
		return colorGrayA64f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
			A: float32(c.A),
		}
	}
	r, g, b, a := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return colorGrayA64f{
		Y: float32(y),
		A: float32(a),
	}
}

type colorGrayA128i struct {
	Y, A int64
}

func (c colorGrayA128i) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	a = uint32(uint16(c.A))
	return y, y, y, a
}

func colorToGrayA128iModel(c color.Color) color.Color {
	if c, ok := c.(colorGrayA128i); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorGrayA128i{
			Y: int64(c.Y),
			A: 0xFFFF,
		}
	case colorGray32f:
		return colorGrayA128i{
			Y: int64(c.Y),
			A: 0xFFFF,
		}
	case colorGray64i:
		return colorGrayA128i{
			Y: int64(c.Y),
			A: 0xFFFF,
		}
	case colorGray64f:
		return colorGrayA128i{
			Y: int64(c.Y),
			A: 0xFFFF,
		}
	case colorGrayA64i:
		return colorGrayA128i{
			Y: int64(c.Y),
			A: int64(c.A),
		}
	case colorGrayA64f:
		return colorGrayA128i{
			Y: int64(c.Y),
			A: int64(c.A),
		}
	case colorGrayA128i:
		return colorGrayA128i{
			Y: int64(c.Y),
			A: int64(c.A),
		}
	case colorGrayA128f:
		return colorGrayA128i{
			Y: int64(c.Y),
			A: int64(c.A),
		}
	case colorRGB96i:
		return colorGrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: 0xFFFF,
		}
	case colorRGB96f:
		return colorGrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: 0xFFFF,
		}
	case colorRGB192i:
		return colorGrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: 0xFFFF,
		}
	case colorRGB192f:
		return colorGrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: 0xFFFF,
		}
	case colorRGBA128i:
		return colorGrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: int64(c.A),
		}
	case colorRGBA128f:
		return colorGrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: int64(c.A),
		}
	case colorRGBA256i:
		return colorGrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: int64(c.A),
		}
	case colorRGBA256f:
		return colorGrayA128i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
			A: int64(c.A),
		}
	}
	r, g, b, a := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return colorGrayA128i{
		Y: int64(y),
		A: int64(a),
	}
}

type colorGrayA128f struct {
	Y, A float64
}

func (c colorGrayA128f) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	a = uint32(uint16(c.A))
	return y, y, y, a
}

func colorToGrayA128fModel(c color.Color) color.Color {
	if c, ok := c.(colorGrayA128f); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorGrayA128f{
			Y: float64(c.Y),
			A: 0xFFFF,
		}
	case colorGray32f:
		return colorGrayA128f{
			Y: float64(c.Y),
			A: 0xFFFF,
		}
	case colorGray64i:
		return colorGrayA128f{
			Y: float64(c.Y),
			A: 0xFFFF,
		}
	case colorGray64f:
		return colorGrayA128f{
			Y: float64(c.Y),
			A: 0xFFFF,
		}
	case colorGrayA64i:
		return colorGrayA128f{
			Y: float64(c.Y),
			A: float64(c.A),
		}
	case colorGrayA64f:
		return colorGrayA128f{
			Y: float64(c.Y),
			A: float64(c.A),
		}
	case colorGrayA128i:
		return colorGrayA128f{
			Y: float64(c.Y),
			A: float64(c.A),
		}
	case colorGrayA128f:
		return colorGrayA128f{
			Y: float64(c.Y),
			A: float64(c.A),
		}
	case colorRGB96i:
		return colorGrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: 0xFFFF,
		}
	case colorRGB96f:
		return colorGrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: 0xFFFF,
		}
	case colorRGB192i:
		return colorGrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: 0xFFFF,
		}
	case colorRGB192f:
		return colorGrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: 0xFFFF,
		}
	case colorRGBA128i:
		return colorGrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: float64(c.A),
		}
	case colorRGBA128f:
		return colorGrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: float64(c.A),
		}
	case colorRGBA256i:
		return colorGrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: float64(c.A),
		}
	case colorRGBA256f:
		return colorGrayA128f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
			A: float64(c.A),
		}
	}
	r, g, b, a := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return colorGrayA128f{
		Y: float64(y),
		A: float64(a),
	}
}
