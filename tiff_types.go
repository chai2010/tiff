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

type TiffType uint16

const (
	TiffType_ClassicTIFF TiffType = 42
	TiffType_BigTIFF     TiffType = 43
)

type ImageType uint16

const (
	ImageType_Nil ImageType = iota
	ImageType_Bilevel
	ImageType_BilevelInvert
	ImageType_Paletted
	ImageType_Gray
	ImageType_GrayInvert
	ImageType_RGB
	ImageType_RGBA
	ImageType_NRGBA
)

type DataType uint16

const (
	DataType_Nil       DataType = 0  // placeholder, invalid
	DataType_Byte      DataType = 1  // 8-bit unsigned integer
	DataType_ASCII     DataType = 2  // 8-bit bytes w/ last byte null
	DataType_Short     DataType = 3  // 16-bit unsigned integer
	DataType_Long      DataType = 4  // 32-bit unsigned integer
	DataType_Rational  DataType = 5  // 64-bit unsigned fraction
	DataType_SByte     DataType = 6  // !8-bit signed integer
	DataType_Undefined DataType = 7  // !8-bit untyped data
	DataType_SShort    DataType = 8  // !16-bit signed integer
	DataType_SLong     DataType = 9  // !32-bit signed integer
	DataType_SRational DataType = 10 // !64-bit signed fraction
	DataType_Float     DataType = 11 // !32-bit IEEE floating point
	DataType_Double    DataType = 12 // !64-bit IEEE floating point
	DataType_IFD       DataType = 13 // %32-bit unsigned integer (offset)
	DataType_Long8     DataType = 16 // BigTIFF 64-bit unsigned integer
	DataType_SLong8    DataType = 17 // BigTIFF 64-bit signed integer
	DataType_IFD8      DataType = 18 // BigTIFF 64-bit unsigned integer (offset)
)

type (
	TagType                     uint16
	TagValue_NewSubfileType     TagType
	TagValue_SubfileType        TagType
	TagValue_CompressionType    TagType
	TagValue_PhotometricType    TagType
	TagValue_PredictorType      TagType
	TagValue_ResolutionUnitType TagType
	TagValue_SampleFormatType   TagType
)

