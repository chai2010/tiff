// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"fmt"
)

type TagType uint16

const (
	TagType_ImageWidth                TagType = 256
	TagType_ImageLength               TagType = 257
	TagType_BitsPerSample             TagType = 258
	TagType_Compression               TagType = 259
	TagType_PhotometricInterpretation TagType = 262
)

const (
	TagType_StripOffsets    TagType = 273
	TagType_SamplesPerPixel TagType = 277
	TagType_RowsPerStrip    TagType = 278
	TagType_StripByteCounts TagType = 279
)

const (
	TagType_TileWidth      TagType = 322
	TagType_TileLength     TagType = 323
	TagType_TileOffsets    TagType = 324
	TagType_TileByteCounts TagType = 325
)

const (
	TagType_XResolution    TagType = 282
	TagType_YResolution    TagType = 283
	TagType_ResolutionUnit TagType = 296
)

const (
	TagType_Predictor    TagType = 317
	TagType_ColorMap     TagType = 320
	TagType_ExtraSamples TagType = 338
	TagType_SampleFormat TagType = 339
)

func (t TagType) String() string {
	if name, ok := _TagTypeTable[t]; ok {
		return name
	}
	return fmt.Sprintf("TagType_Unknown(%d)", uint16(t))
}
