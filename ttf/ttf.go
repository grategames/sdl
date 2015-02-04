// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ttf provides bindings for SDL2_ttf.
//
// Usage
//
// The ttf package is not thread safe and all calls must be serialized in some
// way.  This can be done by only calling the ttf package from a single
// goroutine or wrapping the calls in a sync.Mutex. See
// http://bugzilla.libsdl.org/show_bug.cgi?id=1532 for more information.
package ttf

/*
#cgo pkg-config: sdl2
#cgo LDFLAGS: -lSDL2_ttf
#include "SDL_ttf.h"
*/
import "C"

import (
	"grate/backend/sdl2"
	"fmt"
	"unsafe"
)

func sdlError(i int) error {
	return sdl.SDLError{
		Msg:   sdl.GetError(),
		Value: i,
	}
}

var InvalidFont = fmt.Errorf("Invalid Font")

const (
	MAJOR_VERSION = C.SDL_TTF_MAJOR_VERSION
	MINOR_VERSION = C.SDL_TTF_MINOR_VERSION
	PATCHLEVEL    = C.SDL_TTF_PATCHLEVEL
)

// Version returns the compile-time version of the SDL_ttf library.
func Version() *sdl.Version {
	return &sdl.Version{
		Major: MAJOR_VERSION,
		Minor: MINOR_VERSION,
		Patch: PATCHLEVEL,
	}
}

// LinkedVersion returns the version of the dynamically linked SDL_ttf
// library.
func LinkedVersion() *sdl.Version {
	return (*sdl.Version)(unsafe.Pointer(C.TTF_Linked_Version()))
}

// Font contains font information.
type Font struct {
	font *C.TTF_Font
}

// Init initializes the TTF engine.
func Init() error {
	i := int(C.TTF_Init())
	if i != 0 {
		return sdlError(i)
	}
	return nil
}

// OpenFont opens a font file and creates a font of the specified point size.
// This can load .ttf or .fon files.
func OpenFont(file string, ptsize int) (Font, error) {
	cstr := C.CString(file)
	defer C.free(unsafe.Pointer(cstr))

	f := C.TTF_OpenFont(cstr, C.int(ptsize))
	if f == nil {
		return Font{}, sdlError(0)
	}
	return Font{f}, nil
}

// OpenFontIndex opens a font file and creates a font of the specified point
// size using the specified index. This can load .ttf or .fon files.
func OpenFontIndex(file string, ptsize, index int) (Font, error) {
	cstr := C.CString(file)
	defer C.free(unsafe.Pointer(cstr))

	f := C.TTF_OpenFontIndex(cstr, C.int(ptsize), C.long(index))
	if f == nil {
		return Font{}, sdlError(0)
	}
	return Font{f}, nil
}

type Style int

// Font style settings.
const (
	STYLE_NORMAL        Style = C.TTF_STYLE_NORMAL
	STYLE_BOLD          Style = C.TTF_STYLE_BOLD
	STYLE_ITALIC        Style = C.TTF_STYLE_ITALIC
	STYLE_UNDERLINE     Style = C.TTF_STYLE_UNDERLINE
	STYLE_STRIKETHROUGH Style = C.TTF_STYLE_STRIKETHROUGH
)

// GetStyle gets the style of f.
func (f Font) GetStyle() (Style, error) {
	if f.font == nil {
		return 0, InvalidFont
	}
	return Style(C.TTF_GetFontStyle(f.font)), nil
}

// SetStyle set the style of f.
func (f Font) SetStyle(style Style) error {
	if f.font == nil {
		return InvalidFont
	}
	C.TTF_SetFontStyle(f.font, C.int(style))
	return nil
}

// GetOutline gets the outline of f in pixels.
func (f Font) GetOutline() (int, error) {
	if f.font == nil {
		return 0, InvalidFont
	}
	return int(C.TTF_GetFontOutline(f.font)), nil
}

// SetOutline sets the outline of f in pixels.
func (f Font) SetOutline(outline int) error {
	if f.font == nil {
		return InvalidFont
	}
	C.TTF_SetFontOutline(f.font, C.int(outline))
	return nil
}

type Hint int

