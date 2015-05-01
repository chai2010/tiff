// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

type IFDEntry struct {
	Header   *Header
	Tag      TagType
	DataType DataType
	Count    int
	Offset   int64
	Data     []byte
}

type byIFDEntry []*IFDEntry

func (d byIFDEntry) Len() int           { return len(d) }
func (d byIFDEntry) Less(i, j int) bool { return d[i].Tag < d[j].Tag }
func (d byIFDEntry) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

func NewIFDEntry(hdr *Header, tag TagType, dataType DataType, data interface{}) (p *IFDEntry) {
	p = &IFDEntry{
		Header:   hdr,
		Tag:      tag,
		DataType: dataType,
	}
	p.SetValue(data)
	return p
}

func (p *IFDEntry) Valid() bool {
	if p == nil {
		return false
	}
	if !p.Header.Valid() || !p.Tag.Valid() || !p.DataType.Valid() {
		return false
	}
	if p.Count <= 0 || len(p.Data) == 0 {
		return false
	}
	return true
}

func (p *IFDEntry) Bytes() (entry, data []byte) {
	if p.Header.TiffType == TiffType_ClassicTIFF {
		var buf bytes.Buffer
		binary.Write(&buf, p.Header.ByteOrder, uint16(p.Tag))
		binary.Write(&buf, p.Header.ByteOrder, uint16(p.DataType))
		binary.Write(&buf, p.Header.ByteOrder, uint32(p.Count))

		offsetOrData := make([]byte, 4)
		if len(p.Data) > 4 {
			p.Header.ByteOrder.PutUint32(offsetOrData, uint32(p.Offset))
			data = p.Data
		} else {
			copy(offsetOrData[:], p.Data)
		}
		buf.Write(offsetOrData)
		entry = buf.Bytes()
		return
	} else {
		var buf bytes.Buffer
		binary.Write(&buf, p.Header.ByteOrder, uint16(p.Tag))
		binary.Write(&buf, p.Header.ByteOrder, uint16(p.DataType))
		binary.Write(&buf, p.Header.ByteOrder, uint64(p.Count))

		offsetOrData := make([]byte, 8)
		if len(p.Data) > 8 {
			p.Header.ByteOrder.PutUint64(offsetOrData, uint64(p.Offset))
			data = p.Data
		} else {
			copy(offsetOrData[:], p.Data)
		}
		buf.Write(offsetOrData)
		entry = buf.Bytes()
		return
	}
}

func (p *IFDEntry) GetInts() []int64 {
	switch p.DataType {
	case DataType_Byte:
		dst := make([]int64, p.Count)
		for i := 0; i < p.Count; i++ {
			dst[i] = int64(int8(p.Data[i]))
		}
		return dst
	case DataType_SByte:
		dst := make([]int64, p.Count)
		for i := 0; i < p.Count; i++ {
			dst[i] = int64(uint8(p.Data[i]))
		}
		return dst
	case DataType_Short:
		r := bytes.NewReader(p.Data)
		dst := make([]int64, p.Count)
		for i := 0; i < p.Count; i++ {
			var v uint16
			if err := binary.Read(r, p.Header.ByteOrder, &v); err != nil {
				return nil
			}
			dst[i] = int64(v)
		}
		return dst
	case DataType_SShort:
		r := bytes.NewReader(p.Data)
		dst := make([]int64, p.Count)
		for i := 0; i < p.Count; i++ {
			var v int16
			if err := binary.Read(r, p.Header.ByteOrder, &v); err != nil {
				return nil
			}
			dst[i] = int64(v)
		}
		return dst
	case DataType_Long, DataType_IFD:
		r := bytes.NewReader(p.Data)
		dst := make([]int64, p.Count)
		for i := 0; i < p.Count; i++ {
			var v uint32
			if err := binary.Read(r, p.Header.ByteOrder, &v); err != nil {
				return nil
			}
			dst[i] = int64(v)
		}
		return dst
	case DataType_SLong:
		r := bytes.NewReader(p.Data)
		dst := make([]int64, p.Count)
		for i := 0; i < p.Count; i++ {
			var v int32
			if err := binary.Read(r, p.Header.ByteOrder, &v); err != nil {
				return nil
			}
			dst[i] = int64(v)
		}
		return dst
	case DataType_Long8, DataType_IFD8:
		r := bytes.NewReader(p.Data)
		dst := make([]int64, p.Count)
		for i := 0; i < p.Count; i++ {
			var v uint64
			if err := binary.Read(r, p.Header.ByteOrder, &v); err != nil {
				return nil
			}
			dst[i] = int64(v)
		}
		return dst
	case DataType_SLong8:
		r := bytes.NewReader(p.Data)
		dst := make([]int64, p.Count)
		for i := 0; i < p.Count; i++ {
			var v int64
			if err := binary.Read(r, p.Header.ByteOrder, &v); err != nil {
				return nil
			}
			dst[i] = int64(v)
		}
		return dst
	}
	return nil
}

