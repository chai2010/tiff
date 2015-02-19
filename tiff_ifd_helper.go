// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

func (p *IFD) ImageSize() (width, height int) {
	return
}

func (p *IFD) BitsPerSample() int {
	return 0
}

func (p *IFD) Compression() CompressType {
	return CompressType_Nil
}

func (p *IFD) CellSize() (width, height float64) {
	return
}

func (p *IFD) BlockSize() (width, height int) {
	return
}

func (p *IFD) BlockOffsets() []int64 {
	return nil
}

func (p *IFD) DocumentName() string {
	return ""
}

func (p *IFD) ImageDescription() string {
	return ""
}

func (p *IFD) Make() string {
	return ""
}

func (p *IFD) Model() string {
	return ""
}

func (p *IFD) Copyright() string {
	return ""
}
