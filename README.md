TIFF for Go
===========

PkgDoc: [http://godoc.org/github.com/chai2010/tiff.go](http://godoc.org/github.com/chai2010/tiff.go)

Install
=======

1. `go get -d github.com/chai2010/tiff.go`
2. `cd github.com/chai2010/tiff.go && go install`
3. `go run hello.go`

Example
=======

```Go
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
```

BUGS
====

Report bugs to <chaishushan@gmail.com>.

Thanks!