func (p *IFDEntry) GetFloats() []float64 {
	switch p.DataType {
	case DataType_Float:
		r := bytes.NewReader(p.Data)
		dst := make([]float64, p.Count)
		for i := 0; i < p.Count; i++ {
			var v float32
			if err := binary.Read(r, p.Header.ByteOrder, &v); err != nil {
				return nil
			}
			dst[i] = float64(v)
		}
		return dst
	case DataType_Double:
		r := bytes.NewReader(p.Data)
		dst := make([]float64, p.Count)
		for i := 0; i < p.Count; i++ {
			var v float64
			if err := binary.Read(r, p.Header.ByteOrder, &v); err != nil {
				return nil
			}
			dst[i] = float64(v)
		}
		return dst
	case DataType_Rational, DataType_SRational:
		rats := p.GetRationals()
		floats := make([]float64, len(rats))
		for i := 0; i < len(rats); i++ {
			floats[i] = float64(rats[i][0]) / float64(rats[i][1])
		}
		return floats
	default:
		ints := p.GetInts()
		floats := make([]float64, len(ints))
		for i := 0; i < len(ints); i++ {
			floats[i] = float64(ints[i])
		}
		return floats
	}
}

func (p *IFDEntry) GetRationals() [][2]int64 {
	switch p.DataType {
	case DataType_Rational:
		r := bytes.NewReader(p.Data)
		dst := make([][2]int64, p.Count)
		for i := 0; i < p.Count; i++ {
			var v [2]uint32
			if err := binary.Read(r, p.Header.ByteOrder, &v); err != nil {
				return nil
			}
			dst[i][0] = int64(v[0])
			dst[i][1] = int64(v[1])
		}
		return dst
	case DataType_SRational:
		r := bytes.NewReader(p.Data)
		dst := make([][2]int64, p.Count)
		for i := 0; i < p.Count; i++ {
			var v [2]int32
			if err := binary.Read(r, p.Header.ByteOrder, &v); err != nil {
				return nil
			}
			dst[i][0] = int64(v[0])
			dst[i][1] = int64(v[1])
		}
		return dst
	}
	return nil
}

func (p *IFDEntry) GetString() string {
	switch p.DataType {
	case DataType_ASCII:
		if idx := bytes.Index(p.Data, []byte("\000")); idx >= 0 {
			return string(p.Data[:idx])
		}
		return string(p.Data)
	}
	return ""
}

func (p *IFDEntry) GetValue() interface{} {
	panic("TODO")
}

func (p *IFDEntry) GetUndefined(value interface{}) interface{} {
	if p.DataType != DataType_Undefined {
		return nil
	}
	if err := binary.Read(bytes.NewReader(p.Data), p.Header.ByteOrder, value); err != nil {
		return nil
	}
	return value
}

func (p *IFDEntry) SetInts(value ...int64) (err error) {
	if p.DataType == DataType_Nil {
		if p.Header.IsBigTiff() {
			p.DataType = DataType_Long8
		} else {
			p.DataType = DataType_Long
		}
	}
	switch p.DataType {
	case DataType_Byte:
		tmp := make([]uint8, len(value))
		for i := 0; i < len(tmp); i++ {
			tmp[i] = uint8(value[i])
		}
		var buf bytes.Buffer
		if err = binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(tmp)
	case DataType_SByte:
		tmp := make([]int8, len(value))
		for i := 0; i < len(tmp); i++ {
			tmp[i] = int8(value[i])
		}
		var buf bytes.Buffer
		if err = binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(tmp)
	case DataType_Short:
		tmp := make([]uint16, len(value))
		for i := 0; i < len(tmp); i++ {
			tmp[i] = uint16(value[i])
		}
		var buf bytes.Buffer
		if err = binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(tmp)
	case DataType_SShort:
		tmp := make([]int16, len(value))
		for i := 0; i < len(tmp); i++ {
			tmp[i] = int16(value[i])
		}
		var buf bytes.Buffer
		if err = binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(tmp)
	case DataType_Long, DataType_IFD:
		tmp := make([]uint32, len(value))
		for i := 0; i < len(tmp); i++ {
			tmp[i] = uint32(value[i])
		}
		var buf bytes.Buffer
		if err = binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(tmp)
	case DataType_SLong:
		tmp := make([]int32, len(value))
		for i := 0; i < len(tmp); i++ {
			tmp[i] = int32(value[i])
		}
		var buf bytes.Buffer
		if err = binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(tmp)
	case DataType_Long8, DataType_IFD8:
		tmp := make([]uint64, len(value))
		for i := 0; i < len(tmp); i++ {
			tmp[i] = uint64(value[i])
		}
		var buf bytes.Buffer
		if err = binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(tmp)
	case DataType_SLong8:
		tmp := make([]int64, len(value))
		for i := 0; i < len(tmp); i++ {
			tmp[i] = int64(value[i])
		}
		var buf bytes.Buffer
		if err = binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(tmp)
	}
	return
}

