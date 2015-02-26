// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"fmt"
	"image/color"
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

func (p *Image) ColorAt(x, y int) Color {
	return Color{}
}

func (p *Image) GrayAt(x, y int) color.Gray {
	return color.Gray{}
}
func (p *Image) Gray16At(x, y int) color.Gray16 {
	return color.Gray16{}
}
func (p *Image) Gray32fAt(x, y int) float32 {
	return 0
}

func (p *Image) RGBAt(x, y int) [3]uint8 {
	return [3]uint8{}
}
func (p *Image) RGB48At(x, y int) [3]uint16 {
	return [3]uint16{}
}
func (p *Image) RGB96fAt(x, y int) [3]float32 {
	return [3]float32{}
}

func (p *Image) RGBAAt(x, y int) color.RGBA {
	return color.RGBA{}
}
func (p *Image) RGBA64At(x, y int) color.RGBA64 {
	return color.RGBA64{}
}
func (p *Image) RGBA128fAt(x, y int) [4]float32 {
	return [4]float32{}
}

func (p *Image) SetColor(x, y int, c Color) {
	//
}

func (p *Image) SetGray(x, y int, c color.Gray) {
	//
}
func (p *Image) SetGray16(x, y int, c color.Gray16) {
	//
}
func (p *Image) SetGray32f(x, y int, c float32) {
	//
}

func (p *Image) SetRGB(x, y int, c [3]uint8) {
	//
}
func (p *Image) SetRGB48(x, y int, c [3]uint16) {
	//
}
func (p *Image) SetRGB96f(x, y int, c [3]float32) {
	//
}

func (p *Image) SetRGBA(x, y int, c color.RGBA) {
	//
}
func (p *Image) SetRGBA64(x, y int, c color.RGBA64) {
	//
}
func (p *Image) SetRGBA128f(x, y int, c [4]float32) {
	//
}