// FreeType hinter settings.
const (
	HINTING_NORMAL Hint = C.TTF_HINTING_NORMAL
	HINTING_LIGHT  Hint = C.TTF_HINTING_LIGHT
	HINTING_MONO   Hint = C.TTF_HINTING_MONO
	HINTING_NONE   Hint = C.TTF_HINTING_NONE
)

// GetHinting gets the FreeType hinter setting for f.
func (f Font) GetHinting() (Hint, error) {
	if f.font == nil {
		return 0, InvalidFont
	}
	return Hint(C.TTF_GetFontHinting(f.font)), nil
}

// SetHinting sets the FreeType hinter setting for f.
func (f Font) SetHinting(hinting Hint) error {
	if f.font == nil {
		return InvalidFont
	}
	C.TTF_SetFontHinting(f.font, C.int(hinting))
	return nil
}

// Height gets the total height of f.
func (f Font) Height() (int, error) {
	if f.font == nil {
		return 0, InvalidFont
	}
	return int(C.TTF_FontHeight(f.font)), nil
}

// Ascent gets the offset from the baseline to the top of f.  This is a
// positive value, relative to the baseline.
func (f Font) Ascent() (int, error) {
	if f.font == nil {
		return 0, InvalidFont
	}
	return int(C.TTF_FontAscent(f.font)), nil
}

// Descent gets the off from the baseline to the bottom of f.  This is a
// negative value, relative to the baseline.
func (f Font) Descent() (int, error) {
	if f.font == nil {
		return 0, InvalidFont
	}
	return int(C.TTF_FontDescent(f.font)), nil
}

// LineSkip gets the recommended spacing between lines of text for f.
func (f Font) LineSkip() (int, error) {
	if f.font == nil {
		return 0, InvalidFont
	}
	return int(C.TTF_FontLineSkip(f.font)), nil
}

// GetKerning returns whether or not kerning is allowed for f.
func (f Font) GetKerning() (bool, error) {
	if f.font == nil {
		return false, InvalidFont
	}
	if C.TTF_GetFontKerning(f.font) == 0 {
		return false, nil
	}
	return true, nil
}

// SetKerning sets whether or not kerning is allowed for f.
func (f Font) SetKerning(allowed bool) error {
	if f.font == nil {
		return InvalidFont
	}
	var i C.int
	if allowed {
		i = 1
	}
	C.TTF_SetFontKerning(f.font, i)
	return nil
}

// Faces gets the number of faces of f.
func (f Font) Faces() (int, error) {
	if f.font == nil {
		return 0, InvalidFont
	}
	return int(C.TTF_FontFaces(f.font)), nil
}

// FaceIsFixedWidth checks if the current font face of f is a fixed width
// font.
func (f Font) FaceIsFixedWidth() (bool, error) {
	if f.font == nil {
		return false, InvalidFont
	}
	if C.TTF_FontFaceIsFixedWidth(f.font) == 0 {
		return false, nil
	}
	return true, nil
}

// FaceFamilyName returns the current font face family name of f. 
func (f Font) FaceFamilyName() (string, error) {
	if f.font == nil {
		return "", InvalidFont
	}
	return C.GoString(C.TTF_FontFaceFamilyName(f.font)), nil
}

// FaceStyleName returns the current font face style name of f.
func (f Font) FaceStyleName() (string, error) {
	if f.font == nil {
		return "", InvalidFont
	}
	return C.GoString(C.TTF_FontFaceStyleName(f.font)), nil
}

// GlyphIsProvided checks whether a glyph is provided by f or not.
func (f Font) GlyphIsProvided(ch uint16) (bool, error) {
	if f.font == nil {
		return false, InvalidFont
	}
	if C.TTF_GlyphIsProvided(f.font, C.Uint16(ch)) == 0 {
		return false, nil
	}
	return true, nil
}

