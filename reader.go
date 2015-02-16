// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"io"
	"io/ioutil"
)

type decoder struct {
	r         *seekioReader
	byteOrder binary.ByteOrder
	config    image.Config
	mode      imageMode
	bpp       uint
	features  map[TagType][]uint
	palette   []color.Color

	buf   []byte
	off   int    // Current offset in buf.
	v     uint32 // Buffer value for reading with arbitrary bit depths.
	nbits uint   // Remaining number of bits in v.
}

func (d *decoder) Close() error {
	if d.r != nil {
		d.r.Close()
	}
	*d = decoder{}
	return nil
}

// firstVal returns the first uint of the features entry with the given tag,
// or 0 if the tag does not exist.
func (d *decoder) firstVal(tag TagType) uint {
	f := d.features[tag]
	if len(f) == 0 {
		return 0
	}
	return f[0]
}

// ifdUint decodes the IFD entry in p, which must be of the Byte, Short
// or Long type, and returns the decoded uint values.
func (d *decoder) ifdUint(p []byte) (u []uint, err error) {
	var raw []byte
	datatype := DataType(d.byteOrder.Uint16(p[2:4]))
	count := d.byteOrder.Uint32(p[4:8])
	if datalen := uint32(datatype.ByteSize()) * count; datalen > 4 {
		// The IFD contains a pointer to the real value.
		raw = make([]byte, datalen)
		_, err = d.r.ReadAt(raw, int64(d.byteOrder.Uint32(p[8:12])))
	} else {
		raw = p[8 : 8+datalen]
	}
	if err != nil {
		return nil, err
	}

	u = make([]uint, count)
	switch datatype {
	case DataType_Byte:
		for i := uint32(0); i < count; i++ {
			u[i] = uint(raw[i])
		}
	case DataType_Short:
		for i := uint32(0); i < count; i++ {
			u[i] = uint(d.byteOrder.Uint16(raw[2*i : 2*(i+1)]))
		}
	case DataType_Long:
		for i := uint32(0); i < count; i++ {
			u[i] = uint(d.byteOrder.Uint32(raw[4*i : 4*(i+1)]))
		}
	default:
		return nil, UnsupportedError("data type")
	}
	return u, nil
}

// parseIFD decides whether the the IFD entry in p is "interesting" and
// stows away the data in the decoder.
func (d *decoder) parseIFD(p []byte) error {
	tag := TagType(d.byteOrder.Uint16(p[0:2]))
	switch tag {
	case TagType_BitsPerSample,
		TagType_ExtraSamples,
		TagType_PhotometricInterpretation,
		TagType_Compression,
		TagType_Predictor,
		TagType_StripOffsets,
		TagType_StripByteCounts,
		TagType_RowsPerStrip,
		TagType_TileWidth,
		TagType_TileLength,
		TagType_TileOffsets,
		TagType_TileByteCounts,
		TagType_ImageLength,
		TagType_ImageWidth:
		val, err := d.ifdUint(p)
		if err != nil {
			return err
		}
		d.features[tag] = val
	case TagType_ColorMap:
		val, err := d.ifdUint(p)
		if err != nil {
			return err
		}
		numcolors := len(val) / 3
		if len(val)%3 != 0 || numcolors <= 0 || numcolors > 256 {
			return FormatError("bad ColorMap length")
		}
		d.palette = make([]color.Color, numcolors)
		for i := 0; i < numcolors; i++ {
			d.palette[i] = color.RGBA64{
				uint16(val[i]),
				uint16(val[i+numcolors]),
				uint16(val[i+2*numcolors]),
				0xffff,
			}
		}
	case TagType_SampleFormat:
		// Page 27 of the spec: If the SampleFormat is present and
		// the value is not 1 [= unsigned integer data], a Baseline
		// TIFF reader that cannot handle the SampleFormat value
		// must terminate the import process gracefully.
		val, err := d.ifdUint(p)
		if err != nil {
			return err
		}
		for _, v := range val {
			if v != 1 {
				return UnsupportedError("sample format")
			}
		}
	}
	return nil
}

// readBits reads n bits from the internal buffer starting at the current offset.
func (d *decoder) readBits(n uint) uint32 {
	for d.nbits < n {
		d.v <<= 8
		d.v |= uint32(d.buf[d.off])
		d.off++
		d.nbits += 8
	}
	d.nbits -= n
	rv := d.v >> d.nbits
	d.v &^= rv << d.nbits
	return rv
}

// flushBits discards the unread bits in the buffer used by readBits.
// It is used at the end of a line.
func (d *decoder) flushBits() {
	d.v = 0
	d.nbits = 0
}

