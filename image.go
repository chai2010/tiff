// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"image"
	"image/color"
	"reflect"
	"runtime"
)

const (
	MemPMagic = "MemP" // See https://github.com/chai2010/image
)

const (
	isLittleEndian = (runtime.GOARCH == "386" ||
		runtime.GOARCH == "amd64" ||
		runtime.GOARCH == "arm" ||
		runtime.GOARCH == "arm64")
)

var (
	_ image.Image = (*Image)(nil)
)

// MemP Image Spec (Native Endian), see https://github.com/chai2010/image.
type Image struct {
	MemPMagic string // MemP
	Rect      image.Rectangle
	Channels  int
	DataType  reflect.Kind
	Pix       PixSilce

	// Stride is the Pix stride (in bytes, must align with SizeofKind(p.DataType))
	// between vertically adjacent pixels.
	Stride int
}

func NewImage(r image.Rectangle, channels int, dataType reflect.Kind) *Image {
	m := &Image{
		MemPMagic: MemPMagic,
		Rect:      r,
		Stride:    r.Dx() * channels * SizeofKind(dataType),
		Channels:  channels,
		DataType:  dataType,
	}
	m.Pix = make([]byte, r.Dy()*m.Stride)
	return m
}

func AsMemPImage(m interface{}) (p *Image, ok bool) {
	if p, ok := m.(*Image); ok {
		return p, true
	}

	switch m := m.(type) {
	case *image.Gray:
		return &Image{
			MemPMagic: MemPMagic,
			Rect:      m.Bounds(),
			Stride:    m.Stride,
			Channels:  1,
			DataType:  reflect.Uint8,
			Pix:       m.Pix,
		}, true
	case *image.RGBA:
		return &Image{
			MemPMagic: MemPMagic,
			Rect:      m.Bounds(),
			Stride:    m.Stride,
			Channels:  4,
			DataType:  reflect.Uint8,
			Pix:       m.Pix,
		}, true
	}

	p = new(Image)
	pType := reflect.TypeOf(*p)
	pValue := reflect.ValueOf(p)
	mValue := reflect.ValueOf(m)

	for pValue.Kind() == reflect.Ptr {
		pValue = pValue.Elem()
	}
	for mValue.Kind() == reflect.Ptr {
		mValue = mValue.Elem()
	}

	if mValue.Kind() != reflect.Struct {
		return nil, false
	}
	for i := 0; i < pType.NumField(); i++ {
		pField := pValue.Field(i)
		mField := mValue.FieldByName(pType.Field(i).Name)

		if mField.Kind() != pField.Kind() {
			return nil, false
		}
		pField.Set(mField)
	}

	if p.MemPMagic != MemPMagic {
		// ingore MemPMagic value
	}

	return p, true
}

func NewImageFrom(m image.Image) *Image {
	if p, ok := m.(*Image); ok {
		return p.Clone()
	}
	if p, ok := AsMemPImage(m); ok {
		return p.Clone()
	}

	switch m := m.(type) {
	case *image.Gray:
		b := m.Bounds()
		p := NewImage(b, 1, reflect.Uint8)

		for y := b.Min.Y; y < b.Max.Y; y++ {
			off0 := m.PixOffset(0, y)
			off1 := p.PixOffset(0, y)
			copy(p.Pix[off1:][:p.Stride], m.Pix[off0:][:m.Stride])
			off0 += m.Stride
			off1 += p.Stride
		}
		return p

	case *image.Gray16:
		b := m.Bounds()
		p := NewImage(b, 1, reflect.Uint16)

		for y := b.Min.Y; y < b.Max.Y; y++ {
			off0 := m.PixOffset(0, y)
			off1 := p.PixOffset(0, y)
			copy(p.Pix[off1:][:p.Stride], m.Pix[off0:][:m.Stride])
			off0 += m.Stride
			off1 += p.Stride
		}
		if isLittleEndian {
			p.Pix.SwapEndian(p.DataType)
		}
		return p

	case *image.RGBA:
		b := m.Bounds()
		p := NewImage(b, 4, reflect.Uint8)

		for y := b.Min.Y; y < b.Max.Y; y++ {
			off0 := m.PixOffset(0, y)
			off1 := p.PixOffset(0, y)
			copy(p.Pix[off1:][:p.Stride], m.Pix[off0:][:m.Stride])
			off0 += m.Stride
			off1 += p.Stride
		}
		return p

	case *image.RGBA64:
		b := m.Bounds()
		p := NewImage(b, 4, reflect.Uint16)

		for y := b.Min.Y; y < b.Max.Y; y++ {
			off0 := m.PixOffset(0, y)
			off1 := p.PixOffset(0, y)
			copy(p.Pix[off1:][:p.Stride], m.Pix[off0:][:m.Stride])
			off0 += m.Stride
			off1 += p.Stride
		}
		if isLittleEndian {
			p.Pix.SwapEndian(p.DataType)
		}
		return p

	case *image.YCbCr:
		b := m.Bounds()
		p := NewImage(b, 4, reflect.Uint8)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				R, G, B, A := m.At(x, y).RGBA()

				i := p.PixOffset(x, y)
				p.Pix[i+0] = uint8(R >> 8)
				p.Pix[i+1] = uint8(G >> 8)
				p.Pix[i+2] = uint8(B >> 8)
				p.Pix[i+3] = uint8(A >> 8)
			}
		}
		return p

	default:
		b := m.Bounds()
		p := NewImage(b, 4, reflect.Uint16)
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				R, G, B, A := m.At(x, y).RGBA()

				i := p.PixOffset(x, y)
				p.Pix[i+0] = uint8(R >> 8)
				p.Pix[i+1] = uint8(R)
				p.Pix[i+2] = uint8(G >> 8)
				p.Pix[i+3] = uint8(G)
				p.Pix[i+4] = uint8(B >> 8)
				p.Pix[i+5] = uint8(B)
				p.Pix[i+6] = uint8(A >> 8)
				p.Pix[i+7] = uint8(A)
			}
		}
		return p
	}
}

