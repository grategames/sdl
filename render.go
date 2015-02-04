// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sdl

/*
#include "SDL.h"
*/
import "C"

import (
	"reflect"
	"unsafe"
)

type RendererInfo struct {
	Name                string              // The name of the renderer
	Flags               uint32              // Supported RendererFlags
	Num_texture_formats uint32              // The number of available texture formats
	Texture_formats     [16]PixelFormatEnum // The available texture formats
	Max_texture_width   int32               // The maximimum texture width
	Max_texture_height  int32               // The maximimum texture height
}

// Information on the capabilities of a render driver or context.
type RendererFlags uint32

const (
	// The renderer is a software fallback
	RENDERER_SOFTWARE RendererFlags = C.SDL_RENDERER_SOFTWARE
	// The renderer uses hardware acceleration
	RENDERER_ACCELERATED RendererFlags = C.SDL_RENDERER_ACCELERATED
	// Present is synchronized with the refresh rate
	RENDERER_PRESENTVSYNC RendererFlags = C.SDL_RENDERER_PRESENTVSYNC
	// The renderer supports rendering to texture
	RENDERER_TARGETTEXTURE RendererFlags = C.SDL_RENDERER_TARGETTEXTURE
)

// The access pattern allowed for a texture
type TextureAccess uint32

const (
	// Changes rarely, not lockable
	TEXTUREACCESS_STATIC TextureAccess = C.SDL_TEXTUREACCESS_STATIC
	// Changes frequently, lockable
	TEXTUREACCESS_STREAMING TextureAccess = C.SDL_TEXTUREACCESS_STREAMING
	// Texture can be used as a render target
	TEXTUREACCESS_TARGET TextureAccess = C.SDL_TEXTUREACCESS_TARGET
)

// The texture channel modulation used in renderer.Copy
type TextureModulate uint32

const (
	// No modulation
	TEXTUREMODULATE_NONE TextureModulate = C.SDL_TEXTUREMODULATE_NONE
	// srcC = srcC * color
	TEXTUREMODULATE_COLOR TextureModulate = C.SDL_TEXTUREMODULATE_COLOR
	// srcA = srcA * alpha
	TEXTUREMODULATE_ALPHA TextureModulate = C.SDL_TEXTUREMODULATE_ALPHA
)

type RendererFlip uint32

const (
	// Do not flip
	FLIP_NONE RendererFlip = C.SDL_FLIP_NONE
	// flip horizontally
	FLIP_HORIZONTAL RendererFlip = C.SDL_FLIP_HORIZONTAL
	// flip vertically
	FLIP_VERTICAL RendererFlip = C.SDL_FLIP_VERTICAL
)

// A structure representing rendering state
type Renderer struct {
	ptr *C.SDL_Renderer
}

// An efficient driver-specific representation of pixel data
type Texture struct {
	ptr *C.SDL_Texture
}

// GetNumRenderDrivers returns the number of 2D rendering drivers available for
// the current display.
func GetNumRenderDrivers() int {
	return int(C.SDL_GetNumRenderDrivers())
}

// GetRenderDriverInfo returns information about the 2D rendering driver
// specified by index for the current display.
func GetRenderDriverInfo(index int) (*RendererInfo, error) {
	info := new(C.SDL_RendererInfo)
	r := int(C.SDL_GetRenderDriverInfo(C.int(index), info))

	if r != 0 {
		return nil, sdlError(r)
	}

	return newRendererInfo(info), nil
}

// CreateWindowAndRenderer creates a window and default renderer.
func CreateWindowAndRenderer(width, height int, window_flags WindowFlags) (Window, Renderer, error) {
	var window Window
	var renderer Renderer
	r := int(C.SDL_CreateWindowAndRenderer(C.int(width), C.int(height),
		C.Uint32(window_flags), &window.ptr,
		&renderer.ptr))
	if r != 0 {
		return window, Renderer{}, sdlError(r)
	}
	return window, renderer, nil
}

// Create a 2D rendering context for the window.  index is the index of the
// rendering driver to initialize, or -1 to initialize the first one supporting
// the requested flags.
func (window Window) CreateRenderer(index int, flags RendererFlags) (Renderer, error) {
	renderer := Renderer{}
	renderer.ptr = C.SDL_CreateRenderer(window.ptr,
		C.int(index), C.Uint32(flags))
	if renderer.ptr == nil {
		return renderer, sdlError(0)
	}
	return renderer, nil
}