// minInt returns the smaller of x or y.
func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// decode decodes the raw data of an image.
// It reads from d.buf and writes the strip or tile into dst.
func (d *decoder) decode(dst image.Image, xmin, ymin, xmax, ymax int) error {
	d.off = 0

	// Apply horizontal predictor if necessary.
	// In this case, p contains the color difference to the preceding pixel.
	// See page 64-65 of the spec.
	if d.firstVal(TagType_Predictor) == TagValue_Predictor_Horizontal {
		if d.bpp == 16 {
			var off int
			spp := len(d.features[TagType_BitsPerSample]) // samples per pixel
			bpp := spp * 2                                // bytes per pixel
			for y := ymin; y < ymax; y++ {
				off += spp * 2
				for x := 0; x < (xmax-xmin-1)*bpp; x += 2 {
					v0 := d.byteOrder.Uint16(d.buf[off-bpp : off-bpp+2])
					v1 := d.byteOrder.Uint16(d.buf[off : off+2])
					d.byteOrder.PutUint16(d.buf[off:off+2], v1+v0)
					off += 2
				}
			}
		} else if d.bpp == 8 {
			var off int
			spp := len(d.features[TagType_BitsPerSample]) // samples per pixel
			for y := ymin; y < ymax; y++ {
				off += spp
				for x := 0; x < (xmax-xmin-1)*spp; x++ {
					d.buf[off] += d.buf[off-spp]
					off++
				}
			}
		}
	}

	rMaxX := minInt(xmax, dst.Bounds().Max.X)
	rMaxY := minInt(ymax, dst.Bounds().Max.Y)
	switch d.mode {
	case mGray, mGrayInvert:
		if d.bpp == 16 {
			img := dst.(*image.Gray16)
			for y := ymin; y < rMaxY; y++ {
				for x := xmin; x < rMaxX; x++ {
					v := d.byteOrder.Uint16(d.buf[d.off : d.off+2])
					d.off += 2
					if d.mode == mGrayInvert {
						v = 0xffff - v
					}
					img.SetGray16(x, y, color.Gray16{v})
				}
			}
		} else {
			img := dst.(*image.Gray)
			max := uint32((1 << d.bpp) - 1)
			for y := ymin; y < rMaxY; y++ {
				for x := xmin; x < rMaxX; x++ {
					v := uint8(d.readBits(d.bpp) * 0xff / max)
					if d.mode == mGrayInvert {
						v = 0xff - v
					}
					img.SetGray(x, y, color.Gray{v})
				}
				d.flushBits()
			}
		}
	case mPaletted:
		img := dst.(*image.Paletted)
		for y := ymin; y < rMaxY; y++ {
			for x := xmin; x < rMaxX; x++ {
				img.SetColorIndex(x, y, uint8(d.readBits(d.bpp)))
			}
			d.flushBits()
		}
	case mRGB:
		if d.bpp == 16 {
			img := dst.(*image.RGBA64)
			for y := ymin; y < rMaxY; y++ {
				for x := xmin; x < rMaxX; x++ {
					r := d.byteOrder.Uint16(d.buf[d.off+0 : d.off+2])
					g := d.byteOrder.Uint16(d.buf[d.off+2 : d.off+4])
					b := d.byteOrder.Uint16(d.buf[d.off+4 : d.off+6])
					d.off += 6
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
					img.Pix[i+0] = d.buf[off+0]
					img.Pix[i+1] = d.buf[off+1]
					img.Pix[i+2] = d.buf[off+2]
					img.Pix[i+3] = 0xff
					off += 3
				}
			}
		}
	case mNRGBA:
		if d.bpp == 16 {
			img := dst.(*image.NRGBA64)
			for y := ymin; y < rMaxY; y++ {
				for x := xmin; x < rMaxX; x++ {
					r := d.byteOrder.Uint16(d.buf[d.off+0 : d.off+2])
					g := d.byteOrder.Uint16(d.buf[d.off+2 : d.off+4])
					b := d.byteOrder.Uint16(d.buf[d.off+4 : d.off+6])
					a := d.byteOrder.Uint16(d.buf[d.off+6 : d.off+8])
					d.off += 8
					img.SetNRGBA64(x, y, color.NRGBA64{r, g, b, a})
				}
			}
		} else {
			img := dst.(*image.NRGBA)
			for y := ymin; y < rMaxY; y++ {
				min := img.PixOffset(xmin, y)
				max := img.PixOffset(rMaxX, y)
				buf := d.buf[(y-ymin)*(xmax-xmin)*4 : (y-ymin+1)*(xmax-xmin)*4]
				copy(img.Pix[min:max], buf)
			}
		}
	case mRGBA:
		if d.bpp == 16 {
			img := dst.(*image.RGBA64)
			for y := ymin; y < rMaxY; y++ {
				for x := xmin; x < rMaxX; x++ {
					r := d.byteOrder.Uint16(d.buf[d.off+0 : d.off+2])
					g := d.byteOrder.Uint16(d.buf[d.off+2 : d.off+4])
					b := d.byteOrder.Uint16(d.buf[d.off+4 : d.off+6])
					a := d.byteOrder.Uint16(d.buf[d.off+6 : d.off+8])
					d.off += 8
					img.SetRGBA64(x, y, color.RGBA64{r, g, b, a})
				}
			}
		} else {
			img := dst.(*image.RGBA)
			for y := ymin; y < rMaxY; y++ {
				min := img.PixOffset(xmin, y)
				max := img.PixOffset(rMaxX, y)
				buf := d.buf[(y-ymin)*(xmax-xmin)*4 : (y-ymin+1)*(xmax-xmin)*4]
				copy(img.Pix[min:max], buf)
			}
		}
	}

	return nil
}