func (p *Image) Clone() *Image {
	q := new(Image)
	*q = *p
	q.Pix = append([]byte(nil), p.Pix...)
	return q
}

func (p *Image) Bounds() image.Rectangle {
	return p.Rect
}

func (p *Image) ColorModel() color.Model {
	return ColorModel(p.Channels, p.DataType)
}

func (p *Image) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return Color{
			Channels: p.Channels,
			DataType: p.DataType,
		}
	}
	i := p.PixOffset(x, y)
	n := SizeofPixel(p.Channels, p.DataType)
	return Color{
		Channels: p.Channels,
		DataType: p.DataType,
		Pix:      p.Pix[i:][:n],
	}
}

func (p *Image) PixelAt(x, y int) []byte {
	if !(image.Point{x, y}.In(p.Rect)) {
		return nil
	}
	i := p.PixOffset(x, y)
	n := SizeofPixel(p.Channels, p.DataType)
	return p.Pix[i:][:n]
}

func (p *Image) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	n := SizeofPixel(p.Channels, p.DataType)
	v := p.ColorModel().Convert(c).(Color)
	copy(p.Pix[i:][:n], v.Pix)
}

func (p *Image) SetPixel(x, y int, c []byte) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	n := SizeofPixel(p.Channels, p.DataType)
	copy(p.Pix[i:][:n], c)
}

func (p *Image) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*SizeofPixel(p.Channels, p.DataType)
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
		Pix:      p.Pix[i:],
		Stride:   p.Stride,
		Rect:     r,
		Channels: p.Channels,
		DataType: p.DataType,
	}
}

func (p *Image) StdImage() image.Image {
	switch {
	case p.Channels == 1 && p.DataType == reflect.Uint8:
		return &image.Gray{
			Pix:    p.Pix,
			Stride: p.Stride,
			Rect:   p.Rect,
		}
	case p.Channels == 1 && p.DataType == reflect.Uint16:
		m := &image.Gray16{
			Pix:    p.Pix,
			Stride: p.Stride,
			Rect:   p.Rect,
		}
		if isLittleEndian {
			m.Pix = append([]byte(nil), m.Pix...)
			PixSilce(m.Pix).SwapEndian(p.DataType)
		}
		return m
	case p.Channels == 4 && p.DataType == reflect.Uint8:
		return &image.RGBA{
			Pix:    p.Pix,
			Stride: p.Stride,
			Rect:   p.Rect,
		}
	case p.Channels == 4 && p.DataType == reflect.Uint16:
		m := &image.RGBA64{
			Pix:    p.Pix,
			Stride: p.Stride,
			Rect:   p.Rect,
		}
		if isLittleEndian {
			m.Pix = append([]byte(nil), m.Pix...)
			PixSilce(m.Pix).SwapEndian(p.DataType)
		}
		return m
	}

	return p
}