// CreateSoftwareRenderer creates a 2D software rendering context for a surface.
func (surface *Surface) CreateSoftwareRenderer() (Renderer, error) {
	renderer := Renderer{}
	renderer.ptr = C.SDL_CreateSoftwareRenderer((*C.SDL_Surface)(unsafe.Pointer(surface)))
	if renderer.ptr == nil {
		return renderer, sdlError(0)
	}
	return renderer, nil
}

// GetRenderer returns the renderer associated with the window.
func (window Window) GetRenderer() (Renderer, error) {
	renderer := Renderer{}
	renderer.ptr = C.SDL_GetRenderer(window.ptr)
	if renderer.ptr == nil {
		return renderer, sdlError(0)
	}
	return renderer, nil
}

// GetRendererInfo returns information about the rendering context.
func (renderer Renderer) GetInfo() (*RendererInfo, error) {
	info := new(C.SDL_RendererInfo)
	r := int(C.SDL_GetRendererInfo(renderer.ptr, info))
	if r != 0 {
		return nil, sdlError(r)
	}
	return newRendererInfo(info), nil
}

// CreateTexture creates a texture for the rendering context.
func (renderer Renderer) CreateTexture(format PixelFormatEnum, access TextureAccess, w, h int) (Texture, error) {
	t := Texture{}
	t.ptr = C.SDL_CreateTexture(renderer.ptr,
		C.Uint32(format), C.int(access), C.int(w), C.int(h))
	if t.ptr == nil {
		return t, sdlError(0)
	}
	return t, nil
}

// CreateTextureFromSurface creates a texture from an existing surface.
func (renderer Renderer) CreateTextureFromSurface(surface *Surface) (Texture, error) {
	t := Texture{}
	t.ptr = C.SDL_CreateTextureFromSurface(renderer.ptr,
		(*C.SDL_Surface)(unsafe.Pointer(surface)))
	if t.ptr == nil {
		return t, sdlError(0)
	}
	return t, nil
}

// Query returns the attributes of a texture.
func (texture Texture) Query() (format PixelFormatEnum, access TextureAccess, w, h int, err error) {
	r := int(C.SDL_QueryTexture(texture.ptr,
		(*C.Uint32)(&format), (*C.int)(unsafe.Pointer(&access)),
		(*C.int)(unsafe.Pointer(&w)), (*C.int)(unsafe.Pointer(&h))))
	if r != 0 {
		err = sdlError(r)
	}
	return
}

// SetColorMod sets an additional color value used in render copy operations.
func (texture Texture) SetColorMod(r, g, b uint8) error {
	i := int(C.SDL_SetTextureColorMod(texture.ptr,
		C.Uint8(r), C.Uint8(g), C.Uint8(b)))
	if i != 0 {
		return sdlError(i)
	}
	return nil
}

// GetColorMod returns the additional color value used in render copy operations.
func (texture Texture) GetColorMod() (r, g, b uint8, err error) {
	i := int(C.SDL_GetTextureColorMod(texture.ptr,
		(*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b)))
	if i != 0 {
		err = sdlError(i)
	}
	return
}

