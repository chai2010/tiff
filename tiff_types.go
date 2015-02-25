// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

/*
 * NB: In the comments below,
 *  - items marked with a + are obsoleted by revision 5.0,
 *  - items marked with a ! are introduced in revision 6.0.
 *  - items marked with a % are introduced post revision 6.0.
 *  - items marked with a $ are obsoleted by revision 6.0.
 *  - items marked with a & are introduced by Adobe DNG specification.
 */

const (
	ClassicTiffLittleEnding = "II\x2A\x00"
	ClassicTiffBigEnding    = "MM\x00\x2A"
	BigTiffLittleEnding     = "II\x2B\x00"
	BigTiffBigEnding        = "MM\x00\x2B"
)

type (
	TiffType                    uint16
	ImageType                   uint16
	DataType                    uint16
	TagType                     uint16
	TagValue_NewSubfileType     TagType
	TagValue_SubfileType        TagType
	TagValue_CompressionType    TagType
	TagValue_PhotometricType    TagType
	TagValue_PredictorType      TagType
	TagValue_ResolutionUnitType TagType
)

const (
	_                                 = 0  //
	TiffType_ClassicTIFF    TiffType  = 42 //
	TiffType_BigTIFF        TiffType  = 43 //
	_                                 = 0  //
	DataType_Nil            DataType  = 0  // placeholder, invalid
	DataType_Byte           DataType  = 1  // 8-bit unsigned integer
	DataType_ASCII          DataType  = 2  // 8-bit bytes w/ last byte null
	DataType_Short          DataType  = 3  // 16-bit unsigned integer
	DataType_Long           DataType  = 4  // 32-bit unsigned integer
	DataType_Rational       DataType  = 5  // 64-bit unsigned fraction
	DataType_SByte          DataType  = 6  // !8-bit signed integer
	DataType_Undefined      DataType  = 7  // !8-bit untyped data
	DataType_SShort         DataType  = 8  // !16-bit signed integer
	DataType_SLong          DataType  = 9  // !32-bit signed integer
	DataType_SRational      DataType  = 10 // !64-bit signed fraction
	DataType_Float          DataType  = 11 // !32-bit IEEE floating point
	DataType_Double         DataType  = 12 // !64-bit IEEE floating point
	DataType_IFD            DataType  = 13 // %32-bit unsigned integer (offset)
	_DataType_Unicode       DataType  = 14 // placeholder
	_DataType_Complex       DataType  = 15 // placeholder
	DataType_Long8          DataType  = 16 // BigTIFF 64-bit unsigned integer
	DataType_SLong8         DataType  = 17 // BigTIFF 64-bit signed integer
	DataType_IFD8           DataType  = 18 // BigTIFF 64-bit unsigned integer (offset)
	_                                 = 0  //
	ImageType_Nil           ImageType = 0  //
	ImageType_Bilevel       ImageType = 1  //
	ImageType_BilevelInvert ImageType = 2  //
	ImageType_Paletted      ImageType = 3  //
	ImageType_Gray          ImageType = 4  //
	ImageType_GrayInvert    ImageType = 5  //
	ImageType_RGB           ImageType = 6  //
	ImageType_RGBA          ImageType = 7  //
	ImageType_NRGBA         ImageType = 8  //
	_                                 = 0  //
)

