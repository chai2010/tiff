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
	}
	return nil
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

func (p *IFDEntry) SetBytes(data []byte) {
	return
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
