// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"fmt"
	"image"
)

func newImageWithIFD(r image.Rectangle, ifd *IFD) (m image.Image, err error) {
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