const (
	_                                                                     = 0     // Type(A/B/C/*), Num(1/*), Required, # comment
	TagType_NewSubfileType                    TagType                     = 254   // LONG , 1, # Default=0. subfile data descriptor
	_                                                                     = 0     //
	TagValue_NewSubfileType_Nil               TagValue_NewSubfileType     = 0     //
	TagValue_NewSubfileType_Reduced           TagValue_NewSubfileType     = 1     // # bit0, reduced resolution version
	TagValue_NewSubfileType_Page              TagValue_NewSubfileType     = 2     // # bit1, one page of many
	TagValue_NewSubfileType_Reduced_Page      TagValue_NewSubfileType     = 3     //
	TagValue_NewSubfileType_Mask              TagValue_NewSubfileType     = 4     // # bit2, transparency mask
	TagValue_NewSubfileType_Reduced_Mask      TagValue_NewSubfileType     = 5     //
	TagValue_NewSubfileType_Page_Mask         TagValue_NewSubfileType     = 6     //
	TagValue_NewSubfileType_Reduced_Page_Mask TagValue_NewSubfileType     = 7     //
	_                                                                     = 0     //
	TagType_SubfileType                       TagType                     = 255   // SHORT, 1, # kind of data in subfile
	_                                                                     = 0     //
	TagValue_SubfileType_Image                TagValue_SubfileType        = 1     // # full resolution image data
	TagValue_SubfileType_ReducedImage         TagValue_SubfileType        = 2     // # reduced size image data
	TagValue_SubfileType_Page                 TagValue_SubfileType        = 3     // # one page of many
	_                                                                     = 0     //
	TagType_ImageWidth                        TagType                     = 256   // SHORT/LONG/LONG8, 1, # Required
	TagType_ImageLength                       TagType                     = 257   // SHORT/LONG/LONG8, 1, # Required
	TagType_BitsPerSample                     TagType                     = 258   // SHORT, *, # Default=1. See SamplesPerPixel
	TagType_Compression                       TagType                     = 259   // SHORT, 1, # Default=1
	_                                                                     = 0     //
	TagValue_CompressionType_Nil              TagValue_CompressionType    = 0     //
	TagValue_CompressionType_None             TagValue_CompressionType    = 1     //
	TagValue_CompressionType_CCITT            TagValue_CompressionType    = 2     //
	TagValue_CompressionType_G3               TagValue_CompressionType    = 3     // # Group 3 Fax.
	TagValue_CompressionType_G4               TagValue_CompressionType    = 4     // # Group 4 Fax.
	TagValue_CompressionType_LZW              TagValue_CompressionType    = 5     //
	TagValue_CompressionType_JPEGOld          TagValue_CompressionType    = 6     // # Superseded by cJPEG.
	TagValue_CompressionType_JPEG             TagValue_CompressionType    = 7     //
	TagValue_CompressionType_Deflate          TagValue_CompressionType    = 8     // # zlib compression.
	TagValue_CompressionType_PackBits         TagValue_CompressionType    = 32773 //
	TagValue_CompressionType_DeflateOld       TagValue_CompressionType    = 32946 // # Superseded by cDeflate.
	_                                                                     = 0     //
	TagType_PhotometricInterpretation         TagType                     = 262   // SHORT, 1,
	_                                                                     = 0     //
	TagValue_PhotometricType_WhiteIsZero      TagValue_PhotometricType    = 0     //
	TagValue_PhotometricType_BlackIsZero      TagValue_PhotometricType    = 1     //
	TagValue_PhotometricType_RGB              TagValue_PhotometricType    = 2     //
	TagValue_PhotometricType_Paletted         TagValue_PhotometricType    = 3     //
	TagValue_PhotometricType_TransMask        TagValue_PhotometricType    = 4     // # transparency mask
	TagValue_PhotometricType_CMYK             TagValue_PhotometricType    = 5     //
	TagValue_PhotometricType_YCbCr            TagValue_PhotometricType    = 6     //
	TagValue_PhotometricType_CIELab           TagValue_PhotometricType    = 8     //
	_                                                                     = 0     //
	TagType_Threshholding                     TagType                     = 263   // SHORT, 1, # Default=1
	TagType_CellWidth                         TagType                     = 264   // SHORT, 1,
	TagType_CellLenght                        TagType                     = 265   // SHORT, 1,
	TagType_FillOrder                         TagType                     = 266   // SHORT, 1, # Default=1
	TagType_DocumentName                      TagType                     = 269   // ASCII
	TagType_ImageDescription                  TagType                     = 270   // ASCII
	TagType_Make                              TagType                     = 271   // ASCII
	TagType_Model                             TagType                     = 272   // ASCII
	TagType_StripOffsets                      TagType                     = 273   // SHORT/LONG/LONG8, *, # StripsPerImage
	TagType_Orientation                       TagType                     = 274   // SHORT, 1, # Default=1
	TagType_SamplesPerPixel                   TagType                     = 277   // SHORT, 1, # Default=1
	TagType_RowsPerStrip                      TagType                     = 278   // SHORT/LONG/LONG8, 1,
	TagType_StripByteCounts                   TagType                     = 279   // SHORT/LONG/LONG8, *, # StripsPerImage
	TagType_MinSampleValue                    TagType                     = 280   // SHORT,    *, # Default=0
	TagType_MaxSampleValue                    TagType                     = 281   // SHORT,    *, # Default=2^BitsPerSample-1
	TagType_XResolution                       TagType                     = 282   // RATIONAL, 1, # Required?
	TagType_YResolution                       TagType                     = 283   // RATIONAL, 1, # Required?
	TagType_PlanarConfiguration               TagType                     = 284   // SHORT,    1, # Defaule=1
	TagType_PageName                          TagType                     = 285   // ASCII
	TagType_XPosition                         TagType                     = 286   // RATIONAL,   1
	TagType_YPosition                         TagType                     = 287   // RATIONAL,   1
	TagType_FreeOffsets                       TagType                     = 288   // LONG/LONG8, *
	TagType_FreeByteCounts                    TagType                     = 289   // LONG/LONG8, *
	TagType_GrayResponseUnit                  TagType                     = 290   // SHORT, 1,
	TagType_GrayResponseCurve                 TagType                     = 291   // SHORT, *, # 2**BitPerSample
	TagType_T4Options                         TagType                     = 292   // LONG,  1, # Default=0
	TagType_T6Options                         TagType                     = 293   // LONG,  1, # Default=0
	TagType_ResolutionUnit                    TagType                     = 296   // SHORT, 1, # Default=2
	_                                                                     = 0     //
	TagValue_ResolutionUnitType_None          TagValue_ResolutionUnitType = 1     //
	TagValue_ResolutionUnitType_PerInch       TagValue_ResolutionUnitType = 2     // # Dots per inch.
	TagValue_ResolutionUnitType_PerCM         TagValue_ResolutionUnitType = 3     // # Dots per centimeter.
	_                                                                     = 0     //
	TagType_PageNumber                        TagType                     = 297   // SHORT, 2,
	TagType_TransferFunction                  TagType                     = 301   // SHORT, *, # {1 or SamplesPerPixel}*2**BitPerSample
	TagType_Software                          TagType                     = 305   // ASCII
	TagType_DateTime                          TagType                     = 306   // ASCII, 20, # YYYY:MM:DD HH:MM:SS, include NULL
	TagType_Artist                            TagType                     = 315   // ASCII
	TagType_HostComputer                      TagType                     = 316   // ASCII
	TagType_Predictor                         TagType                     = 317   // SHORT, 1, # Default=1
	_                                                                     = 0     //
	TagValue_PredictorType_None               TagValue_PredictorType      = 1     //
	TagValue_PredictorType_Horizontal         TagValue_PredictorType      = 2     //
	_                                                                     = 0     //
	TagType_WhitePoint                        TagType                     = 318   // RATIONAL, 2
	TagType_PrimaryChromaticities             TagType                     = 319   // RATIONAL, 6
	TagType_ColorMap                          TagType                     = 320   // SHORT, *, # 3*(2**BitPerSample)
	TagType_HalftoneHints                     TagType                     = 321   // SHORT, 2
	TagType_TileWidth                         TagType                     = 322   // SHORT/LONG, 1
	TagType_TileLength                        TagType                     = 323   // SHORT/LONG, 1
	TagType_TileOffsets                       TagType                     = 324   // LONG/LONG8, *, # TilesPerImage
	TagType_TileByteCounts                    TagType                     = 325   // SHORT/LONG, *, # TilesPerImage
	TagType_BadFaxLines                       TagType                     = 326   // ingore # Used in the TIFF-F standard, denotes the number of 'bad' scan lines encountered by the facsimile device.
	TagType_CleanFaxData                      TagType                     = 327   // ingore # Used in the TIFF-F standard, indicates if 'bad' lines encountered during reception are stored in the data, or if 'bad' lines have been replaced by the receiver.
	TagType_ConsecutiveBadFaxLines            TagType                     = 328   // ingore # Used in the TIFF-F standard, denotes the maximum number of consecutive 'bad' scanlines received.
	TagType_SubIFD                            TagType                     = 330   // IFD,   *  # IFD pointer
	TagType_InkSet                            TagType                     = 332   // SHORT, 1, # Default=1
	TagType_InkNames                          TagType                     = 333   // ASCII
	TagType_NumberOfInks                      TagType                     = 334   // SHORT, 1, # Default=4
	TagType_DotRange                          TagType                     = 336   // BYTE/SHORT, # Default=[0,2^BitsPerSample-1]
	TagType_TargetPrinter                     TagType                     = 337   // ASCII
	TagType_ExtraSamples                      TagType                     = 338   // BYTE,  1,
	TagType_SampleFormat                      TagType                     = 339   // SHORT, *, # SamplesPerPixel. Default=1
	_                                                                     = 0     //
	TagValue_SampleFormatType_Uint            TagValue_SampleFormatType   = 1     //
	TagValue_SampleFormatType_TwoInt          TagValue_SampleFormatType   = 2     //
	TagValue_SampleFormatType_Float           TagValue_SampleFormatType   = 3     //
	TagValue_SampleFormatType_Undefined       TagValue_SampleFormatType   = 4     //
	_                                                                     = 0     //
	TagType_SMinSampleValue                   TagType                     = 340   // *,     *, # SamplesPerPixel, try double
	TagType_SMaxSampleValue                   TagType                     = 341   // *,     *, # SamplesPerPixel, try double
	TagType_TransferRange                     TagType                     = 342   // SHORT, 6,
	TagType_ClipPath                          TagType                     = 343   // ingore # Mirrors the essentials of PostScript's path creation functionality.
	TagType_XClipPathUnits                    TagType                     = 344   // ingore # The number of units that span the width of the image, in terms of integer ClipPath coordinates.
	TagType_YClipPathUnits                    TagType                     = 345   // ingore # The number of units that span the height of the image, in terms of integer ClipPath coordinates.
	TagType_Indexed                           TagType                     = 346   // ingore # Aims to broaden the support for indexed images to include support for any color space.
	TagType_JPEGTables                        TagType                     = 347   // ingore # JPEG quantization and/or Huffman tables.
	TagType_OPIProxy                          TagType                     = 351   // ingore # OPI-related.
	TagType_GlobalParametersIFD               TagType                     = 400   // ingore # Used in the TIFF-FX standard to point to an IFD containing tags that are globally applicable to the complete TIFF file.
	TagType_ProfileType                       TagType                     = 401   // ingore # Used in the TIFF-FX standard, denotes the type of data stored in this file or IFD.
	TagType_FaxProfile                        TagType                     = 402   // ingore # Used in the TIFF-FX standard, denotes the 'profile' that applies to this file.
	TagType_CodingMethods                     TagType                     = 403   // ingore # Used in the TIFF-FX standard, indicates which coding methods are used in the file.
	TagType_VersionYear                       TagType                     = 404   // ingore # Used in the TIFF-FX standard, denotes the year of the standard specified by the FaxProfile field.
	TagType_ModeNumber                        TagType                     = 405   // ingore # Used in the TIFF-FX standard, denotes the mode of the standard specified by the FaxProfile field.
	TagType_Decode                            TagType                     = 433   // ingore # Used in the TIFF-F and TIFF-FX standards, holds information about the ITULAB (PhotometricInterpretation = 10) encoding.
	TagType_DefaultImageColor                 TagType                     = 434   // ingore # Defined in the Mixed Raster Content part of RFC 2301, is the default color needed in areas where no image is available.
	TagType_JPEGProc                          TagType                     = 512   // SHORT, 1,
	TagType_JPEGInterchangeFormat             TagType                     = 513   // LONG,  1,
	TagType_JPEGInterchangeFormatLength       TagType                     = 514   // LONG,  1,
	TagType_JPEGRestartInterval               TagType                     = 515   // SHORT, 1,
	TagType_JPEGLosslessPredictors            TagType                     = 517   // SHORT, *, # SamplesPerPixel
	TagType_JPEGPointTransforms               TagType                     = 518   // SHORT, *, # SamplesPerPixel
	TagType_JPEGQTables                       TagType                     = 519   // LONG,  *, # SamplesPerPixel
	TagType_JPEGDCTables                      TagType                     = 520   // LONG,  *, # SamplesPerPixel
	TagType_JPEGACTables                      TagType                     = 521   // LONG,  *, # SamplesPerPixel
	TagType_YCbCrCoefficients                 TagType                     = 529   // RATIONAL, 3
	TagType_YCbCrSubSampling                  TagType                     = 530   // SHORT, 2, # Default=[2,2]
	TagType_YCbCrPositioning                  TagType                     = 531   // SHORT, 1, # Default=1
	TagType_ReferenceBlackWhite               TagType                     = 532   // LONG , *, # 2*SamplesPerPixel
	TagType_StripRowCounts                    TagType                     = 559   // ingore # Defined in the Mixed Raster Content part of RFC 2301, used to replace RowsPerStrip for IFDs with variable-sized strips.
	TagType_XMP                               TagType                     = 700   // ingore # XML packet containing XMP metadata
	TagType_ImageID                           TagType                     = 32781 // ingore # OPI-related.
	TagType_ImageLayer                        TagType                     = 34732 // ingore # Defined in the Mixed Raster Content part of RFC 2301, used to denote the particular function of this Image in the mixed raster scheme.
	TagType_Copyright                         TagType                     = 33432 // ASCII
	TagType_WangAnnotation                    TagType                     = 32932 // ingore # Annotation data, as used in 'Imaging for Windows'.
	TagType_MDFileTag                         TagType                     = 33445 // ingore # Specifies the pixel data format encoding in the Molecular Dynamics GEL file format.
	TagType_MDScalePixel                      TagType                     = 33446 // ingore # Specifies a scale factor in the Molecular Dynamics GEL file format.
	TagType_MDColorTable                      TagType                     = 33447 // ingore # Used to specify the conversion from 16bit to 8bit in the Molecular Dynamics GEL file format.
	TagType_MDLabName                         TagType                     = 33448 // ingore # Name of the lab that scanned this file, as used in the Molecular Dynamics GEL file format.
	TagType_MDSampleInfo                      TagType                     = 33449 // ingore # Information about the sample, as used in the Molecular Dynamics GEL file format.
	TagType_MDPrepDate                        TagType                     = 33450 // ingore # Date the sample was prepared, as used in the Molecular Dynamics GEL file format.
	TagType_MDPrepTime                        TagType                     = 33451 // ingore # Time the sample was prepared, as used in the Molecular Dynamics GEL file format.
	TagType_MDFileUnits                       TagType                     = 33452 // ingore # Units for data in this file, as used in the Molecular Dynamics GEL file format.
	TagType_ModelPixelScaleTag                TagType                     = 33550 // DOUBLE # Used in interchangeable GeoTIFF files.
	TagType_IPTC                              TagType                     = 33723 // ingore # IPTC (International Press Telecommunications Council) metadata.
	TagType_INGRPacketDataTag                 TagType                     = 33918 // ingore # Intergraph Application specific storage.
	TagType_INGRFlagRegisters                 TagType                     = 33919 // ingore # Intergraph Application specific flags.
	TagType_IrasBTransformationMatrix         TagType                     = 33920 // DOUBLE, 17 # Originally part of Intergraph's GeoTIFF tags, but likely understood by IrasB only.
	TagType_ModelTiepointTag                  TagType                     = 33922 // DOUBLE # Originally part of Intergraph's GeoTIFF tags, but now used in interchangeable GeoTIFF files.
	TagType_ModelTransformationTag            TagType                     = 34264 // DOUBLE, 16 # Used in interchangeable GeoTIFF files.
	TagType_Photoshop                         TagType                     = 34377 // ingore # Collection of Photoshop 'Image Resource Blocks'.
	TagType_ExifIFD                           TagType                     = 34665 // IFD    # A pointer to the Exif IFD.
	TagType_ICCProfile                        TagType                     = 34675 // ingore # ICC profile data.
	TagType_GeoKeyDirectoryTag                TagType                     = 34735 // SHORT, *, # >= 4
	TagType_GeoDoubleParamsTag                TagType                     = 34736 // DOUBLE
	TagType_GeoAsciiParamsTag                 TagType                     = 34737 // ASCII
	TagType_GPSIFD                            TagType                     = 34853 // IFD    # A pointer to the Exif-related GPS Info IFD.
	TagType_HylaFAXFaxRecvParams              TagType                     = 34908 // ingore # Used by HylaFAX.
	TagType_HylaFAXFaxSubAddress              TagType                     = 34909 // ingore # Used by HylaFAX.
	TagType_HylaFAXFaxRecvTime                TagType                     = 34910 // ingore # Used by HylaFAX.
	TagType_ImageSourceData                   TagType                     = 37724 // ingore # Used by Adobe Photoshop.
	TagType_InteroperabilityIFD               TagType                     = 40965 // IFD    # A pointer to the Exif-related Interoperability IFD.
	TagType_GDAL_METADATA                     TagType                     = 42112 // ingore # Used by the GDAL library, holds an XML list of name=value 'metadata' values about the image as a whole, and about specific samples.
	TagType_GDAL_NODATA                       TagType                     = 42113 // ingore # Used by the GDAL library, contains an ASCII encoded nodata or background pixel value.
	TagType_OceScanjobDescription             TagType                     = 50215 // ingore # Used in the Oce scanning process.
	TagType_OceApplicationSelector            TagType                     = 50216 // ingore # Used in the Oce scanning process.
	TagType_OceIdentificationNumber           TagType                     = 50217 // ingore # Used in the Oce scanning process.
	TagType_OceImageLogicCharacteristics      TagType                     = 50218 // ingore # Used in the Oce scanning process.
	TagType_DNGVersion                        TagType                     = 50706 // ingore # Used in IFD 0 of DNG files.
	TagType_DNGBackwardVersion                TagType                     = 50707 // ingore # Used in IFD 0 of DNG files.
	TagType_UniqueCameraModel                 TagType                     = 50708 // ingore # Used in IFD 0 of DNG files.
	TagType_LocalizedCameraModel              TagType                     = 50709 // ingore # Used in IFD 0 of DNG files.
	TagType_CFAPlaneColor                     TagType                     = 50710 // ingore # Used in Raw IFD of DNG files.
	TagType_CFALayout                         TagType                     = 50711 // ingore # Used in Raw IFD of DNG files.
	TagType_LinearizationTable                TagType                     = 50712 // ingore # Used in Raw IFD of DNG files.
	TagType_BlackLevelRepeatDim               TagType                     = 50713 // ingore # Used in Raw IFD of DNG files.
	TagType_BlackLevel                        TagType                     = 50714 // ingore # Used in Raw IFD of DNG files.
	TagType_BlackLevelDeltaH                  TagType                     = 50715 // ingore # Used in Raw IFD of DNG files.
	TagType_BlackLevelDeltaV                  TagType                     = 50716 // ingore # Used in Raw IFD of DNG files.
	TagType_WhiteLevel                        TagType                     = 50717 // ingore # Used in Raw IFD of DNG files.
	TagType_DefaultScale                      TagType                     = 50718 // ingore # Used in Raw IFD of DNG files.
	TagType_DefaultCropOrigin                 TagType                     = 50719 // ingore # Used in Raw IFD of DNG files.
	TagType_DefaultCropSize                   TagType                     = 50720 // ingore # Used in Raw IFD of DNG files.
	TagType_ColorMatrix1                      TagType                     = 50721 // ingore # Used in IFD 0 of DNG files.
	TagType_ColorMatrix2                      TagType                     = 50722 // ingore # Used in IFD 0 of DNG files.
	TagType_CameraCalibration1                TagType                     = 50723 // ingore # Used in IFD 0 of DNG files.
	TagType_CameraCalibration2                TagType                     = 50724 // ingore # Used in IFD 0 of DNG files.
	TagType_ReductionMatrix1                  TagType                     = 50725 // ingore # Used in IFD 0 of DNG files.
	TagType_ReductionMatrix2                  TagType                     = 50726 // ingore # Used in IFD 0 of DNG files.
	TagType_AnalogBalance                     TagType                     = 50727 // ingore # Used in IFD 0 of DNG files.
	TagType_AsShotNeutral                     TagType                     = 50728 // ingore # Used in IFD 0 of DNG files.
	TagType_AsShotWhiteXY                     TagType                     = 50729 // ingore # Used in IFD 0 of DNG files.
	TagType_BaselineExposure                  TagType                     = 50730 // ingore # Used in IFD 0 of DNG files.
	TagType_BaselineNoise                     TagType                     = 50731 // ingore # Used in IFD 0 of DNG files.
	TagType_BaselineSharpness                 TagType                     = 50732 // ingore # Used in IFD 0 of DNG files.
	TagType_BayerGreenSplit                   TagType                     = 50733 // ingore # Used in Raw IFD of DNG files.
	TagType_LinearResponseLimit               TagType                     = 50734 // ingore # Used in IFD 0 of DNG files.
	TagType_CameraSerialNumber                TagType                     = 50735 // ingore # Used in IFD 0 of DNG files.
	TagType_LensInfo                          TagType                     = 50736 // ingore # Used in IFD 0 of DNG files.
	TagType_ChromaBlurRadius                  TagType                     = 50737 // ingore # Used in Raw IFD of DNG files.
	TagType_AntiAliasStrength                 TagType                     = 50738 // ingore # Used in Raw IFD of DNG files.
	TagType_DNGPrivateData                    TagType                     = 50740 // ingore # Used in IFD 0 of DNG files.
	TagType_MakerNoteSafety                   TagType                     = 50741 // ingore # Used in IFD 0 of DNG files.
	TagType_CalibrationIlluminant1            TagType                     = 50778 // ingore # Used in IFD 0 of DNG files.
	TagType_CalibrationIlluminant2            TagType                     = 50779 // ingore # Used in IFD 0 of DNG files.
	TagType_BestQualityScale                  TagType                     = 50780 // ingore # Used in Raw IFD of DNG files.
	TagType_AliasLayerMetadata                TagType                     = 50784 // ingore # Alias Sketchbook Pro layer usage description.
	_                                                                     = 0     //
)

