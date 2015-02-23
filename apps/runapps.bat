:: Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

@setlocal

@cd %~dp0

@go run tiffinfo.go                                         ^
  ..\testdata\blue-purple-pink.lzwcompressed.tiff           ^
  ..\testdata\bw-deflate.tiff                               ^
  ..\testdata\bw-packbits.tiff                              ^
  ..\testdata\bw-uncompressed.tiff                          ^
  ..\testdata\no_compress.tiff                              ^
  ..\testdata\no_rps.tiff                                   ^
  ..\testdata\video-001-16bit.tiff                          ^
  ..\testdata\video-001-gray-16bit.tiff                     ^
  ..\testdata\video-001-gray.tiff                           ^
  ..\testdata\video-001-paletted.tiff                       ^
  ..\testdata\video-001-strip-64.tiff                       ^
  ..\testdata\video-001-tile-64x64.tiff                     ^
  ..\testdata\video-001-uncompressed.tiff                   ^
  ..\testdata\video-001.tiff                                ^
                                                            ^
  ..\testdata\BigTIFFSamples\BigTIFF.tif                    ^
  ..\testdata\BigTIFFSamples\BigTIFFLong.tif                ^
  ..\testdata\BigTIFFSamples\BigTIFFLong8.tif               ^
  ..\testdata\BigTIFFSamples\BigTIFFLong8Tiles.tif          ^
  ..\testdata\BigTIFFSamples\BigTIFFMotorola.tif            ^
  ..\testdata\BigTIFFSamples\BigTIFFMotorolaLongStrips.tif  ^
  ..\testdata\BigTIFFSamples\BigTIFFSubIFD4.tif             ^
  ..\testdata\BigTIFFSamples\BigTIFFSubIFD8.tif             ^
  ..\testdata\BigTIFFSamples\Classic.tif                    ^
                                                            ^
  ..\testdata\www.fileformat.info\CCITT_1.TIF               ^
  ..\testdata\www.fileformat.info\CCITT_2.TIF               ^
  ..\testdata\www.fileformat.info\CCITT_3.TIF               ^
  ..\testdata\www.fileformat.info\CCITT_4.TIF               ^
  ..\testdata\www.fileformat.info\CCITT_5.TIF               ^
  ..\testdata\www.fileformat.info\CCITT_6.TIF               ^
  ..\testdata\www.fileformat.info\CCITT_7.TIF               ^
  ..\testdata\www.fileformat.info\CCITT_8.TIF               ^
  ..\testdata\www.fileformat.info\FLAG_T24.TIF              ^
  ..\testdata\www.fileformat.info\G31D.TIF                  ^
  ..\testdata\www.fileformat.info\G31DS.TIF                 ^
  ..\testdata\www.fileformat.info\G32D.TIF                  ^
  ..\testdata\www.fileformat.info\G32DS.TIF                 ^
  ..\testdata\www.fileformat.info\G4.TIF                    ^
  ..\testdata\www.fileformat.info\G4S.TIF                   ^
  ..\testdata\www.fileformat.info\GMARBLES.TIF              ^
  ..\testdata\www.fileformat.info\MARBIBM.TIF               ^
  ..\testdata\www.fileformat.info\MARBLES.TIF               ^
  ..\testdata\www.fileformat.info\XING_T24.TIF              ^
                                                            ^
  ..\testdata\multipage\multipage-gopher.tif                ^
  ..\testdata\multipage\multipage-sample.tif                ^
                                                            ^
  > tiffinfo-out.txt


