TIFF for Go
===========

PkgDoc: [http://godoc.org/github.com/chai2010/tiff](http://godoc.org/github.com/chai2010/tiff)


**Features:**

1. Support BigTiff
2. Support decode multiple image
3. Support RGB format
4. Support Float DataType
5. More ...

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

	tiff "github.com/chai2010/tiff"
)

func main() {
	var data []byte
	var err error

	// Load file data
	if data, err = ioutil.ReadFile("./testdata/multipage/multipage-gopher.tif"); err != nil {
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

		var buf bytes.Buffer
		if err = tiff.Encode(&buf, m[i], nil); err != nil {
			log.Fatal(err)
		}
		if err = ioutil.WriteFile(filename, buf.Bytes(), 0666); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Save %s ok\n", filename)
	}
}
```

BUGS
====

Report bugs to <chaishushan@gmail.com>.

Thanks!
