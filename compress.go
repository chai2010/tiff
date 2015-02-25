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
)

func (p TagValue_CompressionType) ReadAll(r io.Reader) (data []byte, err error) {
	switch p {
	case TagValue_CompressionType_None, TagValue_CompressionType_Nil:
		return p.readAll_None(r)
	case TagValue_CompressionType_CCITT:
		return p.readAll_CCITT(r)
	case TagValue_CompressionType_G3:
		return p.readAll_G3(r)
	case TagValue_CompressionType_G4:
		return p.readAll_G4(r)
	case TagValue_CompressionType_LZW:
		return p.readAll_LZW(r)
	case TagValue_CompressionType_JPEGOld:
		return p.readAll_JPEGOld(r)
	case TagValue_CompressionType_JPEG:
		return p.readAll_JPEG(r)
	case TagValue_CompressionType_Deflate:
		return p.readAll_Deflate(r)
	case TagValue_CompressionType_PackBits:
		return p.readAll_PackBits(r)
	case TagValue_CompressionType_DeflateOld:
		return p.readAll_DeflateOld(r)
	}
	err = fmt.Errorf("tiff: unknown TagValue_CompressionType, %d", int(p))
	return
}

func (p TagValue_CompressionType) readAll_None(r io.Reader) (data []byte, err error) {
	data, err = ioutil.ReadAll(r)
	return
}

func (p TagValue_CompressionType) readAll_CCITT(r io.Reader) (data []byte, err error) {
	err = fmt.Errorf("tiff: unsupport TagValue_CompressionType, %d", int(p))
	return
}

func (p TagValue_CompressionType) readAll_G3(r io.Reader) (data []byte, err error) {
	err = fmt.Errorf("tiff: unsupport TagValue_CompressionType, %d", int(p))
	return
}

func (p TagValue_CompressionType) readAll_G4(r io.Reader) (data []byte, err error) {
	err = fmt.Errorf("tiff: unsupport TagValue_CompressionType, %d", int(p))
	return
}

func (p TagValue_CompressionType) readAll_LZW(r io.Reader) (data []byte, err error) {
	lzwReader := newLzwReader(r, lzwMSB, 8)
	data, err = ioutil.ReadAll(lzwReader)
	lzwReader.Close()
	return
}

func (p TagValue_CompressionType) readAll_JPEGOld(r io.Reader) (data []byte, err error) {
	err = fmt.Errorf("tiff: unsupport TagValue_CompressionType, %d", int(p))
	return
}

func (p TagValue_CompressionType) readAll_JPEG(r io.Reader) (data []byte, err error) {
	err = fmt.Errorf("tiff: unsupport TagValue_CompressionType, %d", int(p))
	return
}

func (p TagValue_CompressionType) readAll_Deflate(r io.Reader) (data []byte, err error) {
	zlibReader, err := zlib.NewReader(r)
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(zlibReader)
	zlibReader.Close()
	return
}

func (p TagValue_CompressionType) readAll_DeflateOld(r io.Reader) (data []byte, err error) {
	zlibReader, err := zlib.NewReader(r)
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(zlibReader)
	zlibReader.Close()
	return
}

func (p TagValue_CompressionType) readAll_PackBits(r io.Reader) (data []byte, err error) {
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