// GlyphMetrics gets the metrics of a glyph.
//
// See http://freetype.sourceforge.net/freetype2/docs/tutorial/step2.html
func (f Font) GlyphMetrics(ch uint16) (minx, maxx, miny, maxy, advance int, err error) {
	if f.font == nil {
		err = InvalidFont
		return
	}
	i := int(C.TTF_GlyphMetrics(f.font, C.Uint16(ch),
		(*C.int)(unsafe.Pointer(&minx)),
		(*C.int)(unsafe.Pointer(&maxx)),
		(*C.int)(unsafe.Pointer(&miny)),
		(*C.int)(unsafe.Pointer(&maxy)),
		(*C.int)(unsafe.Pointer(&advance))))
	if i != 0 {
		err = sdlError(i)
	}
	return
}

// Get the dimensions of a rendered string of text.
func (f Font) Size(text string) (w, h int, err error) {
	if f.font == nil {
		err = InvalidFont
		return
	}
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))

	i := int(C.TTF_SizeUTF8(f.font, cstr, (*C.int)(unsafe.Pointer(&w)),
		(*C.int)(unsafe.Pointer(&h))))
	if i != 0 {
		err = sdlError(i)
	}
	return
}

// RenderTextSolid creates an 8-bit palettized surface and render the given text at
// fast quality with the given font and color.  The 0 pixel is the colorkey, giving
// a transparent background, and the 1 pixel is set to the text color.
func (f Font) RenderTextSolid(text string, fg sdl.Color) (*sdl.Surface, error) {
	if f.font == nil {
		return nil, InvalidFont
	}
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))

	s := (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Solid(f.font,
		cstr, *(*C.SDL_Color)(unsafe.Pointer(&fg)))))
	if s == nil {
		return nil, sdlError(0)
	}
	return s, nil
}

// RenderTextShaded creates an 8-bit palattized surface and renders the given text at
// high quality with the given font and colors.  The 0 pixel is background, while
// other pixels have varying degrees of the foreground color.
func (f Font) RenderTextShaded(text string, fg, bg sdl.Color) (*sdl.Surface, error) {
	if f.font == nil {
		return nil, InvalidFont
	}
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))

	s := (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Shaded(f.font,
		cstr, *(*C.SDL_Color)(unsafe.Pointer(&fg)),
		*(*C.SDL_Color)(unsafe.Pointer(&bg)))))
	if s == nil {
		return nil, sdlError(0)
	}
	return s, nil
}

// RenderTextBlended creates a 32-bit ARGB surface and renders the given text at
// high quality, using alpha blending to dither the font with the given color.
func (f Font) RenderTextBlended(text string, fg sdl.Color) (*sdl.Surface, error) {
	if f.font == nil {
		return nil, InvalidFont
	}
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))

	s := (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Blended(f.font,
		cstr, *(*C.SDL_Color)(unsafe.Pointer(&fg)))))
	if s == nil {
		return nil, sdlError(0)
	}
	return s, nil
}

// RenderTextBlendedWrapped creates a 32-bit ARGB surface and renders the
// given text at high quality, using alpha blending to dither the font with
// the given color.  Text is wrapped to multiple lines on line endings and on
// word boundaries if it extends beyond wrapLength in pixels.
func (f Font) RenderTextBlendedWrapped(text string, fg sdl.Color, wrapLength uint32) (*sdl.Surface, error) {
	if f.font == nil {
		return nil, InvalidFont
	}
	cstr := C.CString(text)
	defer C.free(unsafe.Pointer(cstr))

	s := (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Blended_Wrapped(f.font,
		cstr, *(*C.SDL_Color)(unsafe.Pointer(&fg)), C.Uint32(wrapLength))))
	if s == nil {
		return nil, sdlError(0)
	}
	return s, nil
}

// Close closes the font file.
func (f Font) Close() {
	if f.font == nil {
		return
	}
	C.TTF_CloseFont(f.font)
	f.font = nil
}

// Quit cleans up the TTF engine.
func Quit() {
	C.TTF_Quit()
}

// Wasinit checks if the TTF engine is intialized.
func WasInit() bool {
	if C.TTF_WasInit() == 0 {
		return false
	}
	return true
}

// GetKerningSize gets the kerning size of two glyphs.
func (f Font) GetkerningSize(prevIndex, index int) (int, error) {
	if f.font == nil {
		return 0, InvalidFont
	}
	return int(C.TTF_GetFontKerningSize(f.font, C.int(prevIndex), C.int(index))), nil
}
