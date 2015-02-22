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
	Rect     image.Rectangle
	Channels int
	Kind     reflect.Kind // bool/uint*/int*/float*/Complex*/string
	Pix      interface{}  // At(x,y): (Pix.([]T))[y*Stride+x:][:Channels]
	Stride   int
}

func NewImage(r image.Rectangle, channels int, kind reflect.Kind) *Image {
	var pix interface{}
	var stride = r.Dx() * channels
	switch kind {
	case reflect.Bool:
		pix = make([]bool, stride*r.Dy())
	case reflect.Int:
		pix = make([]int, stride*r.Dy())
	case reflect.Int8:
		pix = make([]int8, stride*r.Dy())
	case reflect.Int16:
		pix = make([]int16, stride*r.Dy())
	case reflect.Int32:
		pix = make([]int32, stride*r.Dy())
	case reflect.Int64:
		pix = make([]int64, stride*r.Dy())
	case reflect.Uint:
		pix = make([]uint, stride*r.Dy())
	case reflect.Uint8:
		pix = make([]uint8, stride*r.Dy())
	case reflect.Uint16:
		pix = make([]uint16, stride*r.Dy())
	case reflect.Uint32:
		pix = make([]uint32, stride*r.Dy())
	case reflect.Uint64:
		pix = make([]uint64, stride*r.Dy())
	case reflect.Complex64:
		pix = make([]complex64, stride*r.Dy())
	case reflect.Complex128:
		pix = make([]complex128, stride*r.Dy())
	case reflect.String:
		pix = make([]string, stride*r.Dy())
	default:
		panic("tiff: NewImage, unknown kind = " + kind.String())
	}
	return &Image{
		Rect:     r,
		Channels: channels,
		Kind:     kind,
		Pix:      pix,
		Stride:   r.Dx() * channels,
	}
}

func (p *Image) Bounds() image.Rectangle {
	return p.Rect
}

func (p *Image) ColorModel() color.Model {
	return nil
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

func (p *Image) Slice(low_high_max ...int) interface{} {
	return nil
}

func (p *Image) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &Image{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &Image{
		Rect:     r,
		Channels: p.Channels,
		Kind:     p.Kind,
		Pix:      p.Slice(i),
		Stride:   p.Stride,
	}
}

func (p *Image) StdImage() image.Image {
	return nil
}

func (p *Image) Opaque() bool {
	return true
}
