// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

func (p TiffType) Valid() bool {
	return p == TiffType_ClassicTIFF || p == TiffType_BigTIFF
}

func (d DataType) Valid() bool {
	if d <= DataType_Nil || d > DataType_IFD8 {
		return false
	}
	if d == _DataType_Unicode {
		return false
	}
	if d == _DataType_Complex {
		return false
	}
	return true
}

func (d DataType) IsIntType() bool {
	switch d {
	case DataType_Byte, DataType_Short, DataType_Long:
		return true
	case DataType_SByte, DataType_SShort, DataType_SLong:
		return true
	case DataType_IFD, DataType_IFD8:
		return true
	}
	return false
}
func (d DataType) IsFloatType() bool {
	switch d {
	case DataType_Float, DataType_Double:
		return true
	}
	return false
}
func (d DataType) IsRationalType() bool {
	switch d {
	case DataType_Rational, DataType_SRational:
		return true
	}
	return false
}
func (d DataType) IsStringType() bool {
	switch d {
	case DataType_ASCII:
		return true
	}
	return false
}

func (d DataType) ByteSize() int {
	switch d {
	case DataType_Byte:
		return 1
	case DataType_ASCII:
		return 1
	case DataType_Short:
		return 2
	case DataType_Long:
		return 4
	case DataType_Rational:
		return 8
	case DataType_SByte:
		return 1
	case DataType_Undefined:
		return 1
	case DataType_SShort:
		return 2
	case DataType_SLong:
		return 4
	case DataType_SRational:
		return 8
	case DataType_Float:
		return 4
	case DataType_Double:
		return 8
	case DataType_IFD:
		return 4 // LONG
	case DataType_Long8:
		return 8
	case DataType_SLong8:
		return 8
	case DataType_IFD8:
		return 8 // LONG8
	}
	return 0
}
