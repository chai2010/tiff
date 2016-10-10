// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"bufio"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/chai2010/tiff/internal/fax"
)

func (p TagValue_CompressionType) Decode(r io.Reader, width, height int) (data []byte, err error) {
	switch p {
	case TagValue_CompressionType_None, TagValue_CompressionType_Nil:
		return p.decode_None(r)
	case TagValue_CompressionType_CCITT:
		return p.decode_CCITT(r)
	case TagValue_CompressionType_G3:
		return p.decode_G3(r, width, height)
	case TagValue_CompressionType_G4:
		return p.decode_G4(r, width, height)
	case TagValue_CompressionType_LZW:
		return p.decode_LZW(r)
	case TagValue_CompressionType_JPEGOld:
		return p.decode_JPEGOld(r)
	case TagValue_CompressionType_JPEG:
		return p.decode_JPEG(r)
	case TagValue_CompressionType_Deflate:
		return p.decode_Deflate(r)
	case TagValue_CompressionType_PackBits:
		return p.decode_PackBits(r)
	case TagValue_CompressionType_DeflateOld:
		return p.decode_DeflateOld(r)
	}
	err = fmt.Errorf("tiff: unknown TagValue_CompressionType, %d", int(p))
	return
}

func (p TagValue_CompressionType) decode_None(r io.Reader) (data []byte, err error) {
	data, err = ioutil.ReadAll(r)
	return
}

func (p TagValue_CompressionType) decode_CCITT(r io.Reader) (data []byte, err error) {
	err = fmt.Errorf("tiff: unsupport TagValue_CompressionType, %d", int(p))
	return
}

func (p TagValue_CompressionType) decode_G3(r io.Reader, width, height int) (data []byte, err error) {
	return p.decode_G4(r, width, height)
}

func (p TagValue_CompressionType) decode_G4(r io.Reader, width, height int) (data []byte, err error) {
	br, ok := r.(io.ByteReader)
	if !ok {
		br = bufio.NewReader(r)
	}

	return fax.DecodeG4Pixels(br, width, height)
}

func (p TagValue_CompressionType) decode_LZW(r io.Reader) (data []byte, err error) {
	lzwReader := newLzwReader(r, lzwMSB, 8)
	data, err = ioutil.ReadAll(lzwReader)
	lzwReader.Close()
	return
}

func (p TagValue_CompressionType) decode_JPEGOld(r io.Reader) (data []byte, err error) {
	err = fmt.Errorf("tiff: unsupport TagValue_CompressionType, %d", int(p))
	return
}

func (p TagValue_CompressionType) decode_JPEG(r io.Reader) (data []byte, err error) {
	err = fmt.Errorf("tiff: unsupport TagValue_CompressionType, %d", int(p))
	return
}

func (p TagValue_CompressionType) decode_Deflate(r io.Reader) (data []byte, err error) {
	zlibReader, err := zlib.NewReader(r)
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(zlibReader)
	zlibReader.Close()
	return
}

func (p TagValue_CompressionType) decode_DeflateOld(r io.Reader) (data []byte, err error) {
	zlibReader, err := zlib.NewReader(r)
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(zlibReader)
	zlibReader.Close()
	return
}

func (p TagValue_CompressionType) decode_PackBits(r io.Reader) (data []byte, err error) {
	type byteReader interface {
		io.Reader
		io.ByteReader
	}

	buf := make([]byte, 128)
	dst := make([]byte, 0, 1024)
	br, ok := r.(byteReader)
	if !ok {
		br = bufio.NewReader(r)
	}

	for {
		b, err := br.ReadByte()
		if err != nil {
			if err == io.EOF {
				return dst, nil
			}
			return nil, err
		}
		code := int(int8(b))
		switch {
		case code >= 0:
			n, err := io.ReadFull(br, buf[:code+1])
			if err != nil {
				return nil, err
			}
			dst = append(dst, buf[:n]...)
		case code == -128:
			// No-op.
		default:
			if b, err = br.ReadByte(); err != nil {
				return nil, err
			}
			for j := 0; j < 1-code; j++ {
				buf[j] = b
			}
			dst = append(dst, buf[:1-code]...)
		}
	}
}
