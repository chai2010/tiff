- *赞助 BTC: 1Cbd6oGAUUyBi7X7MaR4np4nTmQZXVgkCW*
- *赞助 ETH: 0x623A3C3a72186A6336C79b18Ac1eD36e1c71A8a6*
- *Go语言付费QQ群: 1055927514*

----

TIFF for Go
===========

[![Build Status](https://travis-ci.org/chai2010/tiff.svg)](https://travis-ci.org/chai2010/tiff)
[![GoDoc](https://godoc.org/github.com/chai2010/tiff?status.svg)](https://godoc.org/github.com/chai2010/tiff)


**Features:**

1. Support BigTiff
2. Support decode multiple image
3. Support decode subifd image
4. Support RGB format
5. Support Float DataType
6. More ...

Install
=======

1. `go get github.com/chai2010/tiff`
2. `go run hello.go`

Example
=======

```Go
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
```

BUGS
====

Report bugs to <chaishushan@gmail.com>.

Thanks!
