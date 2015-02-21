// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

var _ TagSetter = (*tifTagSetter)(nil)

type tifTagSetter struct {
	TagSetter
}