// SetAlphaMod sets an additional alpha value used in render copy operations.
func (texture Texture) SetAlphaMod(alpha uint8) error {
	r := int(C.SDL_SetTextureAlphaMod(texture.ptr, C.Uint8(alpha)))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// GetAlphaMod returns the additional alpha value used in render copy operations.
func (texture Texture) GetAlphaMod() (alpha uint8, err error) {
	r := int(C.SDL_GetTextureAlphaMod(texture.ptr, (*C.Uint8)(&alpha)))
	if r != 0 {
		err = sdlError(r)
	}
	return
}

// SetBlendMode sets the blend mode used for texture copy operations.
func (texture Texture) SetBlendMode(blendMode BlendMode) error {
	r := int(C.SDL_SetTextureBlendMode(texture.ptr,
		C.SDL_BlendMode(blendMode)))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// GetBlendMode returns the blend mode used for texture copy operations.
func (texture Texture) GetBlendMode() (BlendMode, error) {
	var blendMode BlendMode
	r := int(C.SDL_GetTextureBlendMode(texture.ptr,
		(*C.SDL_BlendMode)(&blendMode)))
	if r != 0 {
		return blendMode, sdlError(r)
	}
	return blendMode, nil
}

// Update updates the given texture rectangle with new pixel data.
func (texture Texture) Update(rect *Rect, pixels unsafe.Pointer, pitch int) error {
	r := int(C.SDL_UpdateTexture(texture.ptr,
		(*C.SDL_Rect)(unsafe.Pointer(rect)), pixels, C.int(pitch)))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// Lock locks a portion of the texture for write-only pixel access.  If rect is
// nil the entire texture will be locked. It returns the locked pixels and the
// pitch for the pixels.  Lock only works if texture was created with
// TEXTUREACCESS_STREAMING.
func (texture Texture) Lock(rect *Rect) (pixels []byte, pitch int, err error) {
	var ptr unsafe.Pointer
	r := int(C.SDL_LockTexture(texture.ptr,
		(*C.SDL_Rect)(unsafe.Pointer(rect)),
		&ptr,
		(*C.int)(unsafe.Pointer(&pitch))))

	if r != 0 {
		err = sdlError(r)
	} else {
		sh := (*reflect.SliceHeader)(unsafe.Pointer(&pixels))
		sh.Data = uintptr(ptr)

		if rect == nil {
			_, _, _, h, qErr := texture.Query()
			if qErr != nil {
				return []byte{}, 0, qErr
			}
			sh.Len = h * pitch
		} else {
			sh.Len = int(rect.H) * pitch
		}
		sh.Cap = sh.Len
	}
	return
}

// Unlock unlocks the texture, uploading the changes to video memory, if needed.
func (texture Texture) Unlock() {
	C.SDL_UnlockTexture(texture.ptr)
}

// RenderTragetSupported determines whether the renderer supports the use of render
// targets.
func (renderer Renderer) RenderTargetSupported() bool {
	if e := C.SDL_RenderTargetSupported(renderer.ptr); e == C.SDL_TRUE {
		return true
	}
	return false
}

// SetRenderTarget sets the texture to as the current rendering target.  If
// texture is nil the default render target is used.
func (renderer Renderer) SetRenderTarget(texture Texture) error {
	r := int(C.SDL_SetRenderTarget(renderer.ptr,
		texture.ptr))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// GetRenderTarget gets the current render target.
//
// FIXME: A nil texture means default render target.
func (renderer Renderer) GetRenderTarget() Texture {
	tex := Texture{}
	tex.ptr = C.SDL_GetRenderTarget(renderer.ptr)
	return tex
}

// SetLogicalSize sets device independent resolution for rendering.
//
// This function uses teh viewport and scaling functionality to allow a fixed
// logical resolution for rendering, regardless of the actual output
// resolution. If the actual output resolution doesn't have the same aspect
// ratio the output rendering will be centered within the output display.
//
// If the output display is a window, mouse events in the window will be
// filtered and scaled so they seem to arrive within the logical resolution.
//
// Note: If this function results in scaling or subpixel drawing by the
// rendering backend, it will be handled using the appropriate quality hints.
func (renderer Renderer) SetLogicalSize(w, h int32) error {
	i := int(C.SDL_RenderSetLogicalSize(renderer.ptr,
		C.int(w), C.int(h)))
	if i != 0 {
		return sdlError(i)
	}
	return nil
}

// GetLogicalSize gets device independent resolution for rendering
func (renderer Renderer) GetLogicalSize() (w, h int32) {
	C.SDL_RenderGetLogicalSize(renderer.ptr,
		(*C.int)(&w), (*C.int)(&h))
	return
}

// SetViewport sets the drawing area for rendering on the current target. If
// rect is nil the viewport is set to the entire target.
func (renderer Renderer) SetViewport(rect *Rect) error {
	r := int(C.SDL_RenderSetViewport(renderer.ptr,
		(*C.SDL_Rect)(unsafe.Pointer(rect))))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// GetViewport returns the drawing area for the current target.
func (renderer Renderer) GetViewport() (*Rect, error) {
	rect := new(Rect)
	C.SDL_RenderGetViewport(renderer.ptr,
		(*C.SDL_Rect)(unsafe.Pointer(rect)))
	if rect == nil {
		return nil, sdlError(0)
	}
	return rect, nil
}

// SetScale sets teh drawing scale for rendering on the current target.
//
// The drawing coordinates are scaled by the x/w scaling factors before they
// are used by the renderer.  This allow resolution independent drawing with
// a single coordinate system.
//
// Note: If this results in scaling or subpixel drawing by the rendering
// backend, it will be handled using the appropriate quality hints.  For
// best results use integer scaling factors.
func (renderer Renderer) SetScale(scaleX, scaleY float32) error {
	i := int(C.SDL_RenderSetScale(renderer.ptr,
		C.float(scaleX), C.float(scaleY)))
	if i != 0 {
		return sdlError(i)
	}
	return nil
}

// GetScale gets the drawing scale for the current target.
func (renderer Renderer) GetScale() (scaleX, scaleY float32) {
	C.SDL_RenderGetScale(renderer.ptr,
		(*C.float)(&scaleX), (*C.float)(&scaleY))
	return
}

// SetDrawColor sets the color used for drawing operations (Rect, Line and Clear).
func (renderer Renderer) SetDrawColor(r, g, b, a uint8) error {
	i := int(C.SDL_SetRenderDrawColor(renderer.ptr,
		C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)))
	if i != 0 {
		return sdlError(i)
	}
	return nil
}

// GetDrawColor returns the color used for drawing operations (Rect, Line and Clear).
func (renderer Renderer) GetDrawColor() (r, g, b, a uint8, err error) {
	i := int(C.SDL_GetRenderDrawColor(renderer.ptr,
		(*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b), (*C.Uint8)(&a)))
	if i != 0 {
		err = sdlError(i)
	}
	return
}

// SetDrawBlendMode sets the blend mode used for drawing operations (Fill and Line).
func (renderer Renderer) SetDrawBlendMode(blendMode BlendMode) error {
	r := int(C.SDL_SetRenderDrawBlendMode(renderer.ptr,
		C.SDL_BlendMode(blendMode)))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// GetDrawBlendMode returns the blend mode used for drawing operations.
func (renderer Renderer) GetDrawBlendMode() (BlendMode, error) {
	var blendMode BlendMode
	r := int(C.SDL_GetRenderDrawBlendMode(renderer.ptr,
		(*C.SDL_BlendMode)(&blendMode)))
	if r != 0 {
		return blendMode, sdlError(r)
	}
	return blendMode, nil
}

// Clear clears the current rendering target with the drawing color.  It clears
// the entire rendering target, ignoring the viewport.
func (renderer Renderer) Clear() error {
	r := int(C.SDL_RenderClear(renderer.ptr))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// DrawPoint draws a point on the current rendering target.
func (renderer Renderer) DrawPoint(x, y int) error {
	r := int(C.SDL_RenderDrawPoint(renderer.ptr, C.int(x), C.int(y)))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// DrawPoints draws multiple points on the current rendering target.
func (renderer Renderer) DrawPoints(points []Point) error {
	var ptr *C.SDL_Point
	if len(points) > 0 {
		ptr = (*C.SDL_Point)(unsafe.Pointer(&points[0]))
	}

	r := int(C.SDL_RenderDrawPoints(renderer.ptr, ptr, C.int(len(points))))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// DrawLine draws a line on the current rendering target.
func (renderer Renderer) DrawLine(x1, y1, x2, y2 int) error {
	r := int(C.SDL_RenderDrawLine(renderer.ptr, C.int(x1), C.int(y1),
		C.int(x2), C.int(y2)))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// DrawLines draws a series of connected lines on the current rendering target.
func (renderer Renderer) DrawLines(points []Point) error {
	var ptr *C.SDL_Point
	if len(points) > 0 {
		ptr = (*C.SDL_Point)(unsafe.Pointer(&points[0]))
	}

	r := int(C.SDL_RenderDrawLines(renderer.ptr, ptr, C.int(len(points))))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// DrawRect draws a rectangle on the current rendering target. If rect is nil
// the entire rendering target is outlined.
func (renderer Renderer) DrawRect(rect *Rect) error {
	r := int(C.SDL_RenderDrawRect(renderer.ptr,
		(*C.SDL_Rect)(unsafe.Pointer(rect))))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// DrawRects draws some number of rectangles on the current rendering target.
func (renderer Renderer) DrawRects(rects []Rect) error {
	var ptr *C.SDL_Rect
	if len(rects) > 0 {
		ptr = (*C.SDL_Rect)(unsafe.Pointer(&rects[0]))
	}

	r := int(C.SDL_RenderDrawRects(renderer.ptr, ptr, C.int(len(rects))))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// FillRect fills a rectangle on the current rendering target with the drawing
// color. If rect is nil the entire rendering target is filled.
func (renderer Renderer) FillRect(rect *Rect) error {
	r := int(C.SDL_RenderFillRect(renderer.ptr,
		(*C.SDL_Rect)(unsafe.Pointer(rect))))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// FillRects fills some number of rectangles on the current rendering target
// with the drawing color.
func (renderer Renderer) FillRects(rects []Rect) error {
	var ptr *C.SDL_Rect
	if len(rects) > 0 {
		ptr = (*C.SDL_Rect)(unsafe.Pointer(&rects[0]))
	}

	r := int(C.SDL_RenderFillRects(renderer.ptr, ptr, C.int(len(rects))))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// Copy copies a portion of the texture to the current rendering target. If
// srcrect is nil the entire texture is copied.  If dstrect is nil the entire
// rendering target is filled.
func (renderer Renderer) Copy(texture Texture, srcrect, dstrect *Rect) error {
	r := int(C.SDL_RenderCopy(renderer.ptr, texture.ptr,
		(*C.SDL_Rect)(unsafe.Pointer(srcrect)),
		(*C.SDL_Rect)(unsafe.Pointer(dstrect))))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// CopyEx copies a portion of the source texture to the current rendering
// target, rotating it by angle around the given center.  If srcrect is nil
// the entire texture is copied.  If dstrect is nil the entire rendering
// target is filled.  If center is nil rotation will be done around
// (dstrect.W/2, dstrect.H/2).
func (renderer Renderer) CopyEx(texture Texture, srcrect, dstrect *Rect,
	angle float64, center *Point, flip RendererFlip) error {
	r := int(C.SDL_RenderCopyEx(renderer.ptr,
		texture.ptr,
		(*C.SDL_Rect)(unsafe.Pointer(srcrect)),
		(*C.SDL_Rect)(unsafe.Pointer(dstrect)), C.double(angle),
		(*C.SDL_Point)(unsafe.Pointer(center)),
		C.SDL_RendererFlip(flip)))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

/*
!!!This needs to try and be memory safe.!!!

// ReadPixels reads pixels from the current rendering target.  This is a very
// slow operation, and should not be used frequently.
//
// WARNING!:  There is no checks to prevent buffer overflows.
func (renderer *Renderer) ReadPixels(rect *Rect, format uint32, pixels unsafe.Pointer, pitch int) int {
	return int(C.SDL_RenderReadPixels((*C.SDL_Renderer)(renderer),
		(*C.SDL_Rect)(unsafe.Pointer(rect)), C.Uint32(format), pixels, C.int(pitch)))
}
*/

// Present updates the screen with the rendering performed.
func (renderer Renderer) Present() {
	C.SDL_RenderPresent(renderer.ptr)
}

// Destroy destroys the texture.
func (texture Texture) Destroy() {
	C.SDL_DestroyTexture(texture.ptr)
	texture.ptr = nil
}

// Destroy destroys the rendering context and frees associated textures.
func (renderer Renderer) Destroy() {
	C.SDL_DestroyRenderer(renderer.ptr)
	renderer.ptr = nil
}

// GLBind binds the texture to the current OpenGL/ES/ES2 context for use with
// OpenGL instructions.
func (texture Texture) GLBind() (w, h float32, err error) {
	r := int(C.SDL_GL_BindTexture(texture.ptr,
		(*C.float)(&w), (*C.float)(&h)))
	if r != 0 {
		err = sdlError(r)
	}
	return
}

// GLUnbind unbinds a texture from the current OpenGL/ES/ES2 context.
func (texture Texture) GLUnbind() error {
	r := int(C.SDL_GL_UnbindTexture(texture.ptr))
	if r != 0 {
		return sdlError(r)
	}
	return nil
}

// newRendererInfo creates a RendererInfo from a C.SDL_RendererInfo.
func newRendererInfo(info *C.SDL_RendererInfo) *RendererInfo {
	name := C.GoString(info.name)

	return &RendererInfo{
		name,
		uint32(info.flags),
		uint32(info.num_texture_formats),
		*(*[16]PixelFormatEnum)(unsafe.Pointer(&info.texture_formats[0])),
		int32(info.max_texture_width),
		int32(info.max_texture_height),
	}
}
