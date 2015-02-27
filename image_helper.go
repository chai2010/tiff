// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"encoding/binary"
	"fmt"
	"math"
	"reflect"
)

func getPixelSize(channels int, dataType reflect.Kind) int {
	switch dataType {
	case reflect.Uint8:
		return 1 * channels
	case reflect.Uint16:
		return 2 * channels
	case reflect.Uint32, reflect.Int32, reflect.Float32:
		return 4 * channels
	case reflect.Uint64, reflect.Int64, reflect.Float64:
		return 8 * channels
	default:
		panic(fmt.Errorf("tiff: getPixelSize, bad dataType %v", dataType))
	}
}

func pGrayAt(pix []byte) colorGray {
	return colorGray{
		Y: pix[1*0],
	}
}
func pSetGray(pix []byte, c colorGray) {
	pix[1*0] = c.Y
}

func pGray16At(pix []byte) colorGray16 {
	return colorGray16{
		Y: binary.BigEndian.Uint16(pix[2*0:]),
	}
}
func pSetGray16(pix []byte, c colorGray16) {
	binary.BigEndian.PutUint16(pix[2*0:], c.Y)
}

func pGray32iAt(pix []byte) colorGray32i {
	return colorGray32i{
		Y: int32(binary.BigEndian.Uint32(pix[4*0:])),
	}
}
func pSetGray32i(pix []byte, c colorGray32i) {
	binary.BigEndian.PutUint32(pix[4*0:], uint32(c.Y))
}

func pGray32fAt(pix []byte) colorGray32f {
	return colorGray32f{
		Y: math.Float32frombits(binary.BigEndian.Uint32(pix[4*0:])),
	}
}
func pSetGray32f(pix []byte, c colorGray32f) {
	binary.BigEndian.PutUint32(pix[4*0:], math.Float32bits(c.Y))
}

func pGray64iAt(pix []byte) colorGray64i {
	return colorGray64i{
		Y: int64(binary.BigEndian.Uint64(pix[8*0:])),
	}
}
func pSetGray64i(pix []byte, c colorGray64i) {
	binary.BigEndian.PutUint64(pix[8*0:], uint64(c.Y))
}

func pGray64fAt(pix []byte) colorGray64f {
	return colorGray64f{
		Y: math.Float64frombits(binary.BigEndian.Uint64(pix[8*0:])),
	}
}
func pSetGray64f(pix []byte, c colorGray64f) {
	binary.BigEndian.PutUint64(pix[8*0:], math.Float64bits(c.Y))
}

func pGrayAAt(pix []byte) colorGrayA {
	return colorGrayA{
		Y: pix[1*0],
		A: pix[1*1],
	}
}
func pSetGrayA(pix []byte, c colorGrayA) {
	pix[1*0] = c.Y
	pix[1*1] = c.A
}

func pGrayA32At(pix []byte) colorGrayA32 {
	return colorGrayA32{
		Y: binary.BigEndian.Uint16(pix[2*0:]),
		A: binary.BigEndian.Uint16(pix[2*1:]),
	}
}
func pSetGrayA32(pix []byte, c colorGrayA32) {
	binary.BigEndian.PutUint16(pix[2*0:], c.Y)
	binary.BigEndian.PutUint16(pix[2*1:], c.A)
}

func pGrayA64iAt(pix []byte) colorGrayA64i {
	return colorGrayA64i{
		Y: int32(binary.BigEndian.Uint32(pix[4*0:])),
		A: int32(binary.BigEndian.Uint32(pix[4*1:])),
	}
}
func pSetGrayA64i(pix []byte, c colorGrayA64i) {
	binary.BigEndian.PutUint32(pix[4*0:], uint32(c.Y))
	binary.BigEndian.PutUint32(pix[4*1:], uint32(c.A))
}

func pGrayA64fAt(pix []byte) colorGrayA64f {
	return colorGrayA64f{
		Y: math.Float32frombits(binary.BigEndian.Uint32(pix[4*0:])),
		A: math.Float32frombits(binary.BigEndian.Uint32(pix[4*1:])),
	}
}
func pSetGrayA64f(pix []byte, c colorGrayA64f) {
	binary.BigEndian.PutUint32(pix[4*0:], math.Float32bits(c.Y))
	binary.BigEndian.PutUint32(pix[4*1:], math.Float32bits(c.A))
}

