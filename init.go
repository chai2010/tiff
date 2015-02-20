// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"image"
)

const (
	ClassicTiffLittleEnding = "II\x2A\x00"
	ClassicTiffBigEnding    = "MM\x00\x2A"
	BigTiffLittleEnding     = "II\x2B\x00"
	BigTiffBigEnding        = "MM\x00\x2B"
)

func init() {
	image.RegisterFormat("tiff", ClassicTiffLittleEnding, Decode, DecodeConfig)
	image.RegisterFormat("tiff", ClassicTiffBigEnding, Decode, DecodeConfig)
	image.RegisterFormat("tiff", BigTiffLittleEnding, Decode, DecodeConfig)
	image.RegisterFormat("tiff", BigTiffBigEnding, Decode, DecodeConfig)
}
