// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"image"
	"io"
	"sort"
)

// The TIFF format allows to choose the order of the different elements freely.
// The basic structure of a TIFF file written by this package is:
//
//   1. Header (8 bytes).
//   2. Image data.
//   3. Image File Directory (IFD).
//   4. "Pointer area" for larger entries in the IFD.

// We only write little-endian TIFF files.
var enc = binary.LittleEndian

// An ifdEntry is a single entry in an Image File Directory.
// A value of type DataType_Rational is composed of two 32-bit values,
// thus data contains two uints (numerator and denominator) for a single number.
type ifdEntry struct {
	tag      TagType
	datatype DataType
	data     []uint32
}

func (e ifdEntry) putData(p []byte) {
	for _, d := range e.data {
		switch e.datatype {
		case DataType_Byte, DataType_ASCII:
			p[0] = byte(d)
			p = p[1:]
		case DataType_Short:
			enc.PutUint16(p, uint16(d))
			p = p[2:]
		case DataType_Long, DataType_Rational:
			enc.PutUint32(p, uint32(d))
			p = p[4:]
		}
	}
}

type byTag []ifdEntry

func (d byTag) Len() int           { return len(d) }
func (d byTag) Less(i, j int) bool { return d[i].tag < d[j].tag }
func (d byTag) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

func encodeGray(w io.Writer, pix []uint8, dx, dy, stride int, predictor bool) error {
	if !predictor {
		return writePix(w, pix, dy, dx, stride)
	}
	buf := make([]byte, dx)
	for y := 0; y < dy; y++ {
		min := y*stride + 0
		max := y*stride + dx
		off := 0
		var v0 uint8
		for i := min; i < max; i++ {
			v1 := pix[i]
			buf[off] = v1 - v0
			v0 = v1
			off++
		}
		if _, err := w.Write(buf); err != nil {
			return err
		}
	}
	return nil
}

func encodeGray16(w io.Writer, pix []uint8, dx, dy, stride int, predictor bool) error {
	buf := make([]byte, dx*2)
	for y := 0; y < dy; y++ {
		min := y*stride + 0
		max := y*stride + dx*2
		off := 0
		var v0 uint16
		for i := min; i < max; i += 2 {
			// An image.Gray16's Pix is in big-endian order.
			v1 := uint16(pix[i])<<8 | uint16(pix[i+1])
			if predictor {
				v0, v1 = v1, v1-v0
			}
			// We only write little-endian TIFF files.
			buf[off+0] = byte(v1)
			buf[off+1] = byte(v1 >> 8)
			off += 2
		}
		if _, err := w.Write(buf); err != nil {
			return err
		}
	}
	return nil
}

func encodeRGBA(w io.Writer, pix []uint8, dx, dy, stride int, predictor bool) error {
	if !predictor {
		return writePix(w, pix, dy, dx*4, stride)
	}
	buf := make([]byte, dx*4)
	for y := 0; y < dy; y++ {
		min := y*stride + 0
		max := y*stride + dx*4
		off := 0
		var r0, g0, b0, a0 uint8
		for i := min; i < max; i += 4 {
			r1, g1, b1, a1 := pix[i+0], pix[i+1], pix[i+2], pix[i+3]
			buf[off+0] = r1 - r0
			buf[off+1] = g1 - g0
			buf[off+2] = b1 - b0
			buf[off+3] = a1 - a0
			off += 4
			r0, g0, b0, a0 = r1, g1, b1, a1
		}
		if _, err := w.Write(buf); err != nil {
			return err
		}
	}
	return nil
}

func encodeRGBA64(w io.Writer, pix []uint8, dx, dy, stride int, predictor bool) error {
	buf := make([]byte, dx*8)
	for y := 0; y < dy; y++ {
		min := y*stride + 0
		max := y*stride + dx*8
		off := 0
		var r0, g0, b0, a0 uint16
		for i := min; i < max; i += 8 {
			// An image.RGBA64's Pix is in big-endian order.
			r1 := uint16(pix[i+0])<<8 | uint16(pix[i+1])
			g1 := uint16(pix[i+2])<<8 | uint16(pix[i+3])
			b1 := uint16(pix[i+4])<<8 | uint16(pix[i+5])
			a1 := uint16(pix[i+6])<<8 | uint16(pix[i+7])
			if predictor {
				r0, r1 = r1, r1-r0
				g0, g1 = g1, g1-g0
				b0, b1 = b1, b1-b0
				a0, a1 = a1, a1-a0
			}
			// We only write little-endian TIFF files.
			buf[off+0] = byte(r1)
			buf[off+1] = byte(r1 >> 8)
			buf[off+2] = byte(g1)
			buf[off+3] = byte(g1 >> 8)
			buf[off+4] = byte(b1)
			buf[off+5] = byte(b1 >> 8)
			buf[off+6] = byte(a1)
			buf[off+7] = byte(a1 >> 8)
			off += 8
		}
		if _, err := w.Write(buf); err != nil {
			return err
		}
	}
	return nil
}

