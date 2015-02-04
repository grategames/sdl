// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
*/
import "C"

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"reflect"
	"unsafe"
)

// Pixels gives you access to surf's pixel data as a byte slice.  This slice
// should not be appended too.  Before accessing the pixel data you may need
// to lock surf.  Pixels returns an empty slice if surf has no pixel data.
func (surf *Surface) Pixels() []byte {
	if unsafe.Pointer(surf.pixels) == nil {
		return []byte{}
	}

	pixels := []byte{}

	sh := (*reflect.SliceHeader)(unsafe.Pointer(&pixels))
	sh.Data = surf.pixels
	sh.Len = int(surf.H * surf.Pitch)
	sh.Cap = sh.Len

	return pixels
}

// MustLock returns true if surf needs to be locked before access.
func (surf *Surface) MustLock() bool {
	return (surf.flags & C.SDL_RLEACCEL) != 0
}

// CreateRGBSurface allocates an RGB surface.  If the depth is 4 or 8 bits, an
// empty palette is allocated for the surface.  If the depth is greater then 8
// bits, the pixel format is set using the flags '[RGB]mask'. Using zeros for
// '[RGB]mask' will set them to their default value based on depth.  However,
// using zero for the Amask results in an Amask of zero.
func CreateRGBSurface(width, height, depth int, Rmask, Gmask, Bmask, Amask uint32) (*Surface, error) {
	r := C.SDL_CreateRGBSurface(0, C.int(width), C.int(height),
		C.int(depth), C.Uint32(Rmask), C.Uint32(Gmask),
		C.Uint32(Bmask), C.Uint32(Amask))
	if r == nil {
		return nil, sdlError(0)
	}
	return (*Surface)(unsafe.Pointer(r)), nil
}

// CreateRGBSurfaceFrom allocates an RGB surface with existing pixel data. If
// the depth is 4 or 8, an empty palette is allocated for the surface.  If the
// depth is greater then 8 bits, the pixel format is set using the flags
// '[RGB]mask'.  Using zeros for '[RGB]mask' will set them to their default
// value based on depth.  However, using zero for the Amask results in an
// Amask of zero.
//
// The pixel data is not copied and it is your responsibility to keep pixels
// from being garbage collected too early.
func CreateRGBSurfaceFrom(pixels unsafe.Pointer, width, height, depth, pitch int, Rmask, Gmask, Bmask, Amask uint32) (*Surface, error) {
	r := C.SDL_CreateRGBSurfaceFrom(pixels, C.int(width), C.int(height),
		C.int(depth), C.int(pitch), C.Uint32(Rmask), C.Uint32(Gmask),
		C.Uint32(Bmask), C.Uint32(Amask))
	if r == nil {
		return nil, sdlError(0)
	}
	return (*Surface)(unsafe.Pointer(r)), nil
}

// Free frees surf.
func (surf *Surface) Free() {
	C.SDL_FreeSurface((*C.SDL_Surface)(unsafe.Pointer(surf)))
}