const (
	_                                                                = 0     // Type(A/B/C/*), Num(1/*), Required, # comment
	TagType_NewSubfileType               TagType                     = 254   // LONG , 1, # Default=0. subfile data descriptor
	_                                                                = 0     //
	TagValue_NewSubfileType_ReducedImage TagValue_NewSubfileType     = 1     // # reduced resolution version
	TagValue_NewSubfileType_Page         TagValue_NewSubfileType     = 2     // # one page of many
	TagValue_NewSubfileType_Mask         TagValue_NewSubfileType     = 4     // # transparency mask
	_                                                                = 0     //
	TagType_SubfileType                  TagType                     = 255   // SHORT, 1, # kind of data in subfile
	_                                                                = 0     //
	TagValue_SubfileType_Image           TagValue_SubfileType        = 1     // # full resolution image data
	TagValue_SubfileType_ReducedImage    TagValue_SubfileType        = 2     // # reduced size image data
	TagValue_SubfileType_Page            TagValue_SubfileType        = 3     // # one page of many
	_                                                                = 0     //
	TagType_ImageWidth                   TagType                     = 256   // SHORT/LONG/LONG8, 1, # Required
	TagType_ImageLength                  TagType                     = 257   // SHORT/LONG/LONG8, 1, # Required
	TagType_BitsPerSample                TagType                     = 258   // SHORT, *, # Default=1. See SamplesPerPixel
	TagType_Compression                  TagType                     = 259   // SHORT, 1, # Default=1
	_                                                                = 0     //
	TagValue_CompressionType_Nil         TagValue_CompressionType    = 0     //
	TagValue_CompressionType_None        TagValue_CompressionType    = 1     //
	TagValue_CompressionType_CCITT       TagValue_CompressionType    = 2     //
	TagValue_CompressionType_G3          TagValue_CompressionType    = 3     // # Group 3 Fax.
	TagValue_CompressionType_G4          TagValue_CompressionType    = 4     // # Group 4 Fax.
	TagValue_CompressionType_LZW         TagValue_CompressionType    = 5     //
	TagValue_CompressionType_JPEGOld     TagValue_CompressionType    = 6     // # Superseded by cJPEG.
	TagValue_CompressionType_JPEG        TagValue_CompressionType    = 7     //
	TagValue_CompressionType_Deflate     TagValue_CompressionType    = 8     // # zlib compression.
	TagValue_CompressionType_PackBits    TagValue_CompressionType    = 32773 //
	TagValue_CompressionType_DeflateOld  TagValue_CompressionType    = 32946 // # Superseded by cDeflate.
	_                                                                = 0     //
	TagType_PhotometricInterpretation    TagType                     = 262   // SHORT, 1,
	_                                                                = 0     //
	TagValue_PhotometricType_WhiteIsZero TagValue_PhotometricType    = 0     //
	TagValue_PhotometricType_BlackIsZero TagValue_PhotometricType    = 1     //
	TagValue_PhotometricType_RGB         TagValue_PhotometricType    = 2     //
	TagValue_PhotometricType_Paletted    TagValue_PhotometricType    = 3     //
	TagValue_PhotometricType_TransMask   TagValue_PhotometricType    = 4     // # transparency mask
	TagValue_PhotometricType_CMYK        TagValue_PhotometricType    = 5     //
	TagValue_PhotometricType_YCbCr       TagValue_PhotometricType    = 6     //
	TagValue_PhotometricType_CIELab      TagValue_PhotometricType    = 8     //
	_                                                                = 0     //
	TagType_Threshholding                TagType                     = 263   // SHORT, 1, # Default=1
	TagType_CellWidth                    TagType                     = 264   // SHORT, 1,
	TagType_CellLenght                   TagType                     = 265   // SHORT, 1,
	TagType_FillOrder                    TagType                     = 266   // SHORT, 1, # Default=1
	TagType_DocumentName                 TagType                     = 269   // ASCII
	TagType_ImageDescription             TagType                     = 270   // ASCII
	TagType_Make                         TagType                     = 271   // ASCII
	TagType_Model                        TagType                     = 272   // ASCII
	TagType_StripOffsets                 TagType                     = 273   // SHORT/LONG/LONG8, *, # StripsPerImage
	TagType_Orientation                  TagType                     = 274   // SHORT, 1, # Default=1
	TagType_SamplesPerPixel              TagType                     = 277   // SHORT, 1, # Default=1
	TagType_RowsPerStrip                 TagType                     = 278   // SHORT/LONG/LONG8, 1,
	TagType_StripByteCounts              TagType                     = 279   // SHORT/LONG/LONG8, *, # StripsPerImage
	TagType_MinSampleValue               TagType                     = 280   // SHORT,    *, # Default=0
	TagType_MaxSampleValue               TagType                     = 281   // SHORT,    *, # Default=2^BitsPerSample-1
	TagType_XResolution                  TagType                     = 282   // RATIONAL, 1, # Required?
	TagType_YResolution                  TagType                     = 283   // RATIONAL, 1, # Required?
	TagType_PlanarConfiguration          TagType                     = 284   // SHORT,    1, # Defaule=1
	TagType_PageName                     TagType                     = 285   // ASCII
	TagType_XPosition                    TagType                     = 286   // RATIONAL,   1
	TagType_YPosition                    TagType                     = 287   // RATIONAL,   1
	TagType_FreeOffsets                  TagType                     = 288   // LONG/LONG8, *
	TagType_FreeByteCounts               TagType                     = 289   // LONG/LONG8, *
	TagType_GrayResponseUnit             TagType                     = 290   // SHORT, 1,
	TagType_GrayResponseCurve            TagType                     = 291   // SHORT, *, # 2**BitPerSample
	TagType_T4Options                    TagType                     = 292   // LONG,  1, # Default=0
	TagType_T6Options                    TagType                     = 293   // LONG,  1, # Default=0
	TagType_ResolutionUnit               TagType                     = 296   // SHORT, 1, # Default=2
	_                                                                = 0     //
	TagValue_ResolutionUnitType_None     TagValue_ResolutionUnitType = 1     //
	TagValue_ResolutionUnitType_PerInch  TagValue_ResolutionUnitType = 2     // # Dots per inch.
	TagValue_ResolutionUnitType_PerCM    TagValue_ResolutionUnitType = 3     // # Dots per centimeter.
	_                                                                = 0     //
	TagType_PageNumber                   TagType                     = 297   // SHORT, 2,
	TagType_TransferFunction             TagType                     = 301   // SHORT, *, # {1 or SamplesPerPixel}*2**BitPerSample
	TagType_Software                     TagType                     = 305   // ASCII
	TagType_DateTime                     TagType                     = 306   // ASCII, 20, # YYYY:MM:DD HH:MM:SS, include NULL
	TagType_Artist                       TagType                     = 315   // ASCII
	TagType_HostComputer                 TagType                     = 316   // ASCII
	TagType_Predictor                    TagType                     = 317   // SHORT, 1, # Default=1
	_                                                                = 0     //
	TagValue_PredictorType_None          TagValue_PredictorType      = 1     //
	TagValue_PredictorType_Horizontal    TagValue_PredictorType      = 2     //
	_                                                                = 0     //
	TagType_WhitePoint                   TagType                     = 318   // RATIONAL, 2
	TagType_PrimaryChromaticities        TagType                     = 319   // RATIONAL, 6
	TagType_ColorMap                     TagType                     = 320   // SHORT, *, # 3*(2**BitPerSample)
	TagType_HalftoneHints                TagType                     = 321   // SHORT, 2
	TagType_TileWidth                    TagType                     = 322   // SHORT/LONG, 1
	TagType_TileLength                   TagType                     = 323   // SHORT/LONG, 1
	TagType_TileOffsets                  TagType                     = 324   // LONG/LONG8, *, # TilesPerImage
	TagType_TileByteCounts               TagType                     = 325   // SHORT/LONG, *, # TilesPerImage
	TagType_SubIFD                       TagType                     = 330   // IFD,   *  # IFD pointer
	TagType_InkSet                       TagType                     = 332   // SHORT, 1, # Default=1
	TagType_InkNames                     TagType                     = 333   // ASCII
	TagType_NumberOfInks                 TagType                     = 334   // SHORT, 1, # Default=4
	TagType_DotRange                     TagType                     = 336   // BYTE/SHORT, # Default=[0,2^BitsPerSample-1]
	TagType_TargetPrinter                TagType                     = 337   // ASCII
	TagType_ExtraSamples                 TagType                     = 338   // BYTE,  1,
	TagType_SampleFormat                 TagType                     = 339   // SHORT, *, # SamplesPerPixel. Default=1
	TagType_SMinSampleValue              TagType                     = 340   // *,     *, # SamplesPerPixel, try double
	TagType_SMaxSampleValue              TagType                     = 341   // *,     *, # SamplesPerPixel, try double
	TagType_TransferRange                TagType                     = 342   // SHORT, 6,
	TagType_JPEGProc                     TagType                     = 512   // SHORT, 1,
	TagType_JPEGInterchangeFormat        TagType                     = 513   // LONG,  1,
	TagType_JPEGInterchangeFormatLngth   TagType                     = 514   // LONG,  1,
	TagType_JPEGRestartInterval          TagType                     = 515   // SHORT, 1,
	TagType_JPEGLosslessPredictors       TagType                     = 517   // SHORT, *, # SamplesPerPixel
	TagType_JPEGPointTransforms          TagType                     = 518   // SHORT, *, # SamplesPerPixel
	TagType_JPEGQTables                  TagType                     = 519   // LONG,  *, # SamplesPerPixel
	TagType_JPEGDCTables                 TagType                     = 520   // LONG,  *, # SamplesPerPixel
	TagType_JPEGACTables                 TagType                     = 521   // LONG,  *, # SamplesPerPixel
	TagType_YCbCrCoefficients            TagType                     = 529   // RATIONAL, 3
	TagType_YCbCrSubSampling             TagType                     = 530   // SHORT, 2, # Default=[2,2]
	TagType_YCbCrPositioning             TagType                     = 531   // SHORT, 1, # Default=1
	TagType_ReferenceBlackWhite          TagType                     = 532   // LONG , *, # 2*SamplesPerPixel
	TagType_Copyright                    TagType                     = 33432 // ASCII
	_                                                                = 0     //
	TagType_GeoKeyDirectoryTag           TagType                     = 34735 // SHORT, *, # >= 4
	TagType_GeoDoubleParamsTag           TagType                     = 34736 // DOUBLE
	TagType_GeoAsciiParamsTag            TagType                     = 34737 // ASCII
	TagType_ModelTiepointTag             TagType                     = 33922 // DOUBLE
	TagType_ModelPixelScaleTag           TagType                     = 33550 // DOUBLE
	TagType_ModelTransformationTag       TagType                     = 34264 // DOUBLE, 16
	TagType_IntergraphMatrixTag          TagType                     = 33920 // DOUBLE, 17
	_                                                                = 0     //
)
