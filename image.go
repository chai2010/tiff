// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"image"
	"image/color"
	"reflect"
)

type Image struct {
	// Pix holds the image's pixels, as pixel values in big-endian order format. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*PixelSize].
	Pix []byte
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle

	// 1:Gray, 2:GrayA, 3:RGB, 4:RGBA
	Channels int
	// Uint8/Uint16/Uint32/Uint64/Int32/Int64/Float32/Float64
	DataType reflect.Kind
}

func NewImage(r image.Rectangle, channels int, dataType reflect.Kind) *Image {
	pixSize := getPixelSize(channels, dataType)
	stride := pixSize * r.Dx()
	return &Image{
		Pix:      make([]uint8, pixSize*r.Dy()),
		Stride:   stride,
		Rect:     r,
		Channels: channels,
		DataType: dataType,
	}
}

func (p *Image) Bounds() image.Rectangle {
	return p.Rect
}

func (p *Image) ColorModel() color.Model {
	return ColorModel
}

func (p *Image) At(x, y int) color.Color {
	return nil
}

func (p *Image) Set(x, y int, c color.Color) {
	//
}

func (p *Image) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*2
}

func (p *Image) SubImage(r image.Rectangle) image.Image {
	return nil
}

func (p *Image) StdImage() (image.Image, bool) {
	return nil, false
}

func (p *Image) Opaque() bool {
	return true
}
