// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"image/color"
)

type colorGray struct {
	Y uint8
}

func (c colorGray) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y)<<8 | uint16(c.Y))
	return y, y, y, 0xFFFF
}

func colorToGrayModel(c color.Color) color.Color {
	if c, ok := c.(colorGray); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return colorGray{Y: uint8(y >> 8)}
}

type colorGray16 struct {
	Y uint16
}

func (c colorGray16) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	return y, y, y, 0xFFFF
}

func colorToGray16Model(c color.Color) color.Color {
	if c, ok := c.(colorGray16); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return colorGray16{Y: uint16(y)}
}

type colorGray32i struct {
	Y int32
}

func (c colorGray32i) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	return y, y, y, 0xFFFF
}

func colorToGray32iModel(c color.Color) color.Color {
	if c, ok := c.(colorGray32i); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorGray32i{
			Y: int32(c.Y),
		}
	case colorGray32f:
		return colorGray32i{
			Y: int32(c.Y),
		}
	case colorGray64i:
		return colorGray32i{
			Y: int32(c.Y),
		}
	case colorGray64f:
		return colorGray32i{
			Y: int32(c.Y),
		}
	case colorGrayA64i:
		return colorGray32i{
			Y: int32(c.Y),
		}
	case colorGrayA64f:
		return colorGray32i{
			Y: int32(c.Y),
		}
	case colorGrayA128i:
		return colorGray32i{
			Y: int32(c.Y),
		}
	case colorGrayA128f:
		return colorGray32i{
			Y: int32(c.Y),
		}
	case colorRGB96i:
		return colorGray32i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
		}
	case colorRGB96f:
		return colorGray32i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
		}
	case colorRGB192i:
		return colorGray32i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
		}
	case colorRGB192f:
		return colorGray32i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
		}
	case colorRGBA128i:
		return colorGray32i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
		}
	case colorRGBA128f:
		return colorGray32i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
		}
	case colorRGBA256i:
		return colorGray32i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
		}
	case colorRGBA256f:
		return colorGray32i{
			Y: colorRgbToGrayI32(int32(c.R), int32(c.G), int32(c.B)),
		}
	}
	r, g, b, _ := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return colorGray32i{Y: int32(y)}
}

type colorGray32f struct {
	Y float32
}

func (c colorGray32f) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	return y, y, y, 0xFFFF
}

func colorToGray32fModel(c color.Color) color.Color {
	if c, ok := c.(colorGray32f); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorGray32f{
			Y: float32(c.Y),
		}
	case colorGray32f:
		return colorGray32f{
			Y: float32(c.Y),
		}
	case colorGray64i:
		return colorGray32f{
			Y: float32(c.Y),
		}
	case colorGray64f:
		return colorGray32f{
			Y: float32(c.Y),
		}
	case colorGrayA64i:
		return colorGray32f{
			Y: float32(c.Y),
		}
	case colorGrayA64f:
		return colorGray32f{
			Y: float32(c.Y),
		}
	case colorGrayA128i:
		return colorGray32f{
			Y: float32(c.Y),
		}
	case colorGrayA128f:
		return colorGray32f{
			Y: float32(c.Y),
		}
	case colorRGB96i:
		return colorGray32f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
		}
	case colorRGB96f:
		return colorGray32f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
		}
	case colorRGB192i:
		return colorGray32f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
		}
	case colorRGB192f:
		return colorGray32f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
		}
	case colorRGBA128i:
		return colorGray32f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
		}
	case colorRGBA128f:
		return colorGray32f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
		}
	case colorRGBA256i:
		return colorGray32f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
		}
	case colorRGBA256f:
		return colorGray32f{
			Y: colorRgbToGrayF32(float32(c.R), float32(c.G), float32(c.B)),
		}
	}
	r, g, b, _ := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return colorGray32f{
		Y: float32(y),
	}
}

type colorGray64i struct {
	Y int64
}

func (c colorGray64i) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	return y, y, y, 0xFFFF
}

func colorToGray64iModel(c color.Color) color.Color {
	if c, ok := c.(colorGray64i); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorGray64i{
			Y: int64(c.Y),
		}
	case colorGray32f:
		return colorGray64i{
			Y: int64(c.Y),
		}
	case colorGray64i:
		return colorGray64i{
			Y: int64(c.Y),
		}
	case colorGray64f:
		return colorGray64i{
			Y: int64(c.Y),
		}
	case colorGrayA64i:
		return colorGray64i{
			Y: int64(c.Y),
		}
	case colorGrayA64f:
		return colorGray64i{
			Y: int64(c.Y),
		}
	case colorGrayA128i:
		return colorGray64i{
			Y: int64(c.Y),
		}
	case colorGrayA128f:
		return colorGray64i{
			Y: int64(c.Y),
		}
	case colorRGB96i:
		return colorGray64i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
		}
	case colorRGB96f:
		return colorGray64i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
		}
	case colorRGB192i:
		return colorGray64i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
		}
	case colorRGB192f:
		return colorGray64i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
		}
	case colorRGBA128i:
		return colorGray64i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
		}
	case colorRGBA128f:
		return colorGray64i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
		}
	case colorRGBA256i:
		return colorGray64i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
		}
	case colorRGBA256f:
		return colorGray64i{
			Y: colorRgbToGrayI64(int64(c.R), int64(c.G), int64(c.B)),
		}
	}
	r, g, b, _ := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return colorGray64i{
		Y: int64(y),
	}
}

type colorGray64f struct {
	Y float64
}

func (c colorGray64f) RGBA() (r, g, b, a uint32) {
	y := uint32(uint16(c.Y))
	return y, y, y, 0xFFFF
}

func colorToGray64fModel(c color.Color) color.Color {
	if c, ok := c.(colorGray64f); ok {
		return c
	}
	switch c := c.(type) {
	case colorGray32i:
		return colorGray64f{
			Y: float64(c.Y),
		}
	case colorGray32f:
		return colorGray64f{
			Y: float64(c.Y),
		}
	case colorGray64i:
		return colorGray64f{
			Y: float64(c.Y),
		}
	case colorGray64f:
		return colorGray64f{
			Y: float64(c.Y),
		}
	case colorGrayA64i:
		return colorGray64f{
			Y: float64(c.Y),
		}
	case colorGrayA64f:
		return colorGray64f{
			Y: float64(c.Y),
		}
	case colorGrayA128i:
		return colorGray64f{
			Y: float64(c.Y),
		}
	case colorGrayA128f:
		return colorGray64f{
			Y: float64(c.Y),
		}
	case colorRGB96i:
		return colorGray64f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
		}
	case colorRGB96f:
		return colorGray64f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
		}
	case colorRGB192i:
		return colorGray64f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
		}
	case colorRGB192f:
		return colorGray64f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
		}
	case colorRGBA128i:
		return colorGray64f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
		}
	case colorRGBA128f:
		return colorGray64f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
		}
	case colorRGBA256i:
		return colorGray64f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
		}
	case colorRGBA256f:
		return colorGray64f{
			Y: colorRgbToGrayF64(float64(c.R), float64(c.G), float64(c.B)),
		}
	}
	r, g, b, _ := c.RGBA()
	y := colorRgbToGray(r, g, b)
	return colorGray64f{
		Y: float64(y),
	}
}
