// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

import (
	"image"
	"os"
)

func Load(filename string) (m image.Image, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()
	return Decode(f)
}

func Save(filename string, m image.Image, opt *Options) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return Encode(f, m, opt)
}

// minInt returns the smaller of x or y.
func minInt(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
