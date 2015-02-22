// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"fmt"
	"image"
	"image/color"
	"io"
)

func (p *IFD) BlocksAcross() int {
	if _, ok := p.TagGetter().GetTileWidth(); ok {
		imageWidth, _ := p.TagGetter().GetImageWidth()
		blockWidth, _ := p.TagGetter().GetTileWidth()
		return int((imageWidth + blockWidth - 1) / blockWidth)
	} else {
		return 1
	}
	return 0
}

func (p *IFD) BlocksDown() int {
	if _, ok := p.TagGetter().GetTileLength(); ok {
		imageHeight, _ := p.TagGetter().GetImageLength()
		blockHeight, _ := p.TagGetter().GetTileLength()
		return int((imageHeight + blockHeight - 1) / blockHeight)
	} else {
		imageHeight, _ := p.TagGetter().GetImageLength()
		blockHeight, ok := p.TagGetter().GetRowsPerStrip()
		if !ok || blockHeight == 0 || blockHeight > imageHeight {
			blockHeight = imageHeight
		}
		return int((imageHeight + blockHeight - 1) / blockHeight)
	}
}

func (p *IFD) BlockBounds(col, row int) image.Rectangle {
	blocksAcross, blocksDown := p.BlocksAcross(), p.BlocksDown()
	if col < 0 || row < 0 || col >= blocksAcross || row >= blocksDown {
		return image.Rectangle{}
	}

	if _, ok := p.TagGetter().GetTileWidth(); ok {
		blockWidth, _ := p.TagGetter().GetTileWidth()
		blockHeight, _ := p.TagGetter().GetTileLength()

		blkW := blockWidth
		blkH := blockHeight

		xmin := col * int(blockWidth)
		ymin := row * int(blockHeight)
		xmax := xmin + int(blkW)
		ymax := ymin + int(blkH)

		return image.Rect(xmin, ymin, xmax, ymax)

	} else {
		imageWidth, _ := p.TagGetter().GetImageWidth()
		imageHeight, _ := p.TagGetter().GetImageLength()

		blockWidth := imageWidth
		blockHeight, ok := p.TagGetter().GetRowsPerStrip()
		if !ok || blockHeight == 0 {
			blockHeight = imageHeight
		}

		blkW := blockWidth
		blkH := blockHeight
		if row == blocksDown-1 && imageHeight%blockHeight != 0 {
			blkH = imageHeight % blockHeight
		}

		xmin := col * int(blockWidth)
		ymin := row * int(blockHeight)
		xmax := xmin + int(blkW)
		ymax := ymin + int(blkH)

		return image.Rect(xmin, ymin, xmax, ymax)
	}
}

func (p *IFD) BlockOffset(col, row int) int64 {
	blocksAcross, blocksDown := p.BlocksAcross(), p.BlocksDown()
	if col < 0 || row < 0 || col >= blocksAcross || row >= blocksDown {
		return 0
	}
	if _, ok := p.TagGetter().GetTileWidth(); ok {
		offsets, ok := p.TagGetter().GetTileOffsets()
		if !ok || len(offsets) != blocksAcross*blocksDown {
			return 0
		}
		return offsets[row*blocksAcross+col]
	} else {
		offsets, ok := p.TagGetter().GetStripOffsets()
		if !ok || len(offsets) != blocksAcross*blocksDown {
			return 0
		}
		return offsets[row*blocksAcross+col]
	}
}

func (p *IFD) BlockCount(col, row int) int64 {
	blocksAcross, blocksDown := p.BlocksAcross(), p.BlocksDown()
	if col < 0 || row < 0 || col >= blocksAcross || row >= blocksDown {
		return 0
	}
	if _, ok := p.TagGetter().GetTileWidth(); ok {
		counts, ok := p.TagGetter().GetTileByteCounts()
		if !ok || len(counts) != blocksAcross*blocksDown {
			return 0
		}
		return counts[row*blocksAcross+col]
	} else {
		counts, ok := p.TagGetter().GetStripByteCounts()
		if !ok || len(counts) != blocksAcross*blocksDown {
			return 0
		}
		return counts[row*blocksAcross+col]
	}
}

func (p *IFD) DecodeBlock(r io.ReadSeeker, col, row int, dst image.Image) (err error) {
	blocksAcross, blocksDown := p.BlocksAcross(), p.BlocksDown()
	if col < 0 || row < 0 || col >= blocksAcross || row >= blocksDown {
		err = fmt.Errorf("tiff: IFD.DecodeBlock, bad col/row = %d/%d", col, row)
		return
	}

	offset := p.BlockOffset(col, row)
	count := p.BlockCount(col, row)

	if _, err = r.Seek(offset, 0); err != nil {
		return
	}
	limitReader := io.LimitReader(r, count)

	var data []byte
	if data, err = p.Compression().ReadAll(limitReader); err != nil {
		return
	}

	rect := p.BlockBounds(col, row)
	predictor, ok := p.TagGetter().GetPredictor()
	if ok && predictor == TagValue_PredictorType_Horizontal {
		if data, err = p.decodePredictor(data, rect); err != nil {
			return
		}
	}

	err = p.decodeBlock(data, dst, rect)
	return
}