func pGrayA128iAt(pix []byte) colorGrayA128i {
	return colorGrayA128i{
		Y: int64(binary.BigEndian.Uint64(pix[8*0:])),
		A: int64(binary.BigEndian.Uint64(pix[8*1:])),
	}
}
func pSetGrayA128i(pix []byte, c colorGrayA128i) {
	binary.BigEndian.PutUint64(pix[8*0:], uint64(c.Y))
	binary.BigEndian.PutUint64(pix[8*1:], uint64(c.A))
}

func pGrayA128fAt(pix []byte) colorGrayA128f {
	return colorGrayA128f{
		Y: math.Float64frombits(binary.BigEndian.Uint64(pix[8*0:])),
		A: math.Float64frombits(binary.BigEndian.Uint64(pix[8*1:])),
	}
}
func pSetGrayA128f(pix []byte, c colorGrayA128f) {
	binary.BigEndian.PutUint64(pix[8*0:], math.Float64bits(c.Y))
	binary.BigEndian.PutUint64(pix[8*1:], math.Float64bits(c.A))
}

func pRGBAt(pix []byte) colorRGB {
	return colorRGB{
		R: pix[1*0],
		G: pix[1*1],
		B: pix[1*2],
	}
}
func pSetRGB(pix []byte, c colorRGB) {
	pix[1*0] = c.R
	pix[1*1] = c.G
	pix[1*2] = c.B
}

func pRGB48At(pix []byte) colorRGB48 {
	return colorRGB48{
		R: binary.BigEndian.Uint16(pix[2*0:]),
		G: binary.BigEndian.Uint16(pix[2*1:]),
		B: binary.BigEndian.Uint16(pix[2*2:]),
	}
}
func pSetRGB48(pix []byte, c colorRGB48) {
	binary.BigEndian.PutUint16(pix[2*0:], c.R)
	binary.BigEndian.PutUint16(pix[2*1:], c.G)
	binary.BigEndian.PutUint16(pix[2*2:], c.B)
}

func pRGB96iAt(pix []byte) colorRGB96i {
	return colorRGB96i{
		R: int32(binary.BigEndian.Uint32(pix[4*0:])),
		G: int32(binary.BigEndian.Uint32(pix[4*1:])),
		B: int32(binary.BigEndian.Uint32(pix[4*2:])),
	}
}
func pSetRGB96i(pix []byte, c colorRGB96i) {
	binary.BigEndian.PutUint32(pix[4*0:], uint32(c.R))
	binary.BigEndian.PutUint32(pix[4*1:], uint32(c.G))
	binary.BigEndian.PutUint32(pix[4*2:], uint32(c.B))
}

func pRGB96fAt(pix []byte) colorRGB96f {
	return colorRGB96f{
		R: math.Float32frombits(binary.BigEndian.Uint32(pix[4*0:])),
		G: math.Float32frombits(binary.BigEndian.Uint32(pix[4*1:])),
		B: math.Float32frombits(binary.BigEndian.Uint32(pix[4*2:])),
	}
}
func pSetRGB96f(pix []byte, c colorRGB96f) {
	binary.BigEndian.PutUint32(pix[4*0:], math.Float32bits(c.R))
	binary.BigEndian.PutUint32(pix[4*1:], math.Float32bits(c.G))
	binary.BigEndian.PutUint32(pix[4*2:], math.Float32bits(c.B))
}

func pRGB192iAt(pix []byte) colorRGB192i {
	return colorRGB192i{
		R: int64(binary.BigEndian.Uint64(pix[8*0:])),
		G: int64(binary.BigEndian.Uint64(pix[8*1:])),
		B: int64(binary.BigEndian.Uint64(pix[8*2:])),
	}
}
func pSetRGB192i(pix []byte, c colorRGB192i) {
	binary.BigEndian.PutUint64(pix[8*0:], uint64(c.R))
	binary.BigEndian.PutUint64(pix[8*1:], uint64(c.G))
	binary.BigEndian.PutUint64(pix[8*2:], uint64(c.B))
}

func pRGB192fAt(pix []byte) colorRGB192f {
	return colorRGB192f{
		R: math.Float64frombits(binary.BigEndian.Uint64(pix[8*0:])),
		G: math.Float64frombits(binary.BigEndian.Uint64(pix[8*1:])),
		B: math.Float64frombits(binary.BigEndian.Uint64(pix[8*2:])),
	}
}
func pSetRGB192f(pix []byte, c colorRGB192f) {
	binary.BigEndian.PutUint64(pix[8*0:], math.Float64bits(c.R))
	binary.BigEndian.PutUint64(pix[8*1:], math.Float64bits(c.G))
	binary.BigEndian.PutUint64(pix[8*2:], math.Float64bits(c.B))
}

