// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tiff

type Options struct {
	EntryMap map[TagType]*IFDEntry
}

func (p *Options) TagGetter() TagGetter {
	return &tifTagGetter{
		EntryMap: p.EntryMap,
	}
}

func (p *Options) TagSetter() TagSetter {
	return &tifTagSetter{
		EntryMap: p.EntryMap,
	}
}
