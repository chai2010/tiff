// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"image/color"
)

var (
	_ color.Color = (*colorGray)(nil)
	_ color.Color = (*colorGray16)(nil)
	_ color.Color = (*colorGray32i)(nil)
	_ color.Color = (*colorGray32f)(nil)
	_ color.Color = (*colorGray64i)(nil)
	_ color.Color = (*colorGray64f)(nil)
	_ color.Color = (*colorGrayA)(nil)
	_ color.Color = (*colorGrayA32)(nil)
	_ color.Color = (*colorGrayA64i)(nil)
	_ color.Color = (*colorGrayA64f)(nil)
	_ color.Color = (*colorGrayA128i)(nil)
	_ color.Color = (*colorGrayA128f)(nil)
	_ color.Color = (*colorRGB)(nil)
	_ color.Color = (*colorRGB48)(nil)
	_ color.Color = (*colorRGB96i)(nil)
	_ color.Color = (*colorRGB96f)(nil)
	_ color.Color = (*colorRGB192i)(nil)
	_ color.Color = (*colorRGB192f)(nil)
	_ color.Color = (*colorRGBA)(nil)
	_ color.Color = (*colorRGBA64)(nil)
	_ color.Color = (*colorRGBA128i)(nil)
	_ color.Color = (*colorRGBA128f)(nil)
	_ color.Color = (*colorRGBA256i)(nil)
	_ color.Color = (*colorRGBA256f)(nil)
)

// Models for the standard color types.
var (
	colorGrayModel      color.Model = color.ModelFunc(colorToGrayModel)
	colorGray16Model    color.Model = color.ModelFunc(colorToGray16Model)
	colorGray32iModel   color.Model = color.ModelFunc(colorToGray32iModel)
	colorGray32fModel   color.Model = color.ModelFunc(colorToGray32fModel)
	colorGray64iModel   color.Model = color.ModelFunc(colorToGray64iModel)
	colorGray64fModel   color.Model = color.ModelFunc(colorToGray64fModel)
	colorGrayAModel     color.Model = color.ModelFunc(colorToGrayAModel)
	colorGrayA32Model   color.Model = color.ModelFunc(colorToGrayA32Model)
	colorGrayA64iModel  color.Model = color.ModelFunc(colorToGrayA64iModel)
	colorGrayA64fModel  color.Model = color.ModelFunc(colorToGrayA64fModel)
	colorGrayA128iModel color.Model = color.ModelFunc(colorToGrayA128iModel)
	colorGrayA128fModel color.Model = color.ModelFunc(colorToGrayA128fModel)
	colorRGBModel       color.Model = color.ModelFunc(colorToRgbModel)
	colorRGB48Model     color.Model = color.ModelFunc(colorToRgb48Model)
	colorRGB96iModel    color.Model = color.ModelFunc(colorToRgb96iModel)
	colorRGB96fModel    color.Model = color.ModelFunc(colorToRgb96fModel)
	colorRGB192iModel   color.Model = color.ModelFunc(colorToRgb192iModel)
	colorRGB192fModel   color.Model = color.ModelFunc(colorToRgb192fModel)
	colorRGBAModel      color.Model = color.ModelFunc(colorToRgbaModel)
	colorRGBA64Model    color.Model = color.ModelFunc(colorToRgba64Model)
	colorRGBA128iModel  color.Model = color.ModelFunc(colorToRgba128iModel)
	colorRGBA128fModel  color.Model = color.ModelFunc(colorToRgba128fModel)
	colorRGBA256iModel  color.Model = color.ModelFunc(colorToRgba256iModel)
	colorRGBA256fModel  color.Model = color.ModelFunc(colorToRgba256fModel)
)

func colorRgbToGray(r, g, b uint32) uint32 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}

func colorRgbToGrayI32(r, g, b int32) int32 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}

func colorRgbToGrayF32(r, g, b float32) float32 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}

func colorRgbToGrayI64(r, g, b int64) int64 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}

func colorRgbToGrayF64(r, g, b float64) float64 {
	y := (299*r + 587*g + 114*b + 500) / 1000
	return y
}
