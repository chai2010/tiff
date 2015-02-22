// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
			var v uint64
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
		return string(runes)
	}
	return ""
}

func (p *IFDEntry) SetInts(...int64) {
	return
}
func (p *IFDEntry) SetUInts(...uint64) {
	return
}
func (p *IFDEntry) SetFloats(...float64) {
	return
}
func (p *IFDEntry) SetURationals(...[2]uint32) {
	return
}

func (p *IFDEntry) String() string {
	switch {
	case p.DataType.IsIntType():
		return fmt.Sprintf("%v: %v", p.Tag, p.GetInts())
	case p.DataType.IsFloatType():
		return fmt.Sprintf("%v: %v", p.Tag, p.GetFloats())
	case p.DataType.IsRationalType():
		return fmt.Sprintf("%v: %v", p.Tag, p.GetRationals())
	case p.DataType.IsStringType():
		return fmt.Sprintf("%v: %v", p.Tag, p.GetString())
	default:
		return fmt.Sprintf("%v: %v", p.Tag, p.Data)
	}
}
