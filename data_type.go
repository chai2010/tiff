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

func (d DataType) Name() string {
	if !d.Valid() {
	}
	switch d {
	case DataType_Nil:
		return "DataType_Nil"
	case DataType_Byte:
		return "DataType_Byte"
	case DataType_ASCII:
		return "DataType_ASCII"
	case DataType_Short:
		return "DataType_Short"
	case DataType_Long:
		return "DataType_Long"
	case DataType_Rational:
		return "DataType_Rational"
	case DataType_SByte:
		return "DataType_SByte"
	case DataType_Undefined:
		return "DataType_Undefined"
	case DataType_SShort:
		return "DataType_SShort"
	case DataType_SLong:
		return "DataType_SLong"
	case DataType_SRational:
		return "DataType_SRational"
	case DataType_Float:
		return "DataType_Float"
	case DataType_Double:
		return "DataType_Double"
	case DataType_IFD:
		return "DataType_IFD"
	case DataType_Unicode:
		return "DataType_Unicode"
	case DataType_Complex:
		return "DataType_Complex"
	case DataType_Long8:
		return "DataType_Long8"
	case DataType_SLong8:
		return "DataType_SLong8"
	case DataType_IFD8:
		return "DataType_IFD8"
	default:
		return fmt.Sprintf("DataType_Unknow(%d)", uint16(d))
	}
}

func (d DataType) String() string {
	return d.Name()
}