// EXIF Tags
type ExifIFD_TagType TagType

const (
	ExifIFD_TagType_ExposureTime             ExifIFD_TagType = 33434 // ingore # Exposure time, given in seconds.
	ExifIFD_TagType_FNumber                  ExifIFD_TagType = 33437 // ingore # The F number.
	ExifIFD_TagType_ExposureProgram          ExifIFD_TagType = 34850 // ingore # The class of the program used by the camera to set exposure when the picture is taken.
	ExifIFD_TagType_SpectralSensitivity      ExifIFD_TagType = 34852 // ingore # Indicates the spectral sensitivity of each channel of the camera used.
	ExifIFD_TagType_ISOSpeedRatings          ExifIFD_TagType = 34855 // ingore # Indicates the ISO Speed and ISO Latitude of the camera or input device as specified in ISO 12232.
	ExifIFD_TagType_OECF                     ExifIFD_TagType = 34856 // ingore # Indicates the Opto-Electric Conversion Function (OECF) specified in ISO 14524.
	ExifIFD_TagType_ExifVersion              ExifIFD_TagType = 36864 // ingore # The version of the supported Exif standard.
	ExifIFD_TagType_DateTimeOriginal         ExifIFD_TagType = 36867 // ingore # The date and time when the original image data was generated.
	ExifIFD_TagType_DateTimeDigitized        ExifIFD_TagType = 36868 // ingore # The date and time when the image was stored as digital data.
	ExifIFD_TagType_ComponentsConfiguration  ExifIFD_TagType = 37121 // ingore # Specific to compressed data; specifies the channels and complements PhotometricInterpretation
	ExifIFD_TagType_CompressedBitsPerPixel   ExifIFD_TagType = 37122 // ingore # Specific to compressed data; states the compressed bits per pixel.
	ExifIFD_TagType_ShutterSpeedValue        ExifIFD_TagType = 37377 // ingore # Shutter speed.
	ExifIFD_TagType_ApertureValue            ExifIFD_TagType = 37378 // ingore # The lens aperture.
	ExifIFD_TagType_BrightnessValue          ExifIFD_TagType = 37379 // ingore # The value of brightness.
	ExifIFD_TagType_ExposureBiasValue        ExifIFD_TagType = 37380 // ingore # The exposure bias.
	ExifIFD_TagType_MaxApertureValue         ExifIFD_TagType = 37381 // ingore # The smallest F number of the lens.
	ExifIFD_TagType_SubjectDistance          ExifIFD_TagType = 37382 // ingore # The distance to the subject, given in meters.
	ExifIFD_TagType_MeteringMode             ExifIFD_TagType = 37383 // ingore # The metering mode.
	ExifIFD_TagType_LightSource              ExifIFD_TagType = 37384 // ingore # The kind of light source.
	ExifIFD_TagType_Flash                    ExifIFD_TagType = 37385 // ingore # Indicates the status of flash when the image was shot.
	ExifIFD_TagType_FocalLength              ExifIFD_TagType = 37386 // ingore # The actual focal length of the lens, in mm.
	ExifIFD_TagType_SubjectArea              ExifIFD_TagType = 37396 // ingore # Indicates the location and area of the main subject in the overall scene.
	ExifIFD_TagType_MakerNote                ExifIFD_TagType = 37500 // ingore # Manufacturer specific information.
	ExifIFD_TagType_UserComment              ExifIFD_TagType = 37510 // ingore # Keywords or comments on the image; complements ImageDescription.
	ExifIFD_TagType_SubsecTime               ExifIFD_TagType = 37520 // ingore # A tag used to record fractions of seconds for the DateTime tag.
	ExifIFD_TagType_SubsecTimeOriginal       ExifIFD_TagType = 37521 // ingore # A tag used to record fractions of seconds for the DateTimeOriginal tag.
	ExifIFD_TagType_SubsecTimeDigitized      ExifIFD_TagType = 37522 // ingore # A tag used to record fractions of seconds for the DateTimeDigitized tag.
	ExifIFD_TagType_FlashpixVersion          ExifIFD_TagType = 40960 // ingore # The Flashpix format version supported by a FPXR file.
	ExifIFD_TagType_ColorSpace               ExifIFD_TagType = 40961 // ingore # The color space information tag is always recorded as the color space specifier.
	ExifIFD_TagType_PixelXDimension          ExifIFD_TagType = 40962 // ingore # Specific to compressed data; the valid width of the meaningful image.
	ExifIFD_TagType_PixelYDimension          ExifIFD_TagType = 40963 // ingore # Specific to compressed data; the valid height of the meaningful image.
	ExifIFD_TagType_RelatedSoundFile         ExifIFD_TagType = 40964 // ingore # Used to record the name of an audio file related to the image data.
	ExifIFD_TagType_FlashEnergy              ExifIFD_TagType = 41483 // ingore # Indicates the strobe energy at the time the image is captured, as measured in Beam Candle Power Seconds
	ExifIFD_TagType_SpatialFrequencyResponse ExifIFD_TagType = 41484 // ingore # Records the camera or input device spatial frequency table and SFR values in the direction of image width, image height, and diagonal direction, as specified in ISO 12233.
	ExifIFD_TagType_FocalPlaneXResolution    ExifIFD_TagType = 41486 // ingore # Indicates the number of pixels in the image width (X) direction per FocalPlaneResolutionUnit on the camera focal plane.
	ExifIFD_TagType_FocalPlaneYResolution    ExifIFD_TagType = 41487 // ingore # Indicates the number of pixels in the image height (Y) direction per FocalPlaneResolutionUnit on the camera focal plane.
	ExifIFD_TagType_FocalPlaneResolutionUnit ExifIFD_TagType = 41488 // ingore # Indicates the unit for measuring FocalPlaneXResolution and FocalPlaneYResolution.
	ExifIFD_TagType_SubjectLocation          ExifIFD_TagType = 41492 // ingore # Indicates the location of the main subject in the scene.
	ExifIFD_TagType_ExposureIndex            ExifIFD_TagType = 41493 // ingore # Indicates the exposure index selected on the camera or input device at the time the image is captured.
	ExifIFD_TagType_SensingMethod            ExifIFD_TagType = 41495 // ingore # Indicates the image sensor type on the camera or input device.
	ExifIFD_TagType_FileSource               ExifIFD_TagType = 41728 // ingore # Indicates the image source.
	ExifIFD_TagType_SceneType                ExifIFD_TagType = 41729 // ingore # Indicates the type of scene.
	ExifIFD_TagType_CFAPattern               ExifIFD_TagType = 41730 // ingore # Indicates the color filter array (CFA) geometric pattern of the image sensor when a one-chip color area sensor is used.
	ExifIFD_TagType_CustomRendered           ExifIFD_TagType = 41985 // ingore # Indicates the use of special processing on image data, such as rendering geared to output.
	ExifIFD_TagType_ExposureMode             ExifIFD_TagType = 41986 // ingore # Indicates the exposure mode set when the image was shot.
	ExifIFD_TagType_WhiteBalance             ExifIFD_TagType = 41987 // ingore # Indicates the white balance mode set when the image was shot.
	ExifIFD_TagType_DigitalZoomRatio         ExifIFD_TagType = 41988 // ingore # Indicates the digital zoom ratio when the image was shot.
	ExifIFD_TagType_FocalLengthIn35mmFilm    ExifIFD_TagType = 41989 // ingore # Indicates the equivalent focal length assuming a 35mm film camera, in mm.
	ExifIFD_TagType_SceneCaptureType         ExifIFD_TagType = 41990 // ingore # Indicates the type of scene that was shot.
	ExifIFD_TagType_GainControl              ExifIFD_TagType = 41991 // ingore # Indicates the degree of overall image gain adjustment.
	ExifIFD_TagType_Contrast                 ExifIFD_TagType = 41992 // ingore # Indicates the direction of contrast processing applied by the camera when the image was shot.
	ExifIFD_TagType_Saturation               ExifIFD_TagType = 41993 // ingore # Indicates the direction of saturation processing applied by the camera when the image was shot.
	ExifIFD_TagType_Sharpness                ExifIFD_TagType = 41994 // ingore # Indicates the direction of sharpness processing applied by the camera when the image was shot.
	ExifIFD_TagType_DeviceSettingDescription ExifIFD_TagType = 41995 // ingore # This tag indicates information on the picture-taking conditions of a particular camera model.
	ExifIFD_TagType_SubjectDistanceRange     ExifIFD_TagType = 41996 // ingore # Indicates the distance to the subject.
	ExifIFD_TagType_ImageUniqueID            ExifIFD_TagType = 42016 // ingore # Indicates an identifier assigned uniquely to each image.
)