func encode(w io.Writer, m image.Image, predictor bool) error {
	bounds := m.Bounds()
	buf := make([]byte, 4*bounds.Dx())
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		off := 0
		if predictor {
			var r0, g0, b0, a0 uint8
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, a := m.At(x, y).RGBA()
				r1 := uint8(r >> 8)
				g1 := uint8(g >> 8)
				b1 := uint8(b >> 8)
				a1 := uint8(a >> 8)
				buf[off+0] = r1 - r0
				buf[off+1] = g1 - g0
				buf[off+2] = b1 - b0
				buf[off+3] = a1 - a0
				off += 4
				r0, g0, b0, a0 = r1, g1, b1, a1
			}
		} else {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				r, g, b, a := m.At(x, y).RGBA()
				buf[off+0] = uint8(r >> 8)
				buf[off+1] = uint8(g >> 8)
				buf[off+2] = uint8(b >> 8)
				buf[off+3] = uint8(a >> 8)
				off += 4
			}
		}
		if _, err := w.Write(buf); err != nil {
			return err
		}
	}
	return nil
}

// writePix writes the internal byte array of an image to w. It is less general
// but much faster then encode. writePix is used when pix directly
// corresponds to one of the TIFF image types.
func writePix(w io.Writer, pix []byte, nrows, length, stride int) error {
	if length == stride {
		_, err := w.Write(pix[:nrows*length])
		return err
	}
	for ; nrows > 0; nrows-- {
		if _, err := w.Write(pix[:length]); err != nil {
			return err
		}
		pix = pix[stride:]
	}
	return nil
}

func writeIFD(w io.Writer, ifdOffset int, d []ifdEntry) error {
	const ifdLen = 12

	var buf [ifdLen]byte
	// Make space for "pointer area" containing IFD entry data
	// longer than 4 bytes.
	parea := make([]byte, 1024)
	pstart := ifdOffset + ifdLen*len(d) + 6
	var o int // Current offset in parea.

	// The IFD has to be written with the tags in ascending order.
	sort.Sort(byTag(d))

	// Write the number of entries in this IFD.
	if err := binary.Write(w, enc, uint16(len(d))); err != nil {
		return err
	}
	for _, ent := range d {
		enc.PutUint16(buf[0:2], uint16(ent.tag))
		enc.PutUint16(buf[2:4], uint16(ent.datatype))
		count := uint32(len(ent.data))
		if ent.datatype == DataType_Rational {
			count /= 2
		}
		enc.PutUint32(buf[4:8], count)
		datalen := int(count * uint32(ent.datatype.ByteSize()))
		if datalen <= 4 {
			ent.putData(buf[8:12])
		} else {
			if (o + datalen) > len(parea) {
				newlen := len(parea) + 1024
				for (o + datalen) > newlen {
					newlen += 1024
				}
				newarea := make([]byte, newlen)
				copy(newarea, parea)
				parea = newarea
			}
			ent.putData(parea[o : o+datalen])
			enc.PutUint32(buf[8:12], uint32(pstart+o))
			o += datalen
		}
		if _, err := w.Write(buf[:]); err != nil {
			return err
		}
	}
	// The IFD ends with the offset of the next IFD in the file,
	// or zero if it is the last one (page 14).
	if err := binary.Write(w, enc, uint32(0)); err != nil {
		return err
	}
	_, err := w.Write(parea[:o])
	return err
}

func EncodeAll(w io.Writer, m [][]image.Image, opt [][]*Options) error {
	return nil
}

