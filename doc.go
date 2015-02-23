// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package tiff offers a rich BigTIFF/TIFF decoder/encoder for Go.

The Classic TIFF specification is at http://partners.adobe.com/public/developer/en/tiff/TIFF6.pdf

The Big TIFF specification is at http://www.remotesensing.org/libtiff/bigtiffdesign.html

The Geo TIFF specification is at http://www.remotesensing.org/geotiff/spec/geotiffhome.html

Features:

	1. Support BigTiff
	2. Support decode multiple image
	3. Support RGB format
	4. Support Float DataType
	5. More ...

Example:

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

Report bugs to <chaishushan@gmail.com>.

Thanks!
*/
package tiff
