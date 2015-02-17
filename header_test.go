// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"bytes"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestHeader_decodeAndEncode_files(t *testing.T) {
	files := []string{
		"BigTIFF.tif",
		"BigTIFFLong.tif",
		"BigTIFFLong8.tif",
		"BigTIFFLong8Tiles.tif",
		"BigTIFFMotorola.tif",
		"BigTIFFMotorolaLongStrips.tif",
		"BigTIFFSubIFD4.tif",
		"BigTIFFSubIFD8.tif",
	}
	for i := 0; i < len(files); i++ {
		data, err := ioutil.ReadFile("./testdata/" + files[i])
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
		h0, err := ReadHeader(bytes.NewReader(data))
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
		h1, err := ReadHeader(bytes.NewReader(h0.Bytes()))
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
		if !reflect.DeepEqual(h0, h1) {
			t.Fatalf("%d: not equal: %v != %v", i, h0, h1)
		}
	}
}

func TestHeader_decodeAndEncode(t *testing.T) {
	headers := []*Header{
		NewHeader(true, 9527),
		NewHeader(false, 9527),
		NewHeader(true, 9527),
		NewHeader(false, 8),
		NewHeader(true, 16),
	}
	for i := 0; i < len(headers); i++ {
		h, err := ReadHeader(bytes.NewReader(headers[i].Bytes()))
		if err != nil {
			t.Fatalf("%d: %v", i, err)
		}
		if !reflect.DeepEqual(headers[i], h) {
			t.Fatalf("%d: not equal: %v != %v", i, headers[i], h)
		}
	}
}

func TestHeader_decodeAndEncode_bad(t *testing.T) {
	headers := []*Header{
		NewHeader(true, 15),
		NewHeader(true, 8),
		NewHeader(true, 7),
		NewHeader(true, 0),
		NewHeader(false, 7),
		NewHeader(false, 6),
		NewHeader(false, 0),
	}
	for i := 0; i < len(headers); i++ {
		_, err := ReadHeader(bytes.NewReader(headers[i].Bytes()))
		if err == nil {
			t.Fatalf("%d: expect err, got %v", i, err)
		}
	}
}
