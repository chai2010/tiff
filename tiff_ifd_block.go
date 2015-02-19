// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"fmt"
	"image"
	"io"
)

func (p *IFD) ReadBlock(r io.ReadSeeker, offset, length int64, dst *Image, rect image.Rectangle) (err error) {
	var data []byte
	if _, err = r.Seek(offset, 0); err != nil {
		return
	}
	limitReader := io.LimitReader(r, length)
	if data, err = p.Compression().ReadAll(limitReader); err != nil {
		return
	}

	if p.Predictor() == TagValue_PredictorType_Horizontal {
		if data, err = p.decodePredictor(rect, data); err != nil {
			return
		}
	}

	return
}

func (p *IFD) decodePredictor(r image.Rectangle, data []byte) (out []byte, err error) {
	bpp := p.BitsPerSample()
	spp := p.SamplesPerPixel()

	switch bpp {
	case 16:
		var off int
		for y := r.Min.Y; y < r.Max.Y; y++ {
			off += spp * 2
			for x := 0; x < r.Dx()*spp*2; x += 2 {
				v0 := p.Header.ByteOrder.Uint16(data[off-spp*2 : off-spp*2+2])
				v1 := p.Header.ByteOrder.Uint16(data[off : off+2])
				p.Header.ByteOrder.PutUint16(data[off:off+2], v1+v0)
				off += 2
			}
		}
	case 8:
		var off int
		for y := r.Min.Y; y < r.Max.Y; y++ {
			off += spp
			for x := 0; x < r.Dx()*spp; x++ {
				data[off] += data[off-spp]
				off++
			}
		}
	default:
		err = fmt.Errorf("tiff: IFD.decodePredictor, bad BitsPerSample = %d", bpp)
		return
	}
	return
}
