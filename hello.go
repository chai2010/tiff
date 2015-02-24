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
	"os"
	"path/filepath"

	tiff "github.com/chai2010/tiff"
)

var files = []string{
	"./testdata/BigTIFFSamples/BigTIFFSubIFD8.tif",
	"./testdata/multipage/multipage-gopher.tif",
}

func main() {
	for _, name := range files {
		// Load file data
		f, err := os.Open(name)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		// Decode tiff
		m, errors, err := tiff.DecodeAll(f)
		if err != nil {
			log.Println(err)
		}

		// Encode tiff
		for i := 0; i < len(m); i++ {
			for j := 0; j < len(m[i]); j++ {
				newname := fmt.Sprintf("%s-%02d-%02d.tiff", filepath.Base(name), i, j)
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
