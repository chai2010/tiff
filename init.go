// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"image"
)

const (
	classicTiffLittleEnding = "II\x2A\x00"
	classicTiffBigEnding    = "MM\x00\x2A"
	bigTiffLittleEnding     = "II\x2B\x00"
	bigTiffBigEnding        = "MM\x00\x2B"
)

func init() {
	image.RegisterFormat("tiff", classicTiffLittleEnding, Decode, DecodeConfig)
	image.RegisterFormat("tiff", classicTiffBigEnding, Decode, DecodeConfig)
	image.RegisterFormat("tiff", bigTiffLittleEnding, Decode, DecodeConfig)
	image.RegisterFormat("tiff", bigTiffBigEnding, Decode, DecodeConfig)
}
