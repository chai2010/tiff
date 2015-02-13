// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

// A tiff image file contains one or more images. The metadata
// of each image is contained in an Image File Directory (IFD),
// which contains entries of 12 bytes each and is described
// on page 14-16 of the specification. An IFD entry consists of
//
//  - a tag, which describes the signification of the entry,
//  - the data type and length of the entry,
//  - the data itself or a pointer to it if it is more than 4 bytes.
//
// The presence of a length means that each IFD is effectively an array.

const (
	ifdLen  = 12 // Length of an IFD entry in bytes.
	ifd8Len = 20
)

// imageMode represents the mode of the image.
type imageMode int

const (
	mBilevel imageMode = iota
	mPaletted
	mGray
	mGrayInvert
	mRGB
	mRGBA
	mNRGBA
)
