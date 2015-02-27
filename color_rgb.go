// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"image/color"
)

type colorRGB struct {
	R, G, B uint8
}

func (c colorRGB) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R)<<8 | uint16(c.R))
	g = uint32(uint16(c.G)<<8 | uint16(c.G))
	b = uint32(uint16(c.B)<<8 | uint16(c.B))
	a = 0xFFFF
	return
}

func colorToRgbModel(c color.Color) color.Color {
	if c, ok := c.(colorRGB); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return colorRGB{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
	}
}

type colorRGB48 struct {
	R, G, B uint16
}

func (c colorRGB48) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = 0xFFFF
	return
}

func colorToRgb48Model(c color.Color) color.Color {
	if c, ok := c.(colorRGB48); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return colorRGB48{
		R: uint16(r),
		G: uint16(g),
		B: uint16(b),
	}
}

type colorRGB96i struct {
	R, G, B int32
}

func (c colorRGB96i) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = 0xFFFF
	return
}

func colorToRgb96iModel(c color.Color) color.Color {
	if c, ok := c.(colorRGB96i); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorRGB96i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
		}
	case colorGray32f:
		return colorRGB96i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
		}
	case colorGray64i:
		return colorRGB96i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
		}
	case colorGray64f:
		return colorRGB96i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
		}
	case colorGrayA64i:
		return colorRGB96i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
		}
	case colorGrayA64f:
		return colorRGB96i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
		}
	case colorGrayA128i:
		return colorRGB96i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
		}
	case colorGrayA128f:
		return colorRGB96i{
			R: int32(c.Y),
			G: int32(c.Y),
			B: int32(c.Y),
		}
	case colorRGB96i:
		return colorRGB96i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
		}
	case colorRGB96f:
		return colorRGB96i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
		}
	case colorRGB192i:
		return colorRGB96i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
		}
	case colorRGB192f:
		return colorRGB96i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
		}
	case colorRGBA128i:
		return colorRGB96i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
		}
	case colorRGBA128f:
		return colorRGB96i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
		}
	case colorRGBA256i:
		return colorRGB96i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
		}
	case colorRGBA256f:
		return colorRGB96i{
			R: int32(c.R),
			G: int32(c.G),
			B: int32(c.B),
		}
	}
	r, g, b, _ := c.RGBA()
	return colorRGB96i{
		R: int32(r),
		G: int32(g),
		B: int32(b),
	}
}

type colorRGB96f struct {
	R, G, B float32
}

func (c colorRGB96f) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = 0xFFFF
	return
}

func colorToRgb96fModel(c color.Color) color.Color {
	if c, ok := c.(colorRGB96f); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorRGB96f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
		}
	case colorGray32f:
		return colorRGB96f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
		}
	case colorGray64i:
		return colorRGB96f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
		}
	case colorGray64f:
		return colorRGB96f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
		}
	case colorGrayA64i:
		return colorRGB96f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
		}
	case colorGrayA64f:
		return colorRGB96f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
		}
	case colorGrayA128i:
		return colorRGB96f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
		}
	case colorGrayA128f:
		return colorRGB96f{
			R: float32(c.Y),
			G: float32(c.Y),
			B: float32(c.Y),
		}
	case colorRGB96i:
		return colorRGB96f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
		}
	case colorRGB96f:
		return colorRGB96f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
		}
	case colorRGB192i:
		return colorRGB96f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
		}
	case colorRGB192f:
		return colorRGB96f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
		}
	case colorRGBA128i:
		return colorRGB96f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
		}
	case colorRGBA128f:
		return colorRGB96f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
		}
	case colorRGBA256i:
		return colorRGB96f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
		}
	case colorRGBA256f:
		return colorRGB96f{
			R: float32(c.R),
			G: float32(c.G),
			B: float32(c.B),
		}
	}
	r, g, b, _ := c.RGBA()
	return colorRGB96f{
		R: float32(r),
		G: float32(g),
		B: float32(b),
	}
}

type colorRGB192i struct {
	R, G, B int64
}

func (c colorRGB192i) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = 0xFFFF
	return
}

func colorToRgb192iModel(c color.Color) color.Color {
	if c, ok := c.(colorRGB192i); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorRGB192i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
		}
	case colorGray32f:
		return colorRGB192i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
		}
	case colorGray64i:
		return colorRGB192i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
		}
	case colorGray64f:
		return colorRGB192i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
		}
	case colorGrayA64i:
		return colorRGB192i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
		}
	case colorGrayA64f:
		return colorRGB192i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
		}
	case colorGrayA128i:
		return colorRGB192i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
		}
	case colorGrayA128f:
		return colorRGB192i{
			R: int64(c.Y),
			G: int64(c.Y),
			B: int64(c.Y),
		}
	case colorRGB96i:
		return colorRGB192i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
		}
	case colorRGB96f:
		return colorRGB192i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
		}
	case colorRGB192i:
		return colorRGB192i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
		}
	case colorRGB192f:
		return colorRGB192i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
		}
	case colorRGBA128i:
		return colorRGB192i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
		}
	case colorRGBA128f:
		return colorRGB192i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
		}
	case colorRGBA256i:
		return colorRGB192i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
		}
	case colorRGBA256f:
		return colorRGB192i{
			R: int64(c.R),
			G: int64(c.G),
			B: int64(c.B),
		}
	}
	r, g, b, _ := c.RGBA()
	return colorRGB192i{
		R: int64(r),
		G: int64(g),
		B: int64(b),
	}
}

type colorRGB192f struct {
	R, G, B float64
}

func (c colorRGB192f) RGBA() (r, g, b, a uint32) {
	r = uint32(uint16(c.R))
	g = uint32(uint16(c.G))
	b = uint32(uint16(c.B))
	a = 0xFFFF
	return
}

func colorToRgb192fModel(c color.Color) color.Color {
	if c, ok := c.(colorRGB192f); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorRGB192f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
		}
	case colorGray32f:
		return colorRGB192f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
		}
	case colorGray64i:
		return colorRGB192f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
		}
	case colorGray64f:
		return colorRGB192f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
		}
	case colorGrayA64i:
		return colorRGB192f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
		}
	case colorGrayA64f:
		return colorRGB192f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
		}
	case colorGrayA128i:
		return colorRGB192f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
		}
	case colorGrayA128f:
		return colorRGB192f{
			R: float64(c.Y),
			G: float64(c.Y),
			B: float64(c.Y),
		}
	case colorRGB96i:
		return colorRGB192f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
		}
	case colorRGB96f:
		return colorRGB192f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
		}
	case colorRGB192i:
		return colorRGB192f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
		}
	case colorRGB192f:
		return colorRGB192f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
		}
	case colorRGBA128i:
		return colorRGB192f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
		}
	case colorRGBA128f:
		return colorRGB192f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
		}
	case colorRGBA256i:
		return colorRGB192f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
		}
	case colorRGBA256f:
		return colorRGB192f{
			R: float64(c.R),
			G: float64(c.G),
			B: float64(c.B),
		}
	}
	r, g, b, _ := c.RGBA()
	return colorRGB192f{
		R: float64(r),
		G: float64(g),
		B: float64(b),
	}
}
