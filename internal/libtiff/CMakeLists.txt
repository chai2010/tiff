# Copyright 2013 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

project(LIBTIFF)

#------------------------------------------------------------------------------

IF(WIN32)
  if(CMAKE_SIZEOF_VOID_P EQUAL 8)
    set(OS win64)
  else()
    set(OS win32)
  endif()
else()
  if(CMAKE_SIZEOF_VOID_P EQUAL 8)
    set(OS posix64)
  else()
    set(OS posix64)
  endif()
endif()

#------------------------------------------------------------------------------

add_definitions(
  #-DXXX
)

include_directories(AFTER
	./3rdparty/zlib
	./libtiff
)

set(TIFF_SRC
	./libtiff/tif_aux.c
	./libtiff/tif_close.c
	./libtiff/tif_codec.c
	./libtiff/tif_color.c
	./libtiff/tif_compress.c
	./libtiff/tif_dir.c
	./libtiff/tif_dirinfo.c
	./libtiff/tif_dirread.c
	./libtiff/tif_dirwrite.c
	./libtiff/tif_dumpmode.c
	./libtiff/tif_error.c
	./libtiff/tif_extension.c
	./libtiff/tif_fax3.c
	./libtiff/tif_fax3sm.c
	./libtiff/tif_getimage.c
	./libtiff/tif_jpeg.c
	./libtiff/tif_ojpeg.c
	./libtiff/tif_flush.c
	./libtiff/tif_luv.c
	./libtiff/tif_lzw.c
	./libtiff/tif_next.c
	./libtiff/tif_open.c
	./libtiff/tif_packbits.c
	./libtiff/tif_pixarlog.c
	./libtiff/tif_predict.c
	./libtiff/tif_print.c
	./libtiff/tif_read.c
	#./libtiff/tif_stream.cxx
	./libtiff/tif_swab.c
	./libtiff/tif_strip.c
	./libtiff/tif_thunder.c
	./libtiff/tif_tile.c
	./libtiff/tif_version.c
	./libtiff/tif_warning.c
	./libtiff/tif_write.c
	./libtiff/tif_zip.c

	./libtiff/tif_win32.c

	./3rdparty/zlib/adler32.c
	./3rdparty/zlib/compress.c
	./3rdparty/zlib/crc32.c
	./3rdparty/zlib/deflate.c
	./3rdparty/zlib/gzclose.c
	./3rdparty/zlib/gzlib.c
	./3rdparty/zlib/gzread.c
	./3rdparty/zlib/gzwrite.c
	./3rdparty/zlib/infback.c
	./3rdparty/zlib/inflate.c 
	./3rdparty/zlib/inftrees.c 
	./3rdparty/zlib/inffast.c
	./3rdparty/zlib/trees.c
	./3rdparty/zlib/uncompr.c
	./3rdparty/zlib/zutil.c
)


#------------------------------------------------------------------------------

add_library(tifflib STATIC
  ${TIFF_SRC}
)
if(CMAKE_BUILD_TYPE STREQUAL "debug")
  set_target_properties(tifflib
    PROPERTIES OUTPUT_NAME "tiff-${OS}-debug"
  )
else()
  set_target_properties(tifflib
    PROPERTIES OUTPUT_NAME "tiff-${OS}"
  )
endif()

install(TARGETS tifflib
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------
# Tiff Demo

include_directories(AFTER
  ./libtiff
)

# webpapp
add_executable(tiffapp
  ./demo/demo.c

  ${TIFF_SRC}
)
install(TARGETS tiffapp
  RUNTIME DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  LIBRARY DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
  ARCHIVE DESTINATION ${CMAKE_CURRENT_SOURCE_DIR}
)

#------------------------------------------------------------------------------

