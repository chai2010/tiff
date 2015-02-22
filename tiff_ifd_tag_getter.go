// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"fmt"
	"time"
)

var _ TagGetter = (*tifTagGetter)(nil)

type tifTagGetter struct {
	EntryMap map[TagType]*IFDEntry
}

func (p *tifTagGetter) GetNewSubfileType() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_NewSubfileType]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetSubfileType() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_SubfileType]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetImageWidth() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_ImageWidth]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetImageLength() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_ImageLength]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetBitsPerSample() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_BitsPerSample]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetCompression() (value CompressType, ok bool) {
	var entry *IFDEntry
	value = CompressType_Nil
	if entry, ok = p.EntryMap[TagType_Compression]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = CompressType(v[0])
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetPhotometricInterpretation() (value TagValue_PhotometricType, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_PhotometricInterpretation]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = TagValue_PhotometricType(v[0])
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetThreshholding() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_Threshholding]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetCellWidth() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_CellWidth]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetCellLenght() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_CellLenght]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetFillOrder() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_FillOrder]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetDocumentName() (value string, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_DocumentName]; ok {
		value = entry.GetString()
	}
	return
}

func (p *tifTagGetter) GetImageDescription() (value string, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_ImageDescription]; ok {
		value = entry.GetString()
	}
	return
}

func (p *tifTagGetter) GetMake() (value string, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_Make]; ok {
		value = entry.GetString()
	}
	return
}

func (p *tifTagGetter) GetModel() (value string, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_Model]; ok {
		value = entry.GetString()
	}
	return
}

func (p *tifTagGetter) GetStripOffsets() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_StripOffsets]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetOrientation() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_Orientation]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetSamplesPerPixel() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_SamplesPerPixel]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetRowsPerStrip() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_RowsPerStrip]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetStripByteCounts() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_StripByteCounts]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetMinSampleValue() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_MinSampleValue]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetMaxSampleValue() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_MaxSampleValue]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetXResolution() (value [2]int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_XResolution]; ok {
		if v := entry.GetInts(); len(v) == 2 {
			value = [2]int64{v[0], v[1]}
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetYResolution() (value [2]int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_YResolution]; ok {
		if v := entry.GetInts(); len(v) == 2 {
			value = [2]int64{v[0], v[1]}
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetPlanarConfiguration() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_PlanarConfiguration]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetPageName() (value string, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_PageName]; ok {
		value = entry.GetString()
	}
	return
}

func (p *tifTagGetter) GetXPosition() (value [2]int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_XPosition]; ok {
		if v := entry.GetInts(); len(v) == 2 {
			value = [2]int64{v[0], v[1]}
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetYPosition() (value [2]int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_YPosition]; ok {
		if v := entry.GetInts(); len(v) == 2 {
			value = [2]int64{v[0], v[1]}
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetFreeOffsets() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_FreeOffsets]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetFreeByteCounts() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_FreeByteCounts]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetGrayResponseUnit() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_GrayResponseUnit]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetGrayResponseCurve() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_GrayResponseCurve]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetT4Options() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_T4Options]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetT6Options() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_T6Options]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetResolutionUnit() (value TagValue_ResolutionUnitType, ok bool) {
	var entry *IFDEntry
	value = TagValue_ResolutionUnitType_None
	if entry, ok = p.EntryMap[TagType_ResolutionUnit]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = TagValue_ResolutionUnitType(v[0])
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetPageNumber() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_PageNumber]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetTransferFunction() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_TransferFunction]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetSoftware() (value string, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_Software]; ok {
		value = entry.GetString()
	}
	return
}

func (p *tifTagGetter) GetDateTime() (value time.Time, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_DateTime]; ok {
		var year, month, day, hour, min, sec int
		if _, err := fmt.Sscanf(entry.GetString(), "%d:%d:%d %d:%d:%d",
			&year, &month, &day,
			&hour, &min, &sec,
		); err != nil {
			ok = false
			return
		}
		value = time.Date(year, time.Month(month), day, hour, min, sec, 0, nil)
	}
	return
}

func (p *tifTagGetter) GetArtist() (value string, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_Artist]; ok {
		value = entry.GetString()
	}
	return
}

func (p *tifTagGetter) GetHostComputer() (value string, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_HostComputer]; ok {
		value = entry.GetString()
	}
	return
}

func (p *tifTagGetter) GetPredictor() (value TagValue_PredictorType, ok bool) {
	var entry *IFDEntry
	value = TagValue_PredictorType_None
	if entry, ok = p.EntryMap[TagType_Predictor]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = TagValue_PredictorType(v[0])
		}
	}
	return
}

