// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
*/
import "C"
import "reflect"
import "unsafe"

// Define alpha as the opacity of a surface
const (
	ALPHA_OPAQUE      = C.SDL_ALPHA_OPAQUE
	ALPHA_TRANSPARENT = C.SDL_ALPHA_TRANSPARENT
)

type PixelFormatEnum uint32

var (
	PIXELFORMAT_UNKNOWN     PixelFormatEnum = C.SDL_PIXELFORMAT_UNKNOWN
	PIXELFORMAT_INDEX1LSB   PixelFormatEnum = C.SDL_PIXELFORMAT_INDEX1LSB
	PIXELFORMAT_INDEX1MSB   PixelFormatEnum = C.SDL_PIXELFORMAT_INDEX1MSB
	PIXELFORMAT_INDEX4LSB   PixelFormatEnum = C.SDL_PIXELFORMAT_INDEX4LSB
	PIXELFORMAT_INDEX4MSB   PixelFormatEnum = C.SDL_PIXELFORMAT_INDEX4MSB
	PIXELFORMAT_INDEX8      PixelFormatEnum = C.SDL_PIXELFORMAT_INDEX8
	PIXELFORMAT_RGB332      PixelFormatEnum = C.SDL_PIXELFORMAT_RGB332
	PIXELFORMAT_RGB444      PixelFormatEnum = C.SDL_PIXELFORMAT_RGB444
	PIXELFORMAT_RGB555      PixelFormatEnum = C.SDL_PIXELFORMAT_RGB555
	PIXELFORMAT_BGR555      PixelFormatEnum = C.SDL_PIXELFORMAT_BGR555
	PIXELFORMAT_ARGB4444    PixelFormatEnum = C.SDL_PIXELFORMAT_ARGB4444
	PIXELFORMAT_RGBA4444    PixelFormatEnum = C.SDL_PIXELFORMAT_RGBA4444
	PIXELFORMAT_ABGR4444    PixelFormatEnum = C.SDL_PIXELFORMAT_ABGR4444
	PIXELFORMAT_BGRA4444    PixelFormatEnum = C.SDL_PIXELFORMAT_BGRA4444
	PIXELFORMAT_ARGB1555    PixelFormatEnum = C.SDL_PIXELFORMAT_ARGB1555
	PIXELFORMAT_RGBA5551    PixelFormatEnum = C.SDL_PIXELFORMAT_RGBA5551
	PIXELFORMAT_ABGR1555    PixelFormatEnum = C.SDL_PIXELFORMAT_ABGR1555
	PIXELFORMAT_BGRA5551    PixelFormatEnum = C.SDL_PIXELFORMAT_BGRA5551
	PIXELFORMAT_RGB565      PixelFormatEnum = C.SDL_PIXELFORMAT_RGB565
	PIXELFORMAT_BGR565      PixelFormatEnum = C.SDL_PIXELFORMAT_BGR565
	PIXELFORMAT_RGB24       PixelFormatEnum = C.SDL_PIXELFORMAT_RGB24
	PIXELFORMAT_BGR24       PixelFormatEnum = C.SDL_PIXELFORMAT_BGR24
	PIXELFORMAT_RGB888      PixelFormatEnum = C.SDL_PIXELFORMAT_RGB888
	PIXELFORMAT_RGBX8888    PixelFormatEnum = C.SDL_PIXELFORMAT_RGBX8888
	PIXELFORMAT_BGR888      PixelFormatEnum = C.SDL_PIXELFORMAT_BGR888
	PIXELFORMAT_BGRX8888    PixelFormatEnum = C.SDL_PIXELFORMAT_BGRX8888
	PIXELFORMAT_ARGB8888    PixelFormatEnum = C.SDL_PIXELFORMAT_ARGB8888
	PIXELFORMAT_RGBA8888    PixelFormatEnum = C.SDL_PIXELFORMAT_RGBA8888
	PIXELFORMAT_ABGR8888    PixelFormatEnum = C.SDL_PIXELFORMAT_ABGR8888
	PIXELFORMAT_BGRA8888    PixelFormatEnum = C.SDL_PIXELFORMAT_BGRA8888
	PIXELFORMAT_ARGB2101010 PixelFormatEnum = C.SDL_PIXELFORMAT_ARGB2101010

	PIXELFORMAT_YV12 PixelFormatEnum = C.SDL_PIXELFORMAT_YV12
	PIXELFORMAT_IYUV PixelFormatEnum = C.SDL_PIXELFORMAT_IYUV
	PIXELFORMAT_YUY2 PixelFormatEnum = C.SDL_PIXELFORMAT_YUY2
	PIXELFORMAT_UYVY PixelFormatEnum = C.SDL_PIXELFORMAT_UYVY
	PIXELFORMAT_YVYU PixelFormatEnum = C.SDL_PIXELFORMAT_YVYU
)

