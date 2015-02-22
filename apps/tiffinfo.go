// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"fmt"
	"log"
	"os"

	tiff "github.com/chai2010/tiff"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("usage: tiffinfo filenames ...")
		os.Exit(1)
	}
	for i := 1; i < len(os.Args); i++ {
		printTiffInfo(os.Args[i])
	}
}

func printTiffInfo(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("printTiffInfo: os.Open, ", err)
	}
	defer f.Close()

	fmt.Println("file:", filename)

	header, err := tiff.ReadHeader(f)
	if err != nil {
		log.Fatal("printTiffInfo: tiff.ReadHeader, ", err)
	}
	fmt.Println(header)

	ifd, err := tiff.ReadIFD(f, header, header.Offset)
	if err != nil {
		log.Fatal("printTiffInfo: tiff.ReadIFD, err =", err)
	}
	fmt.Println(ifd)
}