func openDecoder(r io.Reader) (*decoder, error) {
	d := &decoder{
		r:        openSeekioReader(r, -1),
		features: make(map[TagType][]uint),
	}

	p := make([]byte, 8)
	if _, err := d.r.ReadAt(p, 0); err != nil {
		return nil, err
	}
	switch string(p[0:4]) {
	case classicTiffLittleEnding:
		d.byteOrder = binary.LittleEndian
	case classicTiffBigEnding:
		d.byteOrder = binary.BigEndian
	default:
		return nil, FormatError("malformed header")
	}

	ifdOffset := int64(d.byteOrder.Uint32(p[4:8]))

	// The first two bytes contain the number of entries (12 bytes each).
	if _, err := d.r.ReadAt(p[0:2], ifdOffset); err != nil {
		return nil, err
	}
	numItems := int(d.byteOrder.Uint16(p[0:2]))

	// All IFD entries are read in one chunk.
	p = make([]byte, ifdLen*numItems)
	if _, err := d.r.ReadAt(p, ifdOffset+2); err != nil {
		return nil, err
	}

	for i := 0; i < len(p); i += ifdLen {
		if err := d.parseIFD(p[i : i+ifdLen]); err != nil {
			return nil, err
		}
	}

	d.config.Width = int(d.firstVal(TagType_ImageWidth))
	d.config.Height = int(d.firstVal(TagType_ImageLength))

	if _, ok := d.features[TagType_BitsPerSample]; !ok {
		return nil, FormatError("BitsPerSample tag missing")
	}
	d.bpp = d.firstVal(TagType_BitsPerSample)

	// Determine the image mode.
	switch d.firstVal(TagType_PhotometricInterpretation) {
	case TagValue_Photometric_RGB:
		if d.bpp == 16 {
			for _, b := range d.features[TagType_BitsPerSample] {
				if b != 16 {
					return nil, FormatError("wrong number of samples for 16bit RGB")
				}
			}
		} else {
			for _, b := range d.features[TagType_BitsPerSample] {
				if b != 8 {
					return nil, FormatError("wrong number of samples for 8bit RGB")
				}
			}
		}
		// RGB images normally have 3 samples per pixel.
		// If there are more, ExtraSamples (p. 31-32 of the spec)
		// gives their meaning (usually an alpha channel).
		//
		// This implementation does not support extra samples
		// of an unspecified type.
		switch len(d.features[TagType_BitsPerSample]) {
		case 3:
			d.mode = mRGB
			if d.bpp == 16 {
				d.config.ColorModel = color.RGBA64Model
			} else {
				d.config.ColorModel = color.RGBAModel
			}
		case 4:
			switch d.firstVal(TagType_ExtraSamples) {
			case 1:
				d.mode = mRGBA
				if d.bpp == 16 {
					d.config.ColorModel = color.RGBA64Model
				} else {
					d.config.ColorModel = color.RGBAModel
				}
			case 2:
				d.mode = mNRGBA
				if d.bpp == 16 {
					d.config.ColorModel = color.NRGBA64Model
				} else {
					d.config.ColorModel = color.NRGBAModel
				}
			default:
				return nil, FormatError("wrong number of samples for RGB")
			}
		default:
			return nil, FormatError("wrong number of samples for RGB")
		}
	case TagValue_Photometric_Paletted:
		d.mode = mPaletted
		d.config.ColorModel = color.Palette(d.palette)
	case TagValue_Photometric_WhiteIsZero:
		d.mode = mGrayInvert
		if d.bpp == 16 {
			d.config.ColorModel = color.Gray16Model
		} else {
			d.config.ColorModel = color.GrayModel
		}
	case TagValue_Photometric_BlackIsZero:
		d.mode = mGray
		if d.bpp == 16 {
			d.config.ColorModel = color.Gray16Model
		} else {
			d.config.ColorModel = color.GrayModel
		}
	default:
		return nil, UnsupportedError("color model")
	}

	return d, nil
}

