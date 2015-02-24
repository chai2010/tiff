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
	"path/filepath"

	tiff "github.com/chai2010/tiff"
)

func main() {
	var data []byte
	var err error

	var files = []string{
		"./testdata/BigTIFFSamples/BigTIFFSubIFD8.tif",
		"./testdata/multipage/multipage-gopher.tif",
	}
	for _, filename := range files {
		// Load file data
		if data, err = ioutil.ReadFile(filename); err != nil {
			log.Fatal(err)
		}

		// Decode tiff
		m, errors, err := tiff.DecodeAll(bytes.NewReader(data))
		if err != nil {
			log.Println(err)
		}

		// Encode tiff
		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[i]); j++ {
				newname := fmt.Sprintf("output-%s-frame-%02d-sub-%02d.tiff", filepath.Base(filename), i, j)
				if errors[i][j] != nil {
					log.Printf("%s: %v\n", newname, err)
					continue
				}

				var buf bytes.Buffer
				if err = tiff.Encode(&buf, m[i][j], nil); err != nil {
					log.Fatal(err)
				}
				if err = ioutil.WriteFile(newname, buf.Bytes(), 0666); err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Save %s ok\n", newname)
			}
		}
	}
}