// Encode writes the image m to w. opt determines the options used for
// encoding, such as the compression type. If opt is nil, an uncompressed
// image is written.
func Encode(w io.Writer, m image.Image, opt *Options) error {
	d := m.Bounds().Size()

	compression := CompressType_None
	predictor := false

	_, err := io.WriteString(w, ClassicTiffLittleEnding)
	if err != nil {
		return err
	}

	// Compressed data is written into a buffer first, so that we
	// know the compressed size.
	var buf bytes.Buffer
	// dst holds the destination for the pixel data of the image --
	// either w or a writer to buf.
	var dst io.Writer
	// imageLen is the length of the pixel data in bytes.
	// The offset of the IFD is imageLen + 8 header bytes.
	var imageLen int

	switch compression {
	case CompressType_None:
		dst = w
		// Write IFD offset before outputting pixel data.
		switch m.(type) {
		case *image.Paletted:
			imageLen = d.X * d.Y * 1
		case *image.Gray:
			imageLen = d.X * d.Y * 1
		case *image.Gray16:
			imageLen = d.X * d.Y * 2
		case *image.RGBA64:
			imageLen = d.X * d.Y * 8
		case *image.NRGBA64:
			imageLen = d.X * d.Y * 8
		default:
			imageLen = d.X * d.Y * 4
		}
		err = binary.Write(w, enc, uint32(imageLen+8))
		if err != nil {
			return err
		}
	case CompressType_Deflate:
		dst = zlib.NewWriter(&buf)
	}

	pr := uint32(TagValue_PredictorType_None)
	photometricInterpretation := uint32(TagValue_PhotometricType_RGB)
	samplesPerPixel := uint32(4)
	bitsPerSample := []uint32{8, 8, 8, 8}
	extraSamples := uint32(0)
	colorMap := []uint32{}

	if predictor {
		pr = uint32(TagValue_PredictorType_Horizontal)
	}
	switch m := m.(type) {
	case *image.Paletted:
		photometricInterpretation = uint32(TagValue_PhotometricType_Paletted)
		samplesPerPixel = 1
		bitsPerSample = []uint32{8}
		colorMap = make([]uint32, 256*3)
		for i := 0; i < 256 && i < len(m.Palette); i++ {
			r, g, b, _ := m.Palette[i].RGBA()
			colorMap[i+0*256] = uint32(r)
			colorMap[i+1*256] = uint32(g)
			colorMap[i+2*256] = uint32(b)
		}
		err = encodeGray(dst, m.Pix, d.X, d.Y, m.Stride, predictor)
	case *image.Gray:
		photometricInterpretation = uint32(TagValue_PhotometricType_BlackIsZero)
		samplesPerPixel = 1
		bitsPerSample = []uint32{8}
		err = encodeGray(dst, m.Pix, d.X, d.Y, m.Stride, predictor)
	case *image.Gray16:
		photometricInterpretation = uint32(TagValue_PhotometricType_BlackIsZero)
		samplesPerPixel = 1
		bitsPerSample = []uint32{16}
		err = encodeGray16(dst, m.Pix, d.X, d.Y, m.Stride, predictor)
	case *image.NRGBA:
		extraSamples = 2 // Unassociated alpha.
		err = encodeRGBA(dst, m.Pix, d.X, d.Y, m.Stride, predictor)
	case *image.NRGBA64:
		extraSamples = 2 // Unassociated alpha.
		bitsPerSample = []uint32{16, 16, 16, 16}
		err = encodeRGBA64(dst, m.Pix, d.X, d.Y, m.Stride, predictor)
	case *image.RGBA:
		extraSamples = 1 // Associated alpha.
		err = encodeRGBA(dst, m.Pix, d.X, d.Y, m.Stride, predictor)
	case *image.RGBA64:
		extraSamples = 1 // Associated alpha.
		bitsPerSample = []uint32{16, 16, 16, 16}
		err = encodeRGBA64(dst, m.Pix, d.X, d.Y, m.Stride, predictor)
	default:
		extraSamples = 1 // Associated alpha.
		err = encode(dst, m, predictor)
	}
	if err != nil {
		return err
	}

	if compression != CompressType_None {
		if err = dst.(io.Closer).Close(); err != nil {
			return err
		}
		imageLen = buf.Len()
		if err = binary.Write(w, enc, uint32(imageLen+8)); err != nil {
			return err
		}
		if _, err = buf.WriteTo(w); err != nil {
			return err
		}
	}

	ifd := []ifdEntry{
		{TagType_ImageWidth, DataType_Short, []uint32{uint32(d.X)}},
		{TagType_ImageLength, DataType_Short, []uint32{uint32(d.Y)}},
		{TagType_BitsPerSample, DataType_Short, bitsPerSample},
		{TagType_Compression, DataType_Short, []uint32{uint32(compression)}},
		{TagType_PhotometricInterpretation, DataType_Short, []uint32{photometricInterpretation}},
		{TagType_StripOffsets, DataType_Long, []uint32{8}},
		{TagType_SamplesPerPixel, DataType_Short, []uint32{samplesPerPixel}},
		{TagType_RowsPerStrip, DataType_Short, []uint32{uint32(d.Y)}},
		{TagType_StripByteCounts, DataType_Long, []uint32{uint32(imageLen)}},
		// There is currently no support for storing the image
		// resolution, so give a bogus value of 72x72 dpi.
		{TagType_XResolution, DataType_Rational, []uint32{72, 1}},
		{TagType_YResolution, DataType_Rational, []uint32{72, 1}},
		{TagType_ResolutionUnit, DataType_Short, []uint32{uint32(TagValue_ResolutionUnitType_PerInch)}},
	}
	if pr != uint32(TagValue_PredictorType_None) {
		ifd = append(ifd, ifdEntry{TagType_Predictor, DataType_Short, []uint32{pr}})
	}
	if len(colorMap) != 0 {
		ifd = append(ifd, ifdEntry{TagType_ColorMap, DataType_Short, colorMap})
	}
	if extraSamples > 0 {
		ifd = append(ifd, ifdEntry{TagType_ExtraSamples, DataType_Short, []uint32{extraSamples}})
	}

	return writeIFD(w, imageLen+8, ifd)
}
