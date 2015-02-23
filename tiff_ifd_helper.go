// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"sort"
)

func (p *IFD) Valid() bool {
	if p.Header == nil || !p.Header.Valid() {
		return false
	}
	return true
}

func (p *IFD) TagGetter() TagGetter {
	return &tifTagGetter{
		EntryMap: p.EntryMap,
	}
}

func (p *IFD) TagSetter() TagSetter {
	return &tifTagSetter{
		EntryMap: p.EntryMap,
	}
}

func (p *IFD) Bounds() image.Rectangle {
	var width, height int
	if v, ok := p.TagGetter().GetImageWidth(); ok {
		width = int(v)
	}
	if v, ok := p.TagGetter().GetImageLength(); ok {
		height = int(v)
	}
	return image.Rect(0, 0, width, height)
}

func (p *IFD) Depth() int {
	if v, ok := p.TagGetter().GetBitsPerSample(); ok {
		for i := 1; i < len(v); i++ {
			if v[i] != v[0] {
				return 0
			}
		}
		if len(v) > 0 {
			return int(v[0])
		}
	}
	return 0
}

func (p *IFD) Channels() int {
	if v, ok := p.TagGetter().GetBitsPerSample(); ok {
		return len(v)
	}
	return 0
}

func (p *IFD) ImageType() ImageType {
	var requiredTags = []TagType{
		TagType_ImageWidth,
		TagType_ImageLength,
		TagType_PhotometricInterpretation,
	}
	var requiredTiledTags = []TagType{
		TagType_TileWidth,
		TagType_TileLength,
		TagType_TileOffsets,
		TagType_TileByteCounts,
	}
	var requiredStripTags = []TagType{
		TagType_RowsPerStrip, // default is (2^32-1)
		TagType_StripOffsets,
		TagType_StripByteCounts,
	}

	var isTiled bool
	for _, tag := range requiredTiledTags {
		if _, ok := p.EntryMap[tag]; ok {
			isTiled = true
		}
	}

	if isTiled {
		for _, tag := range requiredTags {
			if _, ok := p.EntryMap[tag]; !ok {
				return ImageType_Nil
			}
		}
		for _, tag := range requiredTiledTags {
			if _, ok := p.EntryMap[tag]; !ok {
				return ImageType_Nil
			}
		}

	} else {
		for _, tag := range requiredTags {
			if _, ok := p.EntryMap[tag]; !ok {
				return ImageType_Nil
			}
		}
		for _, tag := range requiredStripTags {
			if _, ok := p.EntryMap[tag]; !ok {
				if tag != TagType_RowsPerStrip {
					return ImageType_Nil
				}
			}
		}
	}

	var (
		photometric, _                = p.TagGetter().GetPhotometricInterpretation()
		bitsPerSample, _              = p.TagGetter().GetBitsPerSample()
		extraSamples, hasExtraSamples = p.TagGetter().GetExtraSamples()
	)

	switch photometric {
	case TagValue_PhotometricType_WhiteIsZero:
		if len(bitsPerSample) == 1 && bitsPerSample[0] < 8 {
			return ImageType_BilevelInvert
		} else {
			return ImageType_GrayInvert
		}
	case TagValue_PhotometricType_BlackIsZero:
		if len(bitsPerSample) == 1 && bitsPerSample[0] < 8 {
			return ImageType_Bilevel
		} else {
			return ImageType_Gray
		}
	case TagValue_PhotometricType_RGB:
		if p.Channels() == 3 {
			return ImageType_RGB
		}
		if p.Channels() == 4 {
			if hasExtraSamples && extraSamples == 1 {
				return ImageType_RGBA
			}
			if hasExtraSamples && extraSamples == 2 {
				return ImageType_NRGBA
			}
		}
		return ImageType_Nil
	case TagValue_PhotometricType_Paletted:
		return ImageType_Paletted
	case TagValue_PhotometricType_TransMask:
		return ImageType_Nil
	case TagValue_PhotometricType_CMYK:
		return ImageType_Nil
	case TagValue_PhotometricType_YCbCr:
		return ImageType_Nil
	case TagValue_PhotometricType_CIELab:
		return ImageType_Nil
	}

	return ImageType_Nil
}

