// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

func (p *IFD) Valid() bool {
	if p.Header == nil || !p.Header.Valid() {
		return false
	}
	return true
}

func (p *IFD) Size() (width, height int) {
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
	return
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
