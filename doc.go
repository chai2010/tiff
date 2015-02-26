// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package tiff offers a rich TIFF/BigTIFF/GeoTIFF decoder/encoder for Go.

The Classic TIFF specification is at http://partners.adobe.com/public/developer/en/tiff/TIFF6.pdf

The Big TIFF specification is at http://www.remotesensing.org/libtiff/bigtiffdesign.html

The Geo TIFF specification is at http://www.remotesensing.org/geotiff/spec/geotiffhome.html

TIFF Tags: http://www.awaresystems.be/imaging/tiff/tifftags.html

Features:

	1. Support BigTiff
	2. Support decode multiple image
	3. Support decode subifd image
	4. Support RGB format
	5. Support Float DataType
	6. More ...

Example:

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
					newname := fmt.Sprintf("%s-%02d-%02d.tiff", filepath.Base(filename), i, j)
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

Report bugs to <chaishushan@gmail.com>.

Thanks!
*/
package tiff

/*
The Classic TIFF specification is at http://partners.adobe.com/public/developer/en/tiff/TIFF6.pdf

The Big TIFF specification is at http://www.remotesensing.org/libtiff/bigtiffdesign.html

The Geo TIFF specification is at http://www.remotesensing.org/geotiff/spec/geotiffhome.html

Classic TIFF Structure:

	+------------------------------------------------------------------------------+
	|                          Classic TIFF Structure                              |
	|  IFH                                                                         |
	| +-----------------+                                                          |
	| | II/MM ([2]byte) |                                                          |
	| +-----------------+                                                          |
	| | 42     (uint16) |      IFD                                                 |
	| +-----------------+     +------------------+                                 |
	| | Offset (uint32) |---->| Num     (uint16) |                                 |
	| +-----------------+     +------------------+                                 |
	|                         | Entry ([12]byte) |                                 |
	|                         +------------------+                                 |
	|                         | Entry ([12]byte) |                                 |
	|                         +------------------+                                 |
	|                         | ...              |      IFD                        |
	|                         +------------------+    +------------------+         |
	|     IFD Entry           | Offset  (uint32) |--->| Num     (uint16) |         |
	|    +-----------------+  +------------------+    +------------------+         |
	|    | Tag    (uint16) |                          | Entry ([12]byte) |         |
	|    +-----------------+                          +------------------+         |
	|    | Type   (uint16) |<-------------------------| Entry ([12]byte) |         |
	|    +-----------------+                          +------------------+         |
	|    | Count  (uint32) |                          | ...              |         |
	|    +-----------------+                          +------------------+         |
	|    | Offset (uint32) |---> Value                | Offset  (uint32) |--->NULL |
	|    +-----------------+                          +------------------+         |
	|                                                                              |
	+------------------------------------------------------------------------------+

Big TIFF Structure:

	+------------------------------------------------------------------------------+
	|                          Big TIFF Structure                                  |
	|  IFH                                                                         |
	| +-----------------+                                                          |
	| | II/MM ([2byte]) |                                                          |
	| +-----------------+                                                          |
	| | 43     (uint16) |                                                          |
	| +-----------------+                                                          |
	| | 8      (uint16) |                                                          |
	| +-----------------+                                                          |
	| | 0      (uint16) |      IFD                                                 |
	| +-----------------+     +------------------+                                 |
	| | Offset (uint64) |---->| Num     (uint64) |                                 |
	| +-----------------+     +------------------+                                 |
	|                         | Entry ([20]byte) |                                 |
	|                         +------------------+                                 |
	|                         | Entry ([20]byte) |                                 |
	|                         +------------------+                                 |
	|                         | ...              |      IFD                        |
	|                         +------------------+    +------------------+         |
	|     IFD Entry           | Offset  (uint64) |--->| Num     (uint64) |         |
	|    +-----------------+  +------------------+    +------------------+         |
	|    | Tag    (uint16) |                          | Entry ([12]byte) |         |
	|    +-----------------+                          +------------------+         |
	|    | Type   (uint16) |<-------------------------| Entry ([12]byte) |         |
	|    +-----------------+                          +------------------+         |
	|    | Count  (uint64) |                          | ...              |         |
	|    +-----------------+                          +------------------+         |
	|    | Offset (uint64) |---> Value                | Offset  (uint64) |--->NULL |
	|    +-----------------+                          +------------------+         |
	|                                                                              |
	+------------------------------------------------------------------------------+
*/
