// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	tiff "."
)

func main() {
	data, err := ioutil.ReadFile("./testdata/no_compress.tiff")
	if err != nil {
		log.Fatal(err)
	}

	header, err := tiff.ReadHeader(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header)

	ifd, err := tiff.ReadIFD(bytes.NewReader(data), header, header.Offset)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ifd)
}
