// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"reflect"
)

var (
	_ internalImage = (*imageGray)(nil)
	_ internalImage = (*imageGray16)(nil)
	_ internalImage = (*imageGray32i)(nil)
	_ internalImage = (*imageGray32f)(nil)
	_ internalImage = (*imageGray64i)(nil)
	_ internalImage = (*imageGray64f)(nil)
	_ internalImage = (*imageGrayA)(nil)
	_ internalImage = (*imageGrayA32)(nil)
	_ internalImage = (*imageGrayA64i)(nil)
	_ internalImage = (*imageGrayA64f)(nil)
	_ internalImage = (*imageGrayA128i)(nil)
	_ internalImage = (*imageGrayA128f)(nil)
	_ internalImage = (*imageRGB)(nil)
	_ internalImage = (*imageRGB48)(nil)
	_ internalImage = (*imageRGB96i)(nil)
	_ internalImage = (*imageRGB96f)(nil)
	_ internalImage = (*imageRGB192i)(nil)
	_ internalImage = (*imageRGB192f)(nil)
	_ internalImage = (*imageRGBA)(nil)
	_ internalImage = (*imageRGBA64)(nil)
	_ internalImage = (*imageRGBA128i)(nil)
	_ internalImage = (*imageRGBA128f)(nil)
	_ internalImage = (*imageRGBA256i)(nil)
	_ internalImage = (*imageRGBA256f)(nil)
)

type internalImage interface {
	// Pix holds the image's pixels, as pixel values in big-endian order format. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*PixelSize].
	Pix() []byte
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride() int
	// Rect is the image's bounds.
	Rect() image.Rectangle

	// 1:Gray, 2:GrayA, 3:RGB, 4:RGBA
	Channels() int
	// Uint8/Uint16/Int32/Int64/Float32/Float64
	Depth() reflect.Kind

	draw.Image
	SubImage(r image.Rectangle) image.Image
}

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
	// Uint8/Uint16/Int32/Int64/Float32/Float64
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

func NewImageWithIFD(r image.Rectangle, ifd *IFD) (m image.Image, err error) {
	switch ifd.ImageType() {
	case ImageType_Bilevel, ImageType_BilevelInvert:
		m = image.NewGray(r)
	case ImageType_Gray, ImageType_GrayInvert:
		if ifd.Depth() == 16 {
			m = image.NewGray16(r)
		} else {
			m = image.NewGray(r)
		}
	case ImageType_Paletted:
		m = image.NewPaletted(r, ifd.ColorMap())
	case ImageType_NRGBA:
		if ifd.Depth() == 16 {
			m = image.NewNRGBA64(r)
		} else {
			m = image.NewNRGBA(r)
		}
	case ImageType_RGB, ImageType_RGBA:
		if ifd.Depth() == 16 {
			m = image.NewRGBA64(r)
		} else {
			m = image.NewRGBA(r)
		}
	}
	if m == nil {
		err = fmt.Errorf("tiff: Decode, unknown format")
		return
	}
	return
}

func (p *Image) Bounds() image.Rectangle {
	return p.Rect
}

func (p *Image) ColorModel() color.Model {
	m, err := p.asInternalImage()
	if err != nil {
		return nil
	}
	return m.ColorModel()
}

func (p *Image) At(x, y int) color.Color {
	m, err := p.asInternalImage()
	if err != nil {
		return nil
	}
	return m.At(x, y)
}

func (p *Image) Set(x, y int, c color.Color) {
	m, err := p.asInternalImage()
	if err != nil {
		return
	}
	m.Set(x, y, c)
}

func (p *Image) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*2
}

func (p *Image) SubImage(r image.Rectangle) image.Image {
	m, err := p.asInternalImage()
	if err != nil {
		return nil
	}
	return m.SubImage(r)
}

func (p *Image) StdImage() image.Image {
	switch channels, depth := p.Channels, p.DataType; {
	case channels == 1 && depth == reflect.Uint8:
		return &image.Gray{
			Pix:    p.Pix,
			Stride: p.Stride,
			Rect:   p.Rect,
		}
	case channels == 1 && depth == reflect.Uint16:
		return &image.Gray16{
			Pix:    p.Pix,
			Stride: p.Stride,
			Rect:   p.Rect,
		}
	case channels == 4 && depth == reflect.Uint8:
		return &image.RGBA{
			Pix:    p.Pix,
			Stride: p.Stride,
			Rect:   p.Rect,
		}
	case channels == 4 && depth == reflect.Uint16:
		return &image.RGBA64{
			Pix:    p.Pix,
			Stride: p.Stride,
			Rect:   p.Rect,
		}
	}
	return p
}

func (p *Image) asInternalImage() (m internalImage, err error) {
	switch channels, depth := p.Channels, p.DataType; {
	case channels == 1 && depth == reflect.Uint8:
		m = new(imageGray).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 1 && depth == reflect.Uint16:
		m = new(imageGray16).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 1 && depth == reflect.Int32:
		m = new(imageGray32i).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 1 && depth == reflect.Float32:
		m = new(imageGray32f).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 1 && depth == reflect.Int64:
		m = new(imageGray64i).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 1 && depth == reflect.Float64:
		m = new(imageGray64f).Init(p.Pix, p.Stride, p.Rect)
		return

	case channels == 2 && depth == reflect.Uint8:
		m = new(imageGrayA).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 2 && depth == reflect.Uint16:
		m = new(imageGrayA32).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 2 && depth == reflect.Int32:
		m = new(imageGrayA64i).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 2 && depth == reflect.Float32:
		m = new(imageGrayA64f).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 2 && depth == reflect.Int64:
		m = new(imageGrayA128i).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 2 && depth == reflect.Float64:
		m = new(imageGrayA128f).Init(p.Pix, p.Stride, p.Rect)
		return

	case channels == 3 && depth == reflect.Uint8:
		m = new(imageRGB).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 3 && depth == reflect.Uint16:
		m = new(imageRGB48).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 3 && depth == reflect.Int32:
		m = new(imageRGB96i).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 3 && depth == reflect.Float32:
		m = new(imageRGB96f).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 3 && depth == reflect.Int64:
		m = new(imageRGB192i).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 3 && depth == reflect.Float64:
		m = new(imageRGB192f).Init(p.Pix, p.Stride, p.Rect)
		return

	case channels == 4 && depth == reflect.Uint8:
		m = new(imageRGBA).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 4 && depth == reflect.Uint16:
		m = new(imageRGBA64).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 4 && depth == reflect.Int32:
		m = new(imageRGBA128i).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 4 && depth == reflect.Float32:
		m = new(imageRGBA128f).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 4 && depth == reflect.Int64:
		m = new(imageRGBA256i).Init(p.Pix, p.Stride, p.Rect)
		return
	case channels == 4 && depth == reflect.Float64:
		m = new(imageRGBA256f).Init(p.Pix, p.Stride, p.Rect)
		return

	default:
		err = fmt.Errorf("tiff: Image.asInternalImage, invalid format: channels = %v, depth = %v", channels, depth)
		return
	}
}
