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

	"github.com/chai2010/tiff.go"
)

func main() {
	var buf bytes.Buffer
	var data []byte
	var err error

	// Load file data
	if data, err = ioutil.ReadFile("./testdata/video-001.tiff"); err != nil {
		log.Fatal(err)
	}

	// Decode tiff
	m, err := tiff.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	// Encode tiff
	if err = tiff.Encode(&buf, m, nil); err != nil {
		log.Fatal(err)
	}
	if err = ioutil.WriteFile("output.tiff", buf.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Save output.tiff ok\n")
}
