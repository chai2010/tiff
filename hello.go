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

	tiff "github.com/chai2010/tiff"
)

func main() {
	var buf bytes.Buffer
	var data []byte
	var err error

	// Load file data
	if data, err = ioutil.ReadFile("./testdata/BigTIFFSamples/BigTIFFSubIFD8.tif"); err != nil {
		log.Fatal(err)
	}

	// Decode tiff
	m, err := tiff.DecodeAll(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	// Encode tiff
	for i := 0; i < len(m); i++ {
		filename := fmt.Sprintf("output-frame-%02d.tiff", i)
		if err = tiff.Encode(&buf, m[i], nil); err != nil {
			log.Fatal(err)
		}
		if err = ioutil.WriteFile(filename, buf.Bytes(), 0666); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Save %s ok\n", filename)
	}
}