func (p *IFD) ImageConfig() (config image.Config, err error) {
	var (
		imageWidth, _    = p.TagGetter().GetImageWidth()
		imageHeight, _   = p.TagGetter().GetImageLength()
		photometric, _   = p.TagGetter().GetPhotometricInterpretation()
		bitsPerSample, _ = p.TagGetter().GetBitsPerSample()
		extraSamples, _  = p.TagGetter().GetExtraSamples()
	)
	if len(bitsPerSample) == 0 {
		err = fmt.Errorf("tiff: IFD.ColorModel, bad bitsPerSample length")
		return
	}

	config.Width = int(imageWidth)
	config.Height = int(imageHeight)

	switch photometric {
	case TagValue_PhotometricType_RGB:
		if bitsPerSample[0] == 16 {
			for _, b := range bitsPerSample {
				if b != 16 {
					err = fmt.Errorf("tiff: IFD.ColorModel, wrong number of samples for 16bit RGB")
					return
				}
			}
		} else {
			for _, b := range bitsPerSample {
				if b != 8 {
					err = fmt.Errorf("tiff: IFD.ColorModel, wrong number of samples for 8bit RGB")
					return
				}
			}
		}
		switch len(bitsPerSample) {
		case 3:
			if bitsPerSample[0] == 16 {
				config.ColorModel = color.RGBA64Model
			} else {
				config.ColorModel = color.RGBAModel
			}
		case 4:
			switch extraSamples {
			case 1:
				if bitsPerSample[0] == 16 {
					config.ColorModel = color.RGBA64Model
				} else {
					config.ColorModel = color.RGBAModel
				}
			case 2:
				if bitsPerSample[0] == 16 {
					config.ColorModel = color.NRGBA64Model
				} else {
					config.ColorModel = color.NRGBAModel
				}
			default:
				err = fmt.Errorf("tiff: IFD.ColorModel, wrong number of samples for RGB")
				return
			}
		default:
			err = fmt.Errorf("tiff: IFD.ColorModel, wrong number of samples for RGB")
			return
		}
	case TagValue_PhotometricType_Paletted:
		config.ColorModel = color.Palette(p.ColorMap())
	case TagValue_PhotometricType_WhiteIsZero:
		if bitsPerSample[0] == 16 {
			config.ColorModel = color.Gray16Model
		} else {
			config.ColorModel = color.GrayModel
		}
	case TagValue_PhotometricType_BlackIsZero:
		if bitsPerSample[0] == 16 {
			config.ColorModel = color.Gray16Model
		} else {
			config.ColorModel = color.GrayModel
		}
	default:
		err = fmt.Errorf("tiff: decoder.Decode, unsupport color model")
		return
	}
	return
}

func (p *IFD) Compression() CompressType {
	if tag, ok := p.EntryMap[TagType_Compression]; ok {
		if v := tag.GetInts(); len(v) == 1 {
			return CompressType(v[0])
		}
	}
	return CompressType_Nil
}

func (p *IFD) ColorMap() (palette color.Palette) {
	v, ok := p.TagGetter().GetColorMap()
	if !ok {
		return
	}
	palette = make([]color.Color, len(v))
	for i := 0; i < len(palette); i++ {
		palette[i] = color.RGBA64{
			uint16(v[i][0]),
			uint16(v[i][1]),
			uint16(v[i][2]),
			0xffff,
		}
	}
	return
}

func (p *IFD) Bytes() []byte {
	if !p.Valid() {
		return nil
	}

	tagList := make([]*IFDEntry, 0, len(p.EntryMap))
	for _, v := range p.EntryMap {
		tagList = append(tagList, v)
	}
	sort.Sort(byIFDEntry(tagList))

	var buf bytes.Buffer
	if p.Header.TiffType == TiffType_ClassicTIFF {
		binary.Write(&buf, p.Header.ByteOrder, uint32(len(tagList)))
		for i := 0; i < len(tagList); i++ {
			entryBytes, _ := tagList[i].Bytes()
			buf.Write(entryBytes)
		}
		binary.Write(&buf, p.Header.ByteOrder, uint32(p.Offset))
	} else {
		binary.Write(&buf, p.Header.ByteOrder, uint32(len(tagList)))
		for i := 0; i < len(tagList); i++ {
			entryBytes, _ := tagList[i].Bytes()
			buf.Write(entryBytes)
		}
		binary.Write(&buf, p.Header.ByteOrder, uint32(p.Offset))
	}
	return buf.Bytes()
}

func (p *IFD) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "tiff.IFD {\n")

	tagList := make([]*IFDEntry, 0, len(p.EntryMap))
	for _, v := range p.EntryMap {
		tagList = append(tagList, v)
	}
	sort.Sort(byIFDEntry(tagList))

	for _, v := range tagList {
		fmt.Fprintf(&buf, "  %v\n", v)
	}
	fmt.Fprintf(&buf, "  Next: %08x\n", p.Offset)
	fmt.Fprintf(&buf, "}\n")
	return buf.String()
}