func (p *IFDEntry) SetFloats(value ...float64) (err error) {
	if p.DataType == DataType_Nil {
		p.DataType = DataType_Double
	}
	switch p.DataType {
	case DataType_Float:
		tmp := make([]float32, len(value))
		for i := 0; i < len(tmp); i++ {
			tmp[i] = float32(value[i])
		}

		var buf bytes.Buffer
		if err = binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(tmp)
	case DataType_Double:
		var buf bytes.Buffer
		if err = binary.Write(&buf, p.Header.ByteOrder, value); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(value)
	}
	return
}

func (p *IFDEntry) SetRationals(value ...[2]int64) (err error) {
	if p.DataType == DataType_Nil {
		p.DataType = DataType_Rational
	}
	switch p.DataType {
	case DataType_Rational:
		tmp := make([][2]uint32, len(value))
		for i := 0; i < len(tmp); i++ {
			tmp[i][0] = uint32(value[i][0])
			tmp[i][1] = uint32(value[i][1])
		}

		var buf bytes.Buffer
		if err = binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(tmp)
	case DataType_SRational:
		tmp := make([][2]int32, len(value))
		for i := 0; i < len(tmp); i++ {
			tmp[i][0] = int32(value[i][0])
			tmp[i][1] = int32(value[i][1])
		}

		var buf bytes.Buffer
		if err = binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(tmp)
	}
	return
}

func (p *IFDEntry) SetString(value string) (err error) {
	if p.DataType == DataType_Nil {
		p.DataType = DataType_ASCII
	}
	switch p.DataType {
	case DataType_ASCII:
		if idx := strings.Index(value, "\000"); idx >= 0 {
			value = value[:idx]
		}
		p.Data = make([]byte, len(value)+1)
		copy(p.Data, []byte(value))
		p.Data[len(value)] = 0 // +NULL
		p.Count = len(p.Data) + 1
	}
	return
}

func (p *IFDEntry) SetUndefined(value interface{}) (err error) {
	if p.DataType == DataType_Nil {
		p.DataType = DataType_Undefined
	}
	if p.DataType != DataType_Undefined {
		return
	}
	var buf bytes.Buffer
	if err = binary.Write(&buf, p.Header.ByteOrder, value); err != nil {
		return
	}
	p.Data = buf.Bytes()
	p.Count = len(p.Data)
	return
}

func (p *IFDEntry) SetValue(value interface{}) (err error) {
	panic("TODO")
}

func (p *IFDEntry) isOnlyOneValue() bool {
	nums, _ := _TagType_NumsTable[p.Tag]
	return len(nums) == 1 && nums[0] == 1
}

func (p *IFDEntry) String() string {
	switch p.Tag {
	case TagType_NewSubfileType:
		if v := p.GetInts(); len(v) == 1 {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, TagValue_NewSubfileType(v[0]))
		}
	case TagType_SubfileType:
		if v := p.GetInts(); len(v) == 1 {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, TagValue_SubfileType(v[0]))
		}
	case TagType_ImageWidth:
		if v := p.GetInts(); len(v) == 1 {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, int(v[0]))
		}
	case TagType_ImageLength:
		if v := p.GetInts(); len(v) == 1 {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, int(v[0]))
		}
	case TagType_Compression:
		if v := p.GetInts(); len(v) == 1 {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, TagValue_CompressionType(v[0]))
		}
	case TagType_PhotometricInterpretation:
		if v := p.GetInts(); len(v) == 1 {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, TagValue_PhotometricType(v[0]))
		}
	case TagType_RowsPerStrip:
		if v := p.GetInts(); len(v) == 1 {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, int(v[0]))
		}
	case TagType_SamplesPerPixel:
		if v := p.GetInts(); len(v) == 1 {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, int(v[0]))
		}
	case TagType_ResolutionUnit:
		if v := p.GetInts(); len(v) == 1 {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, TagValue_ResolutionUnitType(v[0]))
		}
	case TagType_Predictor:
		if v := p.GetInts(); len(v) == 1 {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, TagValue_PredictorType(v[0]))
		}
	case TagType_StripOffsets, TagType_TileOffsets, TagType_FreeOffsets:
		return fmt.Sprintf("%v(%v): %#08x", p.Tag, p.DataType, p.GetInts())
	case TagType_SampleFormat:
		if v := p.GetInts(); len(v) == 1 {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, TagValue_SampleFormatType(v[0]))
		}
	}
	switch {
	case p.DataType.IsIntType():
		switch p.DataType {
		case DataType_IFD, DataType_IFD8:
			return fmt.Sprintf("%v(%v): %#08x", p.Tag, p.DataType, p.GetInts())
		}
		if v := p.GetInts(); len(v) == 1 && p.isOnlyOneValue() {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, v[0])
		} else {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, v)
		}
	case p.DataType.IsFloatType():
		if v := p.GetFloats(); len(v) == 1 && p.isOnlyOneValue() {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, v[0])
		} else {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, v)
		}
	case p.DataType.IsRationalType():
		if v := p.GetRationals(); len(v) == 1 && p.isOnlyOneValue() {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, v[0])
		} else {
			return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, v)
		}
	case p.DataType.IsStringType():
		return fmt.Sprintf("%v(%v): %q", p.Tag, p.DataType, p.GetString())
	default:
		return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, p.Data)
	}
}
