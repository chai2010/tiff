// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"fmt"
	"time"
)

func (p *IFD) ImageSize() (width, height int) {
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

func (p *IFD) BitsPerSample() (bits []int) {
	if tag, ok := p.EntryMap[TagType_BitsPerSample]; ok {
		v := tag.GetInts()
		bits := make([]int, len(v))
		for i := 0; i < len(v); i++ {
			bits[i] = int(v[i])
		}
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

func (p *IFD) CellSize() (width, height float64) {
	if tag, ok := p.EntryMap[TagType_CellWidth]; ok {
		if v := tag.GetFloats(); len(v) == 1 {
			width = float64(v[0])
		}
	}
	if tag, ok := p.EntryMap[TagType_CellLenght]; ok {
		if v := tag.GetFloats(); len(v) == 1 {
			height = float64(v[0])
		}
	}
	return
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

func (p *IFD) DocumentName() string {
	if tag, ok := p.EntryMap[TagType_DocumentName]; ok {
		return tag.GetString()
	}
	return ""
}

func (p *IFD) ImageDescription() string {
	if tag, ok := p.EntryMap[TagType_ImageDescription]; ok {
		return tag.GetString()
	}
	return ""
}

func (p *IFD) Make() string {
	if tag, ok := p.EntryMap[TagType_Make]; ok {
		return tag.GetString()
	}
	return ""
}

func (p *IFD) Model() string {
	if tag, ok := p.EntryMap[TagType_Model]; ok {
		return tag.GetString()
	}
	return ""
}

func (p *IFD) Software() string {
	if tag, ok := p.EntryMap[TagType_Software]; ok {
		return tag.GetString()
	}
	return ""
}
func (p *IFD) Artist() string {
	if tag, ok := p.EntryMap[TagType_Artist]; ok {
		return tag.GetString()
	}
	return ""
}
func (p *IFD) HostComputer() string {
	if tag, ok := p.EntryMap[TagType_HostComputer]; ok {
		return tag.GetString()
	}
	return ""
}

func (p *IFD) DateTime() (t time.Time) {
	if tag, ok := p.EntryMap[TagType_DateTime]; ok {
		var year, month, day, hour, min, sec int
		if _, err := fmt.Sscanf(tag.GetString(), "%d:%d:%d %d:%d:%d",
			&year, &month, &day,
			&hour, &min, &sec,
		); err == nil {
			t = time.Date(year, time.Month(month), day, hour, min, sec, 0, nil)
		}
	}
	return
}
func (p *IFD) Copyright() string {
	if tag, ok := p.EntryMap[TagType_Copyright]; ok {
		return tag.GetString()
	}
	return ""
}
