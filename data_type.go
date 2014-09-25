// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"fmt"
)

type DataType uint16

const (
	DataType_Nil       DataType = iota //  0, invalid
	DataType_Byte                      //  1
	DataType_ASCII                     //  2
	DataType_Short                     //  3
	DataType_Long                      //  4
	DataType_Rational                  //  5
	DataType_SByte                     //  6
	DataType_Undefined                 //  7
	DataType_SShort                    //  8
	DataType_SLong                     //  9
	DataType_SRational                 // 10
	DataType_Float                     // 11
	DataType_Double                    // 12
	DataType_IFD                       // 13
	DataType_Unicode                   // 14
	DataType_Complex                   // 15
	DataType_Long8                     // 16
	DataType_SLong8                    // 17
	DataType_IFD8                      // 18
)

func (d DataType) Valid() bool {
	if d <= DataType_Nil || d > DataType_IFD8 {
		return false
	}
	return true
}

func (d DataType) Size() int {
	switch d {
	case DataType_Nil:
		return 0
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
		return 0
	case DataType_Unicode:
		return 1
	case DataType_Complex:
		return 8
	case DataType_Long8:
		return 8
	case DataType_SLong8:
		return 8
	case DataType_IFD8:
		return 0
	default:
		return 0
	}
}

func (t DataType) String() string {
	if name, ok := _DataTypeTable[t]; ok {
		return name
	}
	return fmt.Sprintf("DataType_Unknown(%d)", uint16(t))
}
