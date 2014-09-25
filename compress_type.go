// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"fmt"
)

type CompressType uint16

const (
	CompressType_None       CompressType = 1     //
	CompressType_CCITT      CompressType = 2     //
	CompressType_G3         CompressType = 3     // Group 3 Fax.
	CompressType_G4         CompressType = 4     // Group 4 Fax.
	CompressType_LZW        CompressType = 5     //
	CompressType_JPEGOld    CompressType = 6     // Superseded by cJPEG.
	CompressType_JPEG       CompressType = 7     //
	CompressType_Deflate    CompressType = 8     // zlib compression.
	CompressType_PackBits   CompressType = 32773 //
	CompressType_DeflateOld CompressType = 32946 // Superseded by cDeflate.
)

func (t CompressType) String() string {
	if name, ok := _CompressTypeTable[t]; ok {
		return name
	}
	return fmt.Sprintf("CompressType_Unknown(%d)", uint16(t))
}