// DecodeConfig returns the color model and dimensions of a TIFF image without
// decoding the entire image.
func DecodeConfig(r io.Reader) (cfg image.Config, err error) {
	d, err := openDecoder(r)
	if err != nil {
		return image.Config{}, err
	}
	defer func() {
		err = d.Close()
	}()
	cfg = d.config
	return
}

// Decode reads a TIFF image from r and returns it as an image.Image.
// The type of Image returned depends on the contents of the TIFF.
func Decode(r io.Reader) (img image.Image, err error) {
	d, err := openDecoder(r)
	if err != nil {
		return
	}
	defer func() {
		err = d.Close()
	}()

	blockPadding := false
	blockWidth := d.config.Width
	blockHeight := d.config.Height
	blocksAcross := 1
	blocksDown := 1

	var blockOffsets, blockCounts []uint

	if int(d.firstVal(TagType_TileWidth)) != 0 {
		blockPadding = true

		blockWidth = int(d.firstVal(TagType_TileWidth))
		blockHeight = int(d.firstVal(TagType_TileLength))

		blocksAcross = (d.config.Width + blockWidth - 1) / blockWidth
		blocksDown = (d.config.Height + blockHeight - 1) / blockHeight

		blockCounts = d.features[TagType_TileByteCounts]
		blockOffsets = d.features[TagType_TileOffsets]

	} else {
		if int(d.firstVal(TagType_RowsPerStrip)) != 0 {
			blockHeight = int(d.firstVal(TagType_RowsPerStrip))
		}

		blocksDown = (d.config.Height + blockHeight - 1) / blockHeight

		blockOffsets = d.features[TagType_StripOffsets]
		blockCounts = d.features[TagType_StripByteCounts]
	}

	// Check if we have the right number of strips/tiles, offsets and counts.
	if n := blocksAcross * blocksDown; len(blockOffsets) < n || len(blockCounts) < n {
		return nil, FormatError("inconsistent header")
	}

	imgRect := image.Rect(0, 0, d.config.Width, d.config.Height)
	switch d.mode {
	case mGray, mGrayInvert:
		if d.bpp == 16 {
			img = image.NewGray16(imgRect)
		} else {
			img = image.NewGray(imgRect)
		}
	case mPaletted:
		img = image.NewPaletted(imgRect, d.palette)
	case mNRGBA:
		if d.bpp == 16 {
			img = image.NewNRGBA64(imgRect)
		} else {
			img = image.NewNRGBA(imgRect)
		}
	case mRGB, mRGBA:
		if d.bpp == 16 {
			img = image.NewRGBA64(imgRect)
		} else {
			img = image.NewRGBA(imgRect)
		}
	}

	for i := 0; i < blocksAcross; i++ {
		blkW := blockWidth
		if !blockPadding && i == blocksAcross-1 && d.config.Width%blockWidth != 0 {
			blkW = d.config.Width % blockWidth
		}
		for j := 0; j < blocksDown; j++ {
			blkH := blockHeight
			if !blockPadding && j == blocksDown-1 && d.config.Height%blockHeight != 0 {
				blkH = d.config.Height % blockHeight
			}
			offset := int64(blockOffsets[j*blocksAcross+i])
			n := int64(blockCounts[j*blocksAcross+i])
			switch CompressType(d.firstVal(TagType_Compression)) {

			// According to the spec, Compression does not have a default value,
			// but some tools interpret a missing Compression value as none so we do
			// the same.
			case CompressType_None, 0:
				d.buf = make([]byte, n)
				_, err = d.r.ReadAt(d.buf, offset)
			case CompressType_LZW:
				r := newLzwReader(io.NewSectionReader(d.r, offset, n), lzwMSB, 8)
				d.buf, err = ioutil.ReadAll(r)
				r.Close()
			case CompressType_Deflate, CompressType_DeflateOld:
				r, err := zlib.NewReader(io.NewSectionReader(d.r, offset, n))
				if err != nil {
					return nil, err
				}
				d.buf, err = ioutil.ReadAll(r)
				r.Close()
			case CompressType_PackBits:
				d.buf, err = unpackBits(io.NewSectionReader(d.r, offset, n))
			default:
				err = UnsupportedError(fmt.Sprintf("compression value %d", d.firstVal(TagType_Compression)))
			}
			if err != nil {
				return nil, err
			}

			xmin := i * blockWidth
			ymin := j * blockHeight
			xmax := xmin + blkW
			ymax := ymin + blkH
			err = d.decode(img, xmin, ymin, xmax, ymax)
			if err != nil {
				return nil, err
			}
		}
	}
	return
}

func init() {
	image.RegisterFormat("tiff", classicTiffLittleEnding, Decode, DecodeConfig)
	image.RegisterFormat("tiff", classicTiffBigEnding, Decode, DecodeConfig)
}