// Colors gives you access to palette's color data as a Color slice.  The
// slice is backed by palettes color data and should not be appended to.
func (palette *Palette) Colors() []Color {
	colors := []Color{}

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&colors))
	sh.Data = palette.colors
	sh.Len = int(palette.ncolors)
	sh.Cap = sh.Len

	return colors
}

// GetPixelFormatName gets the human readable name of a pixel format
func GetPixelFormatName(format PixelFormatEnum) string {
	return C.GoString(C.SDL_GetPixelFormatName(C.Uint32(format)))
}

// PixelFormatEnumToMasks converts format into a bpp and RGBA masks.
func PixelFormatEnumToMasks(format PixelFormatEnum) (bpp int, rmask, gmask, bmask, amask uint32, err error) {
	result := C.SDL_PixelFormatEnumToMasks(C.Uint32(format), (*C.int)(unsafe.Pointer(&bpp)),
		(*C.Uint32)(&rmask), (*C.Uint32)(&gmask), (*C.Uint32)(&bmask),
		(*C.Uint32)(&amask))
	if result == C.SDL_FALSE {
		err = sdlError(int(result))
	}
	return
}

// MasksToPixelFormatEnum converts a bpp and RGBA masks to a PixelFormatEnum.
func MasksToPixelFormatEnum(bpp int, rmask, gmask, bmask, amask uint32) PixelFormatEnum {
	return PixelFormatEnum(C.SDL_MasksToPixelFormatEnum(C.int(bpp), C.Uint32(rmask), C.Uint32(gmask),
		C.Uint32(bmask), C.Uint32(amask)))
}

// AllocFormat creates a PixelFormat structure from a PixelFormatEnum.
func AllocFormat(format PixelFormatEnum) (*PixelFormat, error) {
	r := (*PixelFormat)(unsafe.Pointer(C.SDL_AllocFormat(C.Uint32(format))))
	if r == nil {
		return nil, sdlError(0)
	}
	return r, nil
}

// Free frees a PixelFormat created by AllocFormat.
func (format *PixelFormat) Free() {
	C.SDL_FreeFormat((*C.SDL_PixelFormat)(unsafe.Pointer(format)))
}

// AllocPalette create a palette structure with the specified number of color
// entries.
func AllocPalette(ncolors int) (*Palette, error) {
	r := (*Palette)(unsafe.Pointer(C.SDL_AllocPalette(C.int(ncolors))))
	if r == nil {
		return nil, sdlError(0)
	}
	return r, nil
}

// SetPalette sets the palette for format.
func (format *PixelFormat) SetPalette(palette *Palette) error {
	r := C.SDL_SetPixelFormatPalette((*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		(*C.SDL_Palette)(unsafe.Pointer(palette)))
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// SetColors sets a range of colors in a palette.
func (palette *Palette) SetColors(colors []Color) error {
	var ptr *C.SDL_Color
	if len(colors) > 0 {
		ptr = (*C.SDL_Color)(unsafe.Pointer(&colors[0]))
	}

	r := C.SDL_SetPaletteColors((*C.SDL_Palette)(unsafe.Pointer(palette)),
		ptr, 0, C.int(len(colors)))
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// Free frees a palette created with AllocPalette.
func (palette *Palette) Free() {
	C.SDL_FreePalette((*C.SDL_Palette)(unsafe.Pointer(palette)))
}

// MapRGB maps an RGB triple to an opaque pixel value for a given pixel format.
func MapRGB(format *PixelFormat, r, g, b uint8) uint32 {
	return uint32(C.SDL_MapRGB((*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		C.Uint8(r), C.Uint8(g), C.Uint8(b)))
}

// MapRGBA maps an RGBA quadruple to a pixel value for a given pixel format.
func MapRGBA(format *PixelFormat, r, g, b, a uint8) uint32 {
	return uint32(C.SDL_MapRGBA((*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)))
}

// GetRGB gets the RGB components from a pixel of the specified format.
func GetRGB(pixel uint32, format *PixelFormat) (r, g, b uint8) {
	C.SDL_GetRGB(C.Uint32(pixel), (*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		(*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b))
	return
}

// GetRGBA gets the RGBA components from a pixel of the specified format.
func GetRGBA(pixel uint32, format *PixelFormat) (r, g, b, a uint8) {
	C.SDL_GetRGBA(C.Uint32(pixel), (*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		(*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b), (*C.Uint8)(&a))
	return
}

// CalculateGammaRamp calculates a 256 entry gamma ramp for a gamma value.
func CalculateGammaRamp(gamma float32, ramp *[256]uint16) {
	C.SDL_CalculateGammaRamp(C.float(gamma), (*C.Uint16)(unsafe.Pointer(&ramp[0])))
}