func (p *IFD) decodePredictor(data []byte, r image.Rectangle) (out []byte, err error) {
	bpp := p.Depth()
	spp := p.Channels()

	switch bpp {
	case 16:
		var off int
		for y := r.Min.Y; y < r.Max.Y; y++ {
			off += spp * 2
			for x := 0; x < (r.Dx()-1)*spp*2; x += 2 {
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
			for x := 0; x < (r.Dx()-1)*spp; x++ {
				data[off] += data[off-spp]
				off++
			}
		}
	default:
		err = fmt.Errorf("tiff: IFD.decodePredictor, bad BitsPerSample = %d", bpp)
		return
	}
	out = data
	return
}

func (p *IFD) decodeBlock(buf []byte, dst image.Image, r image.Rectangle) (err error) {
	xmin, ymin := r.Min.X, r.Min.Y
	xmax, ymax := r.Max.X, r.Max.Y

	rMaxX := minInt(xmax, dst.Bounds().Max.X)
	rMaxY := minInt(ymax, dst.Bounds().Max.Y)

	switch p.ImageType() {
	case ImageType_Gray, ImageType_GrayInvert:
		if p.Depth() == 16 {
			var off int
			img := dst.(*image.Gray16)
			for y := ymin; y < rMaxY; y++ {
				for x := xmin; x < rMaxX; x++ {
					v := p.Header.ByteOrder.Uint16(buf[off : off+2])
					off += 2
					if p.ImageType() == ImageType_GrayInvert {
						v = 0xffff - v
					}
					img.SetGray16(x, y, color.Gray16{v})
				}
			}
		} else {
			bitReader := newBitsReader(buf)
			img := dst.(*image.Gray)
			max := uint32((1 << uint(p.Depth())) - 1)
			for y := ymin; y < rMaxY; y++ {
				for x := xmin; x < rMaxX; x++ {
					v := uint8(bitReader.ReadBits(uint(p.Depth())) * 0xff / max)
					if p.ImageType() == ImageType_GrayInvert {
						v = 0xff - v
					}
					img.SetGray(x, y, color.Gray{v})
				}
			}
		}
	case ImageType_Paletted:
		bitReader := newBitsReader(buf)
		img := dst.(*image.Paletted)
		for y := ymin; y < rMaxY; y++ {
			for x := xmin; x < rMaxX; x++ {
				img.SetColorIndex(x, y, uint8(bitReader.ReadBits(uint(p.Depth()))))
			}
			bitReader.flushBits()
		}
	case ImageType_RGB:
		if p.Depth() == 16 {
			var off int
			img := dst.(*image.RGBA64)
			for y := ymin; y < rMaxY; y++ {
				for x := xmin; x < rMaxX; x++ {
					r := p.Header.ByteOrder.Uint16(buf[off+0 : off+2])
					g := p.Header.ByteOrder.Uint16(buf[off+2 : off+4])
					b := p.Header.ByteOrder.Uint16(buf[off+4 : off+6])
					off += 6
					img.SetRGBA64(x, y, color.RGBA64{r, g, b, 0xffff})
				}
			}
		} else {
			img := dst.(*image.RGBA)
			for y := ymin; y < rMaxY; y++ {
				min := img.PixOffset(xmin, y)
				max := img.PixOffset(rMaxX, y)
				off := (y - ymin) * (xmax - xmin) * 3
				for i := min; i < max; i += 4 {
					img.Pix[i+0] = buf[off+0]
					img.Pix[i+1] = buf[off+1]
					img.Pix[i+2] = buf[off+2]
					img.Pix[i+3] = 0xff
					off += 3
				}
			}
		}
	case ImageType_NRGBA:
		if p.Depth() == 16 {
			var off int
			img := dst.(*image.NRGBA64)
			for y := ymin; y < rMaxY; y++ {
				for x := xmin; x < rMaxX; x++ {
					r := p.Header.ByteOrder.Uint16(buf[off+0 : off+2])
					g := p.Header.ByteOrder.Uint16(buf[off+2 : off+4])
					b := p.Header.ByteOrder.Uint16(buf[off+4 : off+6])
					a := p.Header.ByteOrder.Uint16(buf[off+6 : off+8])
					off += 8
					img.SetNRGBA64(x, y, color.NRGBA64{r, g, b, a})
				}
			}
		} else {
			img := dst.(*image.NRGBA)
			for y := ymin; y < rMaxY; y++ {
				min := img.PixOffset(xmin, y)
				max := img.PixOffset(rMaxX, y)
				copy(img.Pix[min:max], buf[(y-ymin)*(xmax-xmin)*4:(y-ymin+1)*(xmax-xmin)*4])
			}
		}
	case ImageType_RGBA:
		if p.Depth() == 16 {
			var off int
			img := dst.(*image.RGBA64)
			for y := ymin; y < rMaxY; y++ {
				for x := xmin; x < rMaxX; x++ {
					r := p.Header.ByteOrder.Uint16(buf[off+0 : off+2])
					g := p.Header.ByteOrder.Uint16(buf[off+2 : off+4])
					b := p.Header.ByteOrder.Uint16(buf[off+4 : off+6])
					a := p.Header.ByteOrder.Uint16(buf[off+6 : off+8])
					off += 8
					img.SetRGBA64(x, y, color.RGBA64{r, g, b, a})
				}
			}
		} else {
			img := dst.(*image.RGBA)
			for y := ymin; y < rMaxY; y++ {
				min := img.PixOffset(xmin, y)
				max := img.PixOffset(rMaxX, y)
				copy(img.Pix[min:max], buf[(y-ymin)*(xmax-xmin)*4:(y-ymin+1)*(xmax-xmin)*4])
			}
		}
	}

	return
}

func (p *IFD) EncodeBlock(w io.Writer, col, row int, dst *Image) (err error) {
	return
}