// SetPalette sets the palette used by surf.  SetPalette returns an error if
// surf's pixel format does not use a palette.
//
// Note: A single palette can be shared with many surfaces.
func (surf *Surface) SetPalette(palette *Palette) error {
	r := C.SDL_SetSurfacePalette((*C.SDL_Surface)(unsafe.Pointer(surf)),
		(*C.SDL_Palette)(unsafe.Pointer(surf)))
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// Lock sets up surf for directly accessing the pixels.
//
// Between calls to Lock and Unlock, you can write to and read from surf pixel
// data, using the pixel format stored in surf.Format.  Once you are done
// accessing the surface, you should use Unlock to release surf.
//
// Not all surfaces require locking.  If MUSTLOCK(surf) evaluates to true,
// then you can read and write to the surface at any time, and the pixel
// format of the surface will not change.
//
// No operating system or library calls should be made between lock/unlock
// pairs, as critical system locks may be held during this time.
func (surf *Surface) Lock() error {
	r := C.SDL_LockSurface((*C.SDL_Surface)(unsafe.Pointer(surf)))
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// Unlock releases the lock on surf.  See Lock for more information.
func (surf *Surface) Unlock() {
	C.SDL_UnlockSurface((*C.SDL_Surface)(unsafe.Pointer(surf)))
}

/*
Skipping
SDL_LoadBMP_RW()
SDL_LoadBMP()
SDL_SaveBMP_RW()
SDL_SaveBMP()
Undecided on exposing RWops
*/

// SetRLE enables RLE accleration for surf if flag is true, disables RLE
// accleration if flag is false.
//
// Note: If RLE is enabled, colorkey and alpha blending blits are much faster,
// but the surface must be locked before directly accessing the pixels.
func (surf *Surface) SetRLE(flag bool) error {
	var f C.int
	if flag {
		f = 1
	}

	r := C.SDL_SetSurfaceRLE((*C.SDL_Surface)(unsafe.Pointer(surf)), f)
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// SetColorKey the color key (transparent pixel) in surf.  If flag is true
// colorkey is enabled, pixels matching key will be transparent on blit, if
// flag is false colorkey is disabled.  You can use MapRGB to generate key.
func (surf *Surface) SetColorKey(flag bool, key uint32) error {
	var f C.int
	if flag {
		f = 1
	}

	r := C.SDL_SetColorKey((*C.SDL_Surface)(unsafe.Pointer(surf)), f,
		C.Uint32(key))
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// GetColorKey gets the color key (transparent pixel) in surf.
func (surf *Surface) GetColorKey() (key uint32, err error) {
	r := C.SDL_GetColorKey((*C.SDL_Surface)(unsafe.Pointer(surf)),
		(*C.Uint32)(&key))
	if r != 0 {
		err = sdlError(int(r))
	}
	return
}

// SetColorMod sets an additional color value used in blit operations.
func (surf *Surface) SetColorMod(r, g, b uint8) error {
	i := C.SDL_SetSurfaceColorMod((*C.SDL_Surface)(unsafe.Pointer(surf)),
		C.Uint8(r), C.Uint8(g), C.Uint8(b))
	if i != 0 {
		return sdlError(int(i))
	}
	return nil
}

// GetColorMod gets the additional color value used in blit operations.
func (surf *Surface) GetColorMod() (r, g, b uint8, err error) {
	i := C.SDL_GetSurfaceColorMod((*C.SDL_Surface)(unsafe.Pointer(surf)),
		(*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b))
	if i != 0 {
		err = sdlError(int(i))
	}
	return
}

// SetAlphaMod sets an additional alpha value used in blit operations.
func (surf *Surface) SetAlphaMod(alpha uint8) error {
	r := C.SDL_SetSurfaceAlphaMod((*C.SDL_Surface)(unsafe.Pointer(surf)), C.Uint8(alpha))
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// GetAlphaMod gets the additional alpha value used in blit operations.
func (surf *Surface) GetAlphaMod() (alpha uint8, err error) {
	r := C.SDL_GetSurfaceAlphaMod((*C.SDL_Surface)(unsafe.Pointer(surf)),
		(*C.Uint8)(&alpha))
	if r != 0 {
		err = sdlError(int(r))
	}
	return
}

// SetBlendMode sets the blend mode used for blit operations.
func (surf *Surface) SetBlendMode(blendMode BlendMode) error {
	r := C.SDL_SetSurfaceBlendMode((*C.SDL_Surface)(unsafe.Pointer(surf)),
		C.SDL_BlendMode(blendMode))
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// GetBlendMode gets the blend mode used for blit operations.
func (surf *Surface) GetBlendMode() (blendMode BlendMode, err error) {
	r := C.SDL_GetSurfaceBlendMode((*C.SDL_Surface)(unsafe.Pointer(surf)),
		(*C.SDL_BlendMode)(&blendMode))
	if r != 0 {
		err = sdlError(int(r))
	}
	return
}

// SetClipRect sets the clipping rectangle for the destination surface in a
// blit.
//
// If rect is nil, clipping will be disabled.
//
// If the clip rectangle does not intersect the surface, SetClipRect will
// return false, and blits will be completely clipped.  Otherwise
// SetClipRect returns true and blits to surf will be clipped to the
// intersection of the surface area and clipping rectangle.
//
// Note: Blits are automatically clipped to the edges of the source and
// destination surfaces.
func (surf *Surface) SetClipRect(rect *Rect) bool {
	result := C.SDL_SetClipRect((*C.SDL_Surface)(unsafe.Pointer(surf)),
		(*C.SDL_Rect)(unsafe.Pointer(rect)))
	if result == C.SDL_TRUE {
		return true
	}
	return false
}

// GetClipRect gets the clipping rectangle for the destination surface in a
// blit.
func (surf *Surface) GetClipRect() *Rect {
	rect := new(Rect)
	C.SDL_GetClipRect((*C.SDL_Surface)(unsafe.Pointer(surf)),
		(*C.SDL_Rect)(unsafe.Pointer(rect)))
	return rect
}

// Convert creates a new surface of the specified PixelFormat, and then copies
// and maps the given surface to it so the blit of the converted surface will
// be as fast as possible.
func (surf *Surface) Convert(fmt *PixelFormat) (*Surface, error) {
	s := C.SDL_ConvertSurface((*C.SDL_Surface)(unsafe.Pointer(surf)),
		(*C.SDL_PixelFormat)(unsafe.Pointer(fmt)), 0)
	if s == nil {
		return nil, sdlError(0)
	}
	return (*Surface)(unsafe.Pointer(s)), nil
}

// ConvertFormat creates a new surface of the specified PixelFormatEnum, and
// then copies and maps the given surface to it so the blit of the converted
// surface will be as fast as possible.
func (surf *Surface) ConvertFormat(pixel_format PixelFormatEnum) (*Surface, error) {
	s := C.SDL_ConvertSurfaceFormat((*C.SDL_Surface)(unsafe.Pointer(surf)),
		C.Uint32(pixel_format), 0)
	if s == nil {
		return nil, sdlError(0)
	}
	return (*Surface)(unsafe.Pointer(s)), nil
}

/*
Skipping
SDL_ConvertPixels()
Unsafe
*/

// FillRect performs a fast fill of the given rectangle with color.
//
// If rect is nil, the whole surface will be filled with color.
//
// The color should be a pixel of the format used by the surface, and can be
// generated by the MapRGB function.
func (surf *Surface) FillRect(rect *Rect, color uint32) error {
	r := C.SDL_FillRect((*C.SDL_Surface)(unsafe.Pointer(surf)),
		(*C.SDL_Rect)(unsafe.Pointer(rect)), C.Uint32(color))
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// FillRect performs a fast fill of the given rectangles with color.
//
// The color should be a pixel of the format used by the surface, and can be
// generated by the MapRGB function.
func (surf *Surface) FillRects(rects []Rect, color uint32) error {
	var ptr *C.SDL_Rect
	if len(rects) > 0 {
		ptr = (*C.SDL_Rect)(unsafe.Pointer(&rects[0]))
	}
	r := C.SDL_FillRects((*C.SDL_Surface)(unsafe.Pointer(surf)), ptr,
		C.int(len(rects)), C.Uint32(color))
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// Blit preforms a fast blit from src to dst, srcrect is the rectangle copied
// from src to dst, dstrect positions the blit on dst.  Blit will change
// dstrects W and H to match the width and height of the blit.
//
// Blit should not be called on a locked surface.
func (dst *Surface) Blit(src *Surface, srcrect, dstrect *Rect) error {
	r := C.SDL_UpperBlit((*C.SDL_Surface)(unsafe.Pointer(src)),
		(*C.SDL_Rect)(unsafe.Pointer(srcrect)), (*C.SDL_Surface)(unsafe.Pointer(dst)),
		(*C.SDL_Rect)(unsafe.Pointer(dstrect)))
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

/*
Skipping
SDL_LowerBlit()
Seems to be mostly for internal use.
*/

// BlitScaled performs a scaled blit from src to dst, srcrect is the rectangle
// copied from src to dst, dstrect is the rectangle copied to.  The area
// copied from src will be scaled to the width and height of dstrect.
// BlitScaled does not modify dstrect.
//
// BlitScaled should not be called on a locked surface.
func (dst *Surface) BlitScaled(src *Surface, srcrect, dstrect *Rect) error {
	r := C.SDL_UpperBlitScaled((*C.SDL_Surface)(unsafe.Pointer(src)),
		(*C.SDL_Rect)(unsafe.Pointer(srcrect)), (*C.SDL_Surface)(unsafe.Pointer(dst)),
		(*C.SDL_Rect)(unsafe.Pointer(dstrect)))
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

/*
Skipping
SDL_LowerBlitScaled()
Seems to be mostly for internal use.
*/

// CreateRGBSurfaceFromImage creates an RGB surface from a image.Image.
//
// If img is an *image.RGBA, *image.RGBA64, *image.NRGBA, or *image.NRGBA64,
// CreateRGBSurfaceFromImage will create a 32 bit surface with an alpha
// channel.
//
// If img is an *image.Paletted, CreateRGBSurfaceFromImage will create an 8 bit
// surface and copy the image palette to the surface palette.  If the colors
// in img.Palette are color.NRGBA, or color.NRGBA64 the colors in the surface's
// palette will not be alpha-premultiplied otherwise they will.
//
// If img is any other type an error is returned.
// 
// The pixel data is copied to the surface.
func CreateRGBSurfaceFromImage(img image.Image) (*Surface, error) {
	switch i := img.(type) {
	case *image.NRGBA:
		rect := i.Bounds()

		surf, err := CreateRGBSurface(rect.Dx(), rect.Dy(), 32,
			0x000000FF, 0x0000FF00, 0x00FF0000, 0xFF000000)
		if err != nil {
			return nil, err
		}

		copy(surf.Pixels(), i.Pix)
		return surf, nil
	case *image.RGBA, *image.RGBA64, *image.NRGBA64:
		rect := img.Bounds()

		surf, err := CreateRGBSurface(rect.Dx(), rect.Dy(), 32,
			0x000000FF, 0x0000FF00, 0x00FF0000, 0xFF000000)
		if err != nil {
			return nil, err
		}

		data := image.NewNRGBA(rect)
		switch i := img.(type) {
		case *image.RGBA:
			draw.Draw(data, data.Bounds(), i, rect.Min, draw.Src)
		case *image.RGBA64:
			draw.Draw(data, data.Bounds(), i, rect.Min, draw.Src)
		case *image.NRGBA:
			draw.Draw(data, data.Bounds(), i, rect.Min, draw.Src)
		}

		copy(surf.Pixels(), data.Pix)
		return surf, nil
	case *image.Paletted:
		rect := i.Bounds()

		surf, err := CreateRGBSurface(rect.Dx(), rect.Dy(), 8,
			0, 0, 0, 0)
		if err != nil {
			return nil, err
		}

		copy(surf.Pixels(), i.Pix)

		colors := make([]Color, 0, len(i.Palette))
		col := Color{}
		for _, c := range i.Palette {
			switch c1 := c.(type) {
			case color.NRGBA:
				col.R = c1.R
				col.G = c1.G
				col.B = c1.B
			case color.NRGBA64:
				cm := color.NRGBAModel.Convert(c).(color.NRGBA)
				col.R = cm.R
				col.G = cm.G
				col.B = cm.B
			case color.RGBA:
				col.R = c1.R
				col.G = c1.G
				col.B = c1.B
			default:
				cm := color.RGBAModel.Convert(c).(color.RGBA)
				col.R = cm.R
				col.G = cm.G
				col.B = cm.B
			}
			colors = append(colors, col)
		}

		err = surf.Format.Palette.SetColors(colors)
		if err != nil {
			return nil, err
		}

		return surf, nil
	default:
		return nil, fmt.Errorf("unknown image type: %T", img)
	}
	panic("Unreachable")
}