func pRGBAAt(pix []byte) colorRGBA {
	return colorRGBA{
		R: pix[1*0],
		G: pix[1*1],
		B: pix[1*2],
		A: pix[1*3],
	}
}
func pSetRGBA(pix []byte, c colorRGBA) {
	pix[1*0] = c.R
	pix[1*1] = c.G
	pix[1*2] = c.B
	pix[1*3] = c.A
}

func pRGBA64At(pix []byte) colorRGBA64 {
	return colorRGBA64{
		R: binary.BigEndian.Uint16(pix[2*0:]),
		G: binary.BigEndian.Uint16(pix[2*1:]),
		B: binary.BigEndian.Uint16(pix[2*2:]),
		A: binary.BigEndian.Uint16(pix[2*3:]),
	}
}
func pSetRGBA64(pix []byte, c colorRGBA64) {
	binary.BigEndian.PutUint16(pix[2*0:], c.R)
	binary.BigEndian.PutUint16(pix[2*1:], c.G)
	binary.BigEndian.PutUint16(pix[2*2:], c.B)
	binary.BigEndian.PutUint16(pix[2*3:], c.A)
}

func pRGBA128iAt(pix []byte) colorRGBA128i {
	return colorRGBA128i{
		R: int32(binary.BigEndian.Uint32(pix[4*0:])),
		G: int32(binary.BigEndian.Uint32(pix[4*1:])),
		B: int32(binary.BigEndian.Uint32(pix[4*2:])),
		A: int32(binary.BigEndian.Uint32(pix[4*3:])),
	}
}
func pSetRGBA128i(pix []byte, c colorRGBA128i) {
	binary.BigEndian.PutUint32(pix[4*0:], uint32(c.R))
	binary.BigEndian.PutUint32(pix[4*1:], uint32(c.G))
	binary.BigEndian.PutUint32(pix[4*2:], uint32(c.B))
	binary.BigEndian.PutUint32(pix[4*3:], uint32(c.A))
}

func pRGBA128fAt(pix []byte) colorRGBA128f {
	return colorRGBA128f{
		R: math.Float32frombits(binary.BigEndian.Uint32(pix[4*0:])),
		G: math.Float32frombits(binary.BigEndian.Uint32(pix[4*1:])),
		B: math.Float32frombits(binary.BigEndian.Uint32(pix[4*2:])),
		A: math.Float32frombits(binary.BigEndian.Uint32(pix[4*3:])),
	}
}
func pSetRGBA128f(pix []byte, c colorRGBA128f) {
	binary.BigEndian.PutUint32(pix[4*0:], math.Float32bits(c.R))
	binary.BigEndian.PutUint32(pix[4*1:], math.Float32bits(c.G))
	binary.BigEndian.PutUint32(pix[4*2:], math.Float32bits(c.B))
	binary.BigEndian.PutUint32(pix[4*3:], math.Float32bits(c.A))
}

func pRGBA256iAt(pix []byte) colorRGBA256i {
	return colorRGBA256i{
		R: int64(binary.BigEndian.Uint64(pix[8*0:])),
		G: int64(binary.BigEndian.Uint64(pix[8*1:])),
		B: int64(binary.BigEndian.Uint64(pix[8*2:])),
		A: int64(binary.BigEndian.Uint64(pix[8*3:])),
	}
}
func pSetRGBA256i(pix []byte, c colorRGBA256i) {
	binary.BigEndian.PutUint64(pix[8*0:], uint64(c.R))
	binary.BigEndian.PutUint64(pix[8*1:], uint64(c.G))
	binary.BigEndian.PutUint64(pix[8*2:], uint64(c.B))
	binary.BigEndian.PutUint64(pix[8*3:], uint64(c.A))
}

func pRGBA256fAt(pix []byte) colorRGBA256f {
	return colorRGBA256f{
		R: math.Float64frombits(binary.BigEndian.Uint64(pix[8*0:])),
		G: math.Float64frombits(binary.BigEndian.Uint64(pix[8*1:])),
		B: math.Float64frombits(binary.BigEndian.Uint64(pix[8*2:])),
		A: math.Float64frombits(binary.BigEndian.Uint64(pix[8*3:])),
	}
}
func pSetRGBA256f(pix []byte, c colorRGBA256f) {
	binary.BigEndian.PutUint64(pix[8*0:], math.Float64bits(c.R))
	binary.BigEndian.PutUint64(pix[8*1:], math.Float64bits(c.G))
	binary.BigEndian.PutUint64(pix[8*2:], math.Float64bits(c.B))
	binary.BigEndian.PutUint64(pix[8*3:], math.Float64bits(c.A))
}
