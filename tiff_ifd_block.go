// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"fmt"
	"image"
	"io"
)

func (p *IFD) BlockSize() (width, height int) {
	if value, ok := p.TagGetter().GetTileWidth(); ok {
		width = int(value)
	}
	if value, ok := p.TagGetter().GetTileLength(); ok {
		height = int(value)
	}
	return
}

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
		if !ok || blockHeight == 0 {
			blockHeight = imageHeight
		}
		return int((imageHeight + blockHeight - 1) / blockHeight)
	}
}

func (p *IFD) BlockBounds(col, row int) image.Rectangle {
	blocksAcross, blocksDown := p.BlocksAcross(), p.BlocksDown()
	if col < 0 || row < 0 || col >= blocksAcross-1 || row >= blocksDown {
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

		blockWidth, _ := p.TagGetter().GetTileWidth()
		blockHeight, ok := p.TagGetter().GetRowsPerStrip()
		if !ok || blockHeight == 0 {
			blockHeight = imageHeight
		}

		blkW := blockWidth
		blkH := blockHeight

		if col == blocksAcross-1 && imageWidth%blockWidth != 0 {
			blkW = imageWidth % blockWidth
		}
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
	if col < 0 || row < 0 || col >= blocksAcross-1 || row >= blocksDown {
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
	if col < 0 || row < 0 || col >= blocksAcross-1 || row >= blocksDown {
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

func (p *IFD) DecodeBlock(r io.Reader, col, row int, dst *Image) (err error) {
	return
}

func (p *IFD) EncodeBlock(w io.Writer, col, row int, dst *Image) (err error) {
	return
}

func (p *IFD) _ReadBlock(r io.ReadSeeker, offset, length int64, dst *Image, rect image.Rectangle) (err error) {
	var data []byte
	if _, err = r.Seek(offset, 0); err != nil {
		return
	}
	limitReader := io.LimitReader(r, length)
	if data, err = p.Compression().ReadAll(limitReader); err != nil {
		return
	}

	predictor, ok := p.TagGetter().GetPredictor()
	if ok && predictor == TagValue_PredictorType_Horizontal {
		if data, err = p.decodePredictor(rect, data); err != nil {
			return
		}
	}

	return
}

func (p *IFD) decodePredictor(r image.Rectangle, data []byte) (out []byte, err error) {
	bpp := p.Depth()
	spp := p.Channels()

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
