// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"bytes"
	"image"
	"io/ioutil"
	"os"
	"testing"
)

var roundtripTests = []struct {
	filename string
}{
	{"video-001.tiff"},
	{"video-001-16bit.tiff"},
	{"video-001-gray.tiff"},
	{"video-001-gray-16bit.tiff"},
	{"video-001-paletted.tiff"},
	{"bw-packbits.tiff"},
}

func openImage(filename string) (image.Image, error) {
	f, err := os.Open(testdataDir + filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Decode(f)
}

func Test_loadAndSave(t *testing.T) {
	tmpname := "_test_loadAndSave_temp.tiff"
	defer os.Remove(tmpname)

	for _, rt := range roundtripTests {
		img, err := Load(testdataDir + rt.filename)
		if err != nil {
			t.Fatal(err)
		}
		err = Save(tmpname, img, nil)
		if err != nil {
			t.Fatal(err)
		}

		img2, err := Load(tmpname)
		if err != nil {
			t.Fatal(err)
		}
		compare(t, img, img2)
	}
}

func TestRoundtrip(t *testing.T) {
	for _, rt := range roundtripTests {
		img, err := openImage(rt.filename)
		if err != nil {
			t.Fatal(err)
		}
		out := new(bytes.Buffer)
		err = Encode(out, img, nil)
		if err != nil {
			t.Fatal(err)
		}

		img2, err := Decode(bytes.NewReader(out.Bytes()))
		if err != nil {
			t.Fatal(err)
		}
		compare(t, img, img2)
	}
}

// TestRoundtrip2 tests that encoding and decoding an image whose
// origin is not (0, 0) gives the same thing.
func TestRoundtrip2(t *testing.T) {
	m0 := image.NewRGBA(image.Rect(3, 4, 9, 8))
	for i := range m0.Pix {
		m0.Pix[i] = byte(i)
	}
	out := new(bytes.Buffer)
	if err := Encode(out, m0, nil); err != nil {
		t.Fatal(err)
	}
	m1, err := Decode(bytes.NewReader(out.Bytes()))
	if err != nil {
		t.Fatal(err)
	}
	compare(t, m0, m1)
}

func benchmarkEncode(b *testing.B, name string, pixelSize int) {
	img, err := openImage(name)
	if err != nil {
		b.Fatal(err)
	}
	s := img.Bounds().Size()
	b.SetBytes(int64(s.X * s.Y * pixelSize))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Encode(ioutil.Discard, img, nil)
	}
}

func BenchmarkEncode(b *testing.B)         { benchmarkEncode(b, "video-001.tiff", 4) }
func BenchmarkEncodePaletted(b *testing.B) { benchmarkEncode(b, "video-001-paletted.tiff", 1) }
func BenchmarkEncodeGray(b *testing.B)     { benchmarkEncode(b, "video-001-gray.tiff", 1) }
func BenchmarkEncodeGray16(b *testing.B)   { benchmarkEncode(b, "video-001-gray-16bit.tiff", 2) }
func BenchmarkEncodeRGBA(b *testing.B)     { benchmarkEncode(b, "video-001.tiff", 4) }
func BenchmarkEncodeRGBA64(b *testing.B)   { benchmarkEncode(b, "video-001-16bit.tiff", 8) }
