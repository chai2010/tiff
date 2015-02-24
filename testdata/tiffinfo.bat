:: Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
:: Use of this source code is governed by a BSD-style
:: license that can be found in the LICENSE file.

@setlocal

@cd %~dp0

@go run ../apps/tiffinfo.go                                 ^
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
  > z_testdata-info.txt


@go run ../apps/tiffinfo.go                                 ^
  ..\testdata\BigTIFFSamples\BigTIFF.tif                    ^
  ..\testdata\BigTIFFSamples\BigTIFFLong.tif                ^
  ..\testdata\BigTIFFSamples\BigTIFFLong8.tif               ^
  ..\testdata\BigTIFFSamples\BigTIFFLong8Tiles.tif          ^
  ..\testdata\BigTIFFSamples\BigTIFFMotorola.tif            ^
  ..\testdata\BigTIFFSamples\BigTIFFMotorolaLongStrips.tif  ^
  ..\testdata\BigTIFFSamples\BigTIFFSubIFD4.tif             ^
  ..\testdata\BigTIFFSamples\BigTIFFSubIFD8.tif             ^
  ..\testdata\BigTIFFSamples\Classic.tif                    ^
  > z_testdata-BigTIFFSamples-info.txt

@go run ../apps/tiffinfo.go                                 ^
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
  > z_testdata-www.fileformat.info-info.txt

@go run ../apps/tiffinfo.go                                 ^
  ..\testdata\multipage\multipage-gopher.tif                ^
  ..\testdata\multipage\multipage-sample.tif                ^
  > z_testdata-multipage-info.txt


