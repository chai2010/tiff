// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unicode/utf16"
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

func NewIFDEntry(hdr *Header, tag TagType, dataType DataType, data ...interface{}) *IFDEntry {
	return nil
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
	case DataType_Long:
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
	case DataType_Long8:
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
			floats[i] = float64(rats[i][0]) / float64(rats[i][0])
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
	case DataType_Unicode:
		r := bytes.NewReader(p.Data)
		runes := make([]rune, p.Count)
		for i := 0; i < p.Count; i++ {
			var v uint16
			if err := binary.Read(r, p.Header.ByteOrder, &v); err != nil {
				return ""
			}
			runes[i] = rune(v)
		}
		for i := 0; i < p.Count; i++ {
			if runes[i] == 0 {
				runes = runes[:i]
				break
			}
		}
		return string(runes)
	}
	return ""
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

func (p *IFDEntry) SetInts(value ...int64) {
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
		if err := binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
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
		if err := binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
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
		if err := binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
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
		if err := binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
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
		if err := binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
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
		if err := binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
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
		if err := binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
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
		if err := binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(tmp)
	}
	return
}

func (p *IFDEntry) SetFloats(value ...float64) {
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
		if err := binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(tmp)
	case DataType_Double:
		var buf bytes.Buffer
		if err := binary.Write(&buf, p.Header.ByteOrder, value); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(value)
	}
	return
}

func (p *IFDEntry) SetRationals(value ...[2]int64) {
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
		if err := binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
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
		if err := binary.Write(&buf, p.Header.ByteOrder, tmp); err != nil {
			return
		}
		p.Data = buf.Bytes()
		p.Count = len(tmp)
	}
	return
}

func (p *IFDEntry) SetString(value string) {
	if p.DataType == DataType_Nil {
		p.DataType = DataType_ASCII
	}
	switch p.DataType {
	case DataType_ASCII:
		p.Data = make([]byte, len(value)+1)
		copy(p.Data, []byte(value))
		p.Data[len(value)] = 0 // +NULL
		p.Count = len(p.Data) + 1
	case DataType_Unicode:
		u16Data := utf16.Encode([]rune(value))
		var buf bytes.Buffer
		if err := binary.Write(&buf, p.Header.ByteOrder, u16Data); err != nil {
			return
		}
		binary.Write(&buf, p.Header.ByteOrder, uint16(0)) // +NULL
		p.Data = buf.Bytes()
		p.Count = len(u16Data) + 1
	}
	return
}

func (p *IFDEntry) SetUndefined(value interface{}) {
	if p.DataType == DataType_Nil {
		p.DataType = DataType_Undefined
	}
	if p.DataType != DataType_Undefined {
		return
	}
	var buf bytes.Buffer
	if err := binary.Write(&buf, p.Header.ByteOrder, value); err != nil {
		return
	}
	p.Data = buf.Bytes()
	p.Count = len(p.Data)
	return
}

func (p *IFDEntry) String() string {
	switch {
	case p.DataType.IsIntType():
		return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, p.GetInts())
	case p.DataType.IsFloatType():
		return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, p.GetFloats())
	case p.DataType.IsRationalType():
		return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, p.GetRationals())
	case p.DataType.IsStringType():
		return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, p.GetString())
	default:
		return fmt.Sprintf("%v(%v): %v", p.Tag, p.DataType, p.Data)
	}
}
