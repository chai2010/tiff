// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"fmt"
)

type TagType uint16

const (
	_                                  TagType = 0     // Type, Num
	TagType_NewSubfileType             TagType = 254   // LONG, 1
	TagType_SubfileType                TagType = 255   // SHORT, 1
	TagType_ImageWidth                 TagType = 256   // SHORT/LONG/LONG8, 1
	TagType_ImageLength                TagType = 257   // SHORT/LONG/LONG8, 1
	TagType_BitsPerSample              TagType = 258   // SHORT, SamplesPerPixel
	TagType_Compression                TagType = 259   // SHORT, 1
	TagType_PhotometricInterpretation  TagType = 262   // SHORT, 1
	TagType_Threshholding              TagType = 263   // SHORT, 1
	TagType_CellWidth                  TagType = 264   // SHORT, 1
	TagType_CellLenght                 TagType = 265   // SHORT, 1
	TagType_FillOrder                  TagType = 266   // SHORT, 1
	TagType_DocumentName               TagType = 269   // ASCII
	TagType_ImageDescription           TagType = 270   // ASCII
	TagType_Make                       TagType = 271   // ASCII
	TagType_Model                      TagType = 272   // ASCII
	TagType_StripOffsets               TagType = 273   // SHORT/LONG/LONG8, StripsPerImage
	TagType_Orientation                TagType = 274   // SHORT, 1
	TagType_SamplesPerPixel            TagType = 277   // SHORT, 1
	TagType_RowsPerStrip               TagType = 278   // SHORT/LONG/LONG8, 1
	TagType_StripByteCounts            TagType = 279   // SHORT/LONG/LONG8, StripsPerImage
	TagType_MinSampleValue             TagType = 280   // SHORT, SamplesPerPixel
	TagType_MaxSampleValue             TagType = 281   // SHORT, SamplesPerPixel
	TagType_XResolution                TagType = 282   // RATIONAL, 1
	TagType_YResolution                TagType = 283   // RATIONAL, 1
	TagType_PlanarConfiguration        TagType = 284   // SHORT
	TagType_PageName                   TagType = 285   // ASCII
	TagType_XPosition                  TagType = 286   // RATIONAL
	TagType_YPosition                  TagType = 287   // RATIONAL
	TagType_FreeOffsets                TagType = 288   // LONG/LONG8
	TagType_FreeByteCounts             TagType = 289   // LONG/LONG8
	TagType_GrayResponseUnit           TagType = 290   // SHORT, 1
	TagType_GrayResponseCurve          TagType = 291   // SHORT, 2**BitPerSample
	TagType_T4Options                  TagType = 292   // LONG, 1
	TagType_T6Options                  TagType = 293   // LONG, 1
	TagType_ResolutionUnit             TagType = 296   // SHORT, 1
	TagType_PageNumber                 TagType = 297   // SHORT, 2
	TagType_TransferFunction           TagType = 301   // SHORT, {1 or SamplesPerPixel}*2**BitPerSample
	TagType_Software                   TagType = 305   // ASCII
	TagType_DateTime                   TagType = 306   // ASCII, 20
	TagType_Artist                     TagType = 315   // ASCII
	TagType_HostComputer               TagType = 316   // ASCII
	TagType_Predictor                  TagType = 317   // SHORT, 1
	TagType_WhitePoint                 TagType = 318   // RATIONAL, 2
	TagType_PrimaryChromaticities      TagType = 319   // RATIONAL, 6
	TagType_ColorMap                   TagType = 320   // SHORT, 3*(2**BitPerSample)
	TagType_HalftoneHints              TagType = 321   // SHORT, 2
	TagType_TileWidth                  TagType = 322   // SHORT/LONG, 1
	TagType_TileLength                 TagType = 323   // SHORT/LONG, 1
	TagType_TileOffsets                TagType = 324   // LONG/LONG8, TilesPerImage
	TagType_TileByteCounts             TagType = 325   // SHORT/LONG, TilesPerImage
	TagType_InkSet                     TagType = 332   // SHORT, 1
	TagType_InkNames                   TagType = 333   // ASCII
	TagType_NumberOfInks               TagType = 334   // SHORT
	TagType_DotRange                   TagType = 336   // BYTE/SHORT, 2 or 2*NumberOfInks
	TagType_TargetPrinter              TagType = 337   // ASCII
	TagType_ExtraSamples               TagType = 338   // BYTE
	TagType_SampleFormat               TagType = 339   // SHORT, SamplesPerPixel
	TagType_SMinSampleValue            TagType = 340   // *, SamplesPerPixel
	TagType_SMaxSampleValue            TagType = 341   // *, SamplesPerPixel
	TagType_TransferRange              TagType = 342   // SHORT, 6
	TagType_JPEGProc                   TagType = 512   // SHORT, 1
	TagType_JPEGInterchangeFormat      TagType = 513   // LONG, 1
	TagType_JPEGInterchangeFormatLngth TagType = 514   // LONG, 1
	TagType_JPEGRestartInterval        TagType = 515   // SHORT, 1
	TagType_JPEGLosslessPredictors     TagType = 517   // SHORT, SamplesPerPixel
	TagType_JPEGPointTransforms        TagType = 518   // SHORT, SamplesPerPixel
	TagType_JPEGQTables                TagType = 519   // LONG, SamplesPerPixel
	TagType_JPEGDCTables               TagType = 520   // LONG, SamplesPerPixel
	TagType_JPEGACTables               TagType = 521   // LONG, SamplesPerPixel
	TagType_YCbCrCoefficients          TagType = 529   // RATIONAL, 3
	TagType_YCbCrSubSampling           TagType = 530   // SHORT, 2
	TagType_YCbCrPositioning           TagType = 531   // SHORT, 1
	TagType_ReferenceBlackWhite        TagType = 532   // LONG, 2*SamplesPerPixel
	TagType_Copyright                  TagType = 33432 // ASCII
)

const (
	TagType_GeoKeyDirectoryTag     TagType = 34735 // SHORT, >= 4
	TagType_GeoDoubleParamsTag     TagType = 34736 // DOUBLE
	TagType_GeoAsciiParamsTag      TagType = 34737 // ASCII
	TagType_ModelTiepointTag       TagType = 33922 // DOUBLE
	TagType_ModelPixelScaleTag     TagType = 33550 // DOUBLE
	TagType_ModelTransformationTag TagType = 34264 // DOUBLE, 16
	TagType_IntergraphMatrixTag    TagType = 33920 // DOUBLE, 17
)

func (t TagType) String() string {
	if name, ok := _TagTypeTable[t]; ok {
		return name
	}
	return fmt.Sprintf("TagType_Unknown(%d)", uint16(t))
}
