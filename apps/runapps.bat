:: Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

@setlocal

@cd %~dp0

@go run tiffinfo.go                               ^
  ..\testdata\blue-purple-pink.lzwcompressed.tiff ^
  ..\testdata\bw-deflate.tiff                     ^
  ..\testdata\bw-packbits.tiff                    ^
  ..\testdata\bw-uncompressed.tiff                ^
  ..\testdata\no_compress.tiff                    ^
  ..\testdata\no_rps.tiff                         ^
  ..\testdata\video-001-16bit.tiff                ^
  ..\testdata\video-001-gray-16bit.tiff           ^
  ..\testdata\video-001-gray.tiff                 ^
  ..\testdata\video-001-paletted.tiff             ^
  ..\testdata\video-001-strip-64.tiff             ^
  ..\testdata\video-001-tile-64x64.tiff           ^
  ..\testdata\video-001-uncompressed.tiff         ^
  ..\testdata\video-001.tiff                      ^
  > tiffinfo-out.txt