// GPS Tags
type GPSIFD_TagType TagType

const (
	GPSIFD_TagType_GPSVersionID        GPSIFD_TagType = 0  // ingore # Indicates the version of GPSInfoIFD.
	GPSIFD_TagType_GPSLatitudeRef      GPSIFD_TagType = 1  // ingore # Indicates whether the latitude is north or south latitude.
	GPSIFD_TagType_GPSLatitude         GPSIFD_TagType = 2  // ingore # Indicates the latitude.
	GPSIFD_TagType_GPSLongitudeRef     GPSIFD_TagType = 3  // ingore # Indicates whether the longitude is east or west longitude.
	GPSIFD_TagType_GPSLongitude        GPSIFD_TagType = 4  // ingore # Indicates the longitude.
	GPSIFD_TagType_GPSAltitudeRef      GPSIFD_TagType = 5  // ingore # Indicates the altitude used as the reference altitude.
	GPSIFD_TagType_GPSAltitude         GPSIFD_TagType = 6  // ingore # Indicates the altitude based on the reference in GPSAltitudeRef.
	GPSIFD_TagType_GPSTimeStamp        GPSIFD_TagType = 7  // ingore # Indicates the time as UTC (Coordinated Universal Time).
	GPSIFD_TagType_GPSSatellites       GPSIFD_TagType = 8  // ingore # Indicates the GPS satellites used for measurements.
	GPSIFD_TagType_GPSStatus           GPSIFD_TagType = 9  // ingore # Indicates the status of the GPS receiver when the image is recorded.
	GPSIFD_TagType_GPSMeasureMode      GPSIFD_TagType = 10 // ingore # Indicates the GPS measurement mode.
	GPSIFD_TagType_GPSDOP              GPSIFD_TagType = 11 // ingore # Indicates the GPS DOP (data degree of precision).
	GPSIFD_TagType_GPSSpeedRef         GPSIFD_TagType = 12 // ingore # Indicates the unit used to express the GPS receiver speed of movement.
	GPSIFD_TagType_GPSSpeed            GPSIFD_TagType = 13 // ingore # Indicates the speed of GPS receiver movement.
	GPSIFD_TagType_GPSTrackRef         GPSIFD_TagType = 14 // ingore # Indicates the reference for giving the direction of GPS receiver movement.
	GPSIFD_TagType_GPSTrack            GPSIFD_TagType = 15 // ingore # Indicates the direction of GPS receiver movement.
	GPSIFD_TagType_GPSImgDirectionRef  GPSIFD_TagType = 16 // ingore # Indicates the reference for giving the direction of the image when it is captured.
	GPSIFD_TagType_GPSImgDirection     GPSIFD_TagType = 17 // ingore # Indicates the direction of the image when it was captured.
	GPSIFD_TagType_GPSMapDatum         GPSIFD_TagType = 18 // ingore # Indicates the geodetic survey data used by the GPS receiver.
	GPSIFD_TagType_GPSDestLatitudeRef  GPSIFD_TagType = 19 // ingore # Indicates whether the latitude of the destination point is north or south latitude.
	GPSIFD_TagType_GPSDestLatitude     GPSIFD_TagType = 20 // ingore # Indicates the latitude of the destination point.
	GPSIFD_TagType_GPSDestLongitudeRef GPSIFD_TagType = 21 // ingore # Indicates whether the longitude of the destination point is east or west longitude.
	GPSIFD_TagType_GPSDestLongitude    GPSIFD_TagType = 22 // ingore # Indicates the longitude of the destination point.
	GPSIFD_TagType_GPSDestBearingRef   GPSIFD_TagType = 23 // ingore # Indicates the reference used for giving the bearing to the destination point.
	GPSIFD_TagType_GPSDestBearing      GPSIFD_TagType = 24 // ingore # Indicates the bearing to the destination point.
	GPSIFD_TagType_GPSDestDistanceRef  GPSIFD_TagType = 25 // ingore # Indicates the unit used to express the distance to the destination point.
	GPSIFD_TagType_GPSDestDistance     GPSIFD_TagType = 26 // ingore # Indicates the distance to the destination point.
	GPSIFD_TagType_GPSProcessingMethod GPSIFD_TagType = 27 // ingore # A character string recording the name of the method used for location finding.
	GPSIFD_TagType_GPSAreaInformation  GPSIFD_TagType = 28 // ingore # A character string recording the name of the GPS area.
	GPSIFD_TagType_GPSDateStamp        GPSIFD_TagType = 29 // ingore # A character string recording date and time information relative to UTC (Coordinated Universal Time).
	GPSIFD_TagType_GPSDifferential     GPSIFD_TagType = 30 // ingore # Indicates whether differential correction is applied to the GPS receiver.
)

// Interoperability Tags
type InteroperabilityIFD_TagType TagType

const (
	InteroperabilityIFD_TagType_InteroperabilityIndex InteroperabilityIFD_TagType = 1 // ingore # Indicates the identification of the Interoperability rule.
)
