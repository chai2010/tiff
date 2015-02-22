// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"fmt"
	"image"
	"io"
)

// DecodeConfig returns the color model and dimensions of a TIFF image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (cfg image.Config, err error) {
	rs := openSeekioReader(r, -1)
	defer rs.Close()

	hdr, err := ReadHeader(rs)
	if err != nil {
		return
	}
	ifd, err := ReadIFD(rs, hdr, hdr.Offset)
	if err != nil {
		return
	}
	cfg, err = ifd.ImageConfig()
	if err != nil {
		return
	}
	return
}

func DecodeAll(r io.Reader) (m []image.Image, err error) {
	return
}

// Decode reads a TIFF image from r and returns it as an image.Image.
// The type of Image returned depends on the contents of the TIFF.
func Decode(r io.Reader) (m image.Image, err error) {
	rs := openSeekioReader(r, -1)
	defer rs.Close()

	hdr, err := ReadHeader(rs)
	if err != nil {
		return
	}
	ifd, err := ReadIFD(rs, hdr, hdr.Offset)
	if err != nil {
		return
	}
	cfg, err := ifd.ImageConfig()
	if err != nil {
		return
	}

	imgRect := image.Rect(0, 0, cfg.Width, cfg.Height)
	imageType := ifd.ImageType()

	switch imageType {
	case ImageType_Gray, ImageType_GrayInvert:
		if ifd.Depth() == 16 {
			m = image.NewGray16(imgRect)
		} else {
			m = image.NewGray(imgRect)
		}
	case ImageType_Paletted:
		m = image.NewPaletted(imgRect, ifd.ColorMap())
	case ImageType_NRGBA:
		if ifd.Depth() == 16 {
			m = image.NewNRGBA64(imgRect)
		} else {
			m = image.NewNRGBA(imgRect)
		}
	case ImageType_RGB, ImageType_RGBA:
		if ifd.Depth() == 16 {
			m = image.NewRGBA64(imgRect)
		} else {
			m = image.NewRGBA(imgRect)
		}
	}
	if m == nil {
		err = fmt.Errorf("tiff: Decode, unknown format")
		return
	}

	blocksAcross := ifd.BlocksAcross()
	blocksDown := ifd.BlocksDown()

	for i := 0; i < blocksAcross; i++ {
		for j := 0; j < blocksDown; j++ {
			if err = ifd.DecodeBlock(rs, i, j, m); err != nil {
				return
			}
		}
	}
	return
}