@go run ../apps/tiffinfo.go                                    ^
  ..\testdata\geotiff\gdal_eg\cea.tif                          ^
  ..\testdata\geotiff\GeogToWGS84GeoKey\GeogToWGS84GeoKey5.tif ^
  ..\testdata\geotiff\intergraph\albers27.tif                  ^
  ..\testdata\geotiff\intergraph\azi_equi.tif                  ^
  ..\testdata\geotiff\intergraph\bip_obl_con.tif               ^
  ..\testdata\geotiff\intergraph\bonne.tif                     ^
  ..\testdata\geotiff\intergraph\brit_nat.tif                  ^
  ..\testdata\geotiff\intergraph\cass_sold.tif                 ^
  ..\testdata\geotiff\intergraph\cham_tri.tif                  ^
  ..\testdata\geotiff\intergraph\cyl_equi.tif                  ^
  ..\testdata\geotiff\intergraph\eckert4.tif                   ^
  ..\testdata\geotiff\intergraph\equi_conic.tif                ^
  ..\testdata\geotiff\intergraph\gauss_k.tif                   ^
  ..\testdata\geotiff\intergraph\gen_pers.tif                  ^
  ..\testdata\geotiff\intergraph\geoc.tif                      ^
  ..\testdata\geotiff\intergraph\geog.tif                      ^
  ..\testdata\geotiff\intergraph\gnomonic.tif                  ^
  ..\testdata\geotiff\intergraph\indo_poly.tif                 ^
  ..\testdata\geotiff\intergraph\krovak.tif                    ^
  ..\testdata\geotiff\intergraph\laborde.tif                   ^
  ..\testdata\geotiff\intergraph\lamb_az.tif                   ^
  ..\testdata\geotiff\intergraph\lamb_conf.tif                 ^
  ..\testdata\geotiff\intergraph\lamb_conf_obl.tif             ^
  ..\testdata\geotiff\intergraph\lsr.tif                       ^
  ..\testdata\geotiff\intergraph\merc.tif                      ^
  ..\testdata\geotiff\intergraph\merc_obl.tif                  ^
  ..\testdata\geotiff\intergraph\miller.tif                    ^
  ..\testdata\geotiff\intergraph\mod_poly.tif                  ^
  ..\testdata\geotiff\intergraph\molle.tif                     ^
  ..\testdata\geotiff\intergraph\nz_map_grid.tif               ^
  ..\testdata\geotiff\intergraph\ortho.tif                     ^
  ..\testdata\geotiff\intergraph\poly.tif                      ^
  ..\testdata\geotiff\intergraph\rect_grid.tif                 ^
  ..\testdata\geotiff\intergraph\rect_skew_ortho.tif           ^
  ..\testdata\geotiff\intergraph\robinson.tif                  ^
  ..\testdata\geotiff\intergraph\sinus.tif                     ^
  ..\testdata\geotiff\intergraph\spc_obl_merc.tif              ^
  ..\testdata\geotiff\intergraph\spc_obl_merc_pre.tif          ^
  ..\testdata\geotiff\intergraph\spcs27.tif                    ^
  ..\testdata\geotiff\intergraph\spcs83.tif                    ^
  ..\testdata\geotiff\intergraph\stereo.tif                    ^
  ..\testdata\geotiff\intergraph\stereo3.tif                   ^
  ..\testdata\geotiff\intergraph\stereo_np.tif                 ^
  ..\testdata\geotiff\intergraph\stereo_sp.tif                 ^
  ..\testdata\geotiff\intergraph\stereo_up.tif                 ^
  ..\testdata\geotiff\intergraph\trans_merc.tif                ^
  ..\testdata\geotiff\intergraph\utm.tif                       ^
  ..\testdata\geotiff\intergraph\van_der_grin.tif              ^
  ..\testdata\geotiff\other\erdas_spnad83.tif                  ^
  ..\testdata\geotiff\pci_eg\acea.tif                          ^
  ..\testdata\geotiff\pci_eg\ae.tif                            ^
  ..\testdata\geotiff\pci_eg\econic.tif                        ^
  ..\testdata\geotiff\pci_eg\er.tif                            ^
  ..\testdata\geotiff\pci_eg\gnomonic.tif                      ^
  ..\testdata\geotiff\pci_eg\laea.tif                          ^
  ..\testdata\geotiff\pci_eg\latlong.tif                       ^
  ..\testdata\geotiff\pci_eg\lcc-27.tif                        ^
  ..\testdata\geotiff\pci_eg\mc.tif                            ^
  ..\testdata\geotiff\pci_eg\mercator.tif                      ^
  ..\testdata\geotiff\pci_eg\meter.tif                         ^
  ..\testdata\geotiff\pci_eg\oblqmerc.tif                      ^
  ..\testdata\geotiff\pci_eg\og.tif                            ^
  ..\testdata\geotiff\pci_eg\pc.tif                            ^
  ..\testdata\geotiff\pci_eg\ps.tif                            ^
  ..\testdata\geotiff\pci_eg\rob.tif                           ^
  ..\testdata\geotiff\pci_eg\sg.tif                            ^
  ..\testdata\geotiff\pci_eg\sin.tif                           ^
  ..\testdata\geotiff\pci_eg\spaf27.tif                        ^
  ..\testdata\geotiff\pci_eg\spcs27.tif                        ^
  ..\testdata\geotiff\pci_eg\spif83.tif                        ^
  ..\testdata\geotiff\pci_eg\tm.tif                            ^
  ..\testdata\geotiff\pci_eg\utm11-27.tif                      ^
  ..\testdata\geotiff\pci_eg\vdg.tif                           ^
  ..\testdata\geotiff\spot\chicago\SP27GTIF.TIF                ^
  ..\testdata\geotiff\spot\chicago\UTM2GTIF.TIF                ^
  ..\testdata\geotiff\zi_imaging\image0.tif                    ^
  ..\testdata\geotiff\zi_imaging\image1.tif                    ^
  ..\testdata\geotiff\zi_imaging\image2.tif                    ^
  ..\testdata\geotiff\zi_imaging\image3.tif                    ^
  ..\testdata\geotiff\zi_imaging\image4.tif                    ^
  ..\testdata\geotiff\zi_imaging\image5.tif                    ^
  ..\testdata\geotiff\zi_imaging\image6.tif                    ^
  ..\testdata\geotiff\zi_imaging\image7.tif                    ^
  ..\testdata\geotiff\zi_imaging\tp_image0.tif                 ^
  ..\testdata\geotiff\zi_imaging\tp_image1.tif                 ^
  ..\testdata\geotiff\zi_imaging\tp_image2.tif                 ^
  ..\testdata\geotiff\zi_imaging\tp_image3.tif                 ^
  ..\testdata\geotiff\zi_imaging\tp_image4.tif                 ^
  ..\testdata\geotiff\zi_imaging\tp_image5.tif                 ^
  ..\testdata\geotiff\zi_imaging\tp_image6.tif                 ^
  ..\testdata\geotiff\zi_imaging\tp_image7.tif                 ^
  > z_testdata-geotiff-info.txt



