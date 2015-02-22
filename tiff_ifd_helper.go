// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"image"
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
	if tag, ok := p.EntryMap[TagType_ImageWidth]; ok {
		if v := tag.GetInts(); len(v) == 1 {
			width = int(v[0])
		}
	}
	if tag, ok := p.EntryMap[TagType_ImageLength]; ok {
		if v := tag.GetInts(); len(v) == 1 {
			height = int(v[0])
		}
	}
	return image.Rect(0, 0, width, height)
}

func (p *IFD) Depth() int {
	if tag, ok := p.EntryMap[TagType_BitsPerSample]; ok {
		v := tag.GetInts()
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
	if tag, ok := p.EntryMap[TagType_SamplesPerPixel]; ok {
		if v := tag.GetInts(); len(v) == 1 {
			return int(v[0])
		}
	}
	return 0
}

func (p *IFD) ImageType() ImageType {
	var requiredTags = []TagType{
		TagType_ImageWidth,
		TagType_ImageLength,
		TagType_PhotometricInterpretation,
		TagType_XResolution,
		TagType_YResolution,
		TagType_ResolutionUnit,
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
		if _, ok := p.EntryMap[tag]; !ok {
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
		photometric, _                      = p.TagGetter().GetPhotometricInterpretation()
		_, hasBitsPerSample                 = p.TagGetter().GetBitsPerSample()
		samplesPerPixel, hasSamplesPerPixel = p.TagGetter().GetSamplesPerPixel()
		extraSamples, hasExtraSamples       = p.TagGetter().GetExtraSamples()
	)

	switch photometric {
	case TagValue_PhotometricType_WhiteIsZero:
		if !hasBitsPerSample {
			return ImageType_BilevelInvert
		} else {
			return ImageType_GrayInvert
		}
	case TagValue_PhotometricType_BlackIsZero:
		if !hasBitsPerSample {
			return ImageType_Bilevel
		} else {
			return ImageType_Gray
		}
	case TagValue_PhotometricType_RGB:
		if hasSamplesPerPixel && len(samplesPerPixel) == 3 {
			return ImageType_RGB
		}
		if hasSamplesPerPixel && len(samplesPerPixel) == 4 {
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

func (p *IFD) Compression() CompressType {
	if tag, ok := p.EntryMap[TagType_Compression]; ok {
		if v := tag.GetInts(); len(v) == 1 {
			return CompressType(v[0])
		}
	}
	return CompressType_Nil
}

func (p *IFD) ColorMap() [][3]uint16 {
	if tag, ok := p.EntryMap[TagType_ColorMap]; ok {
		v := tag.GetInts()
		if len(v) == 0 || len(v)%3 != 0 {
			return nil
		}
		colorMap := make([][3]uint16, len(v)/3)
		for i := 0; i < len(v); i += 3 {
			colorMap[i/3] = [3]uint16{
				uint16(v[i+0]),
				uint16(v[i+1]),
				uint16(v[i+2]),
			}
		}
	}
	return nil
}

func (p *IFD) BlockSize() (width, height int) {
	if _, ok := p.EntryMap[TagType_TileWidth]; ok {
		if tag, ok := p.EntryMap[TagType_TileWidth]; ok {
			if v := tag.GetInts(); len(v) == 1 {
				width = int(v[0])
			}
		}
		if tag, ok := p.EntryMap[TagType_TileLength]; ok {
			if v := tag.GetInts(); len(v) == 1 {
				height = int(v[0])
			}
		}

	} else {
		if tag, ok := p.EntryMap[TagType_ImageWidth]; ok {
			if v := tag.GetInts(); len(v) == 1 {
				width = int(v[0])
			}
		}
		if tag, ok := p.EntryMap[TagType_RowsPerStrip]; ok {
			if v := tag.GetInts(); len(v) == 1 {
				height = int(v[0])
			}
		}
		if height == 0 {
			if tag, ok := p.EntryMap[TagType_ImageLength]; ok {
				if v := tag.GetInts(); len(v) == 1 {
					height = int(v[0])
				}
			}
		}
	}
	return
}

func (p *IFD) BlockOffsets() []int64 {
	if _, ok := p.EntryMap[TagType_TileWidth]; ok {
		if tag, ok := p.EntryMap[TagType_TileOffsets]; ok {
			return tag.GetInts()
		}
	} else {
		if tag, ok := p.EntryMap[TagType_StripOffsets]; ok {
			return tag.GetInts()
		}
	}
	return nil
}

func (p *IFD) BlockCounts() []int64 {
	if _, ok := p.EntryMap[TagType_TileWidth]; ok {
		if tag, ok := p.EntryMap[TagType_TileByteCounts]; ok {
			return tag.GetInts()
		}
	} else {
		if tag, ok := p.EntryMap[TagType_StripByteCounts]; ok {
			return tag.GetInts()
		}
	}
	return nil
}

func (p *IFD) Predictor() TagValue_PredictorType {
	if tag, ok := p.EntryMap[TagType_Predictor]; ok {
		if v := tag.GetInts(); len(v) == 1 {
			return TagValue_PredictorType(v[0])
		}
	}
	return TagValue_PredictorType_None
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
	for _, v := range p.EntryMap {
		fmt.Fprintf(&buf, "  %v\n", v)
	}
	fmt.Fprintf(&buf, "  Next: %08x\n", p.Offset)
	fmt.Fprintf(&buf, "}\n")
	return buf.String()
}