func (p *tifTagGetter) GetWhitePoint() (value [][2]int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_WhitePoint]; ok {
		value = entry.GetRationals()
	}
	return
}

func (p *tifTagGetter) GetPrimaryChromaticities() (value [][2]int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_PrimaryChromaticities]; ok {
		value = entry.GetRationals()
	}
	return
}

func (p *tifTagGetter) GetColorMap() (value [][3]uint16, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_ColorMap]; ok {
		v := entry.GetInts()
		if len(v) == 0 || len(v)%3 != 0 {
			ok = false
			return
		}
		value := make([][3]uint16, len(v)/3)
		for i := 0; i < len(v); i += 3 {
			value[i/3] = [3]uint16{
				uint16(v[i+0]),
				uint16(v[i+1]),
				uint16(v[i+2]),
			}
		}
	}
	return
}

func (p *tifTagGetter) GetHalftoneHints() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_HalftoneHints]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetTileWidth() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_TileWidth]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetTileLength() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_TileLength]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetTileOffsets() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_TileOffsets]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetTileByteCounts() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_TileByteCounts]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetInkSet() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_InkSet]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetInkNames() (value string, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_InkNames]; ok {
		value = entry.GetString()
	}
	return
}

func (p *tifTagGetter) GetNumberOfInks() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_NumberOfInks]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetDotRange() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_DotRange]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetTargetPrinter() (value string, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_TargetPrinter]; ok {
		value = entry.GetString()
	}
	return
}

func (p *tifTagGetter) GetExtraSamples() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_ExtraSamples]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetSampleFormat() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_SampleFormat]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetSMinSampleValue() (value []float64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_SMinSampleValue]; ok {
		value = entry.GetFloats()
	}
	return
}

func (p *tifTagGetter) GetSMaxSampleValue() (value []float64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_SMaxSampleValue]; ok {
		value = entry.GetFloats()
	}
	return
}

func (p *tifTagGetter) GetTransferRange() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_TransferRange]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetJPEGProc() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_JPEGProc]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetJPEGInterchangeFormat() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_JPEGInterchangeFormat]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetJPEGInterchangeFormatLngth() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_JPEGInterchangeFormatLngth]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetJPEGRestartInterval() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_JPEGRestartInterval]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetJPEGLosslessPredictors() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_JPEGLosslessPredictors]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetJPEGPointTransforms() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_JPEGPointTransforms]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetJPEGQTables() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_JPEGQTables]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetJPEGDCTables() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_JPEGDCTables]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetJPEGACTables() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_JPEGACTables]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetYCbCrCoefficients() (value [][2]int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_YCbCrCoefficients]; ok {
		value = entry.GetRationals()
	}
	return
}

func (p *tifTagGetter) GetYCbCrSubSampling() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_YCbCrSubSampling]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetYCbCrPositioning() (value int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_YCbCrPositioning]; ok {
		if v := entry.GetInts(); len(v) == 1 {
			value = v[0]
		} else {
			ok = false
		}
	}
	return
}

func (p *tifTagGetter) GetReferenceBlackWhite() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_ReferenceBlackWhite]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetCopyright() (value string, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_Copyright]; ok {
		value = entry.GetString()
	}
	return
}

func (p *tifTagGetter) GetGeoKeyDirectoryTag() (value []int64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_GeoKeyDirectoryTag]; ok {
		value = entry.GetInts()
	}
	return
}

func (p *tifTagGetter) GetGeoDoubleParamsTag() (value []float64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_GeoDoubleParamsTag]; ok {
		value = entry.GetFloats()
	}
	return
}

func (p *tifTagGetter) GetGeoAsciiParamsTag() (value string, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_GeoAsciiParamsTag]; ok {
		value = entry.GetString()
	}
	return
}

func (p *tifTagGetter) GetModelTiepointTag() (value []float64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_ModelTiepointTag]; ok {
		value = entry.GetFloats()
	}
	return
}

func (p *tifTagGetter) GetModelPixelScaleTag() (value []float64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_ModelPixelScaleTag]; ok {
		value = entry.GetFloats()
	}
	return
}

func (p *tifTagGetter) GetModelTransformationTag() (value []float64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_ModelTransformationTag]; ok {
		value = entry.GetFloats()
	}
	return
}

func (p *tifTagGetter) GetIntergraphMatrixTag() (value []float64, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_IntergraphMatrixTag]; ok {
		value = entry.GetFloats()
	}
	return
}

func (p *tifTagGetter) GetUnknown(tag TagType) (value []byte, ok bool) {
	var entry *IFDEntry
	if entry, ok = p.EntryMap[TagType_IntergraphMatrixTag]; ok {
		value = entry.Data
	}
	return
}

func (p *tifTagGetter) private() {
	return
}
